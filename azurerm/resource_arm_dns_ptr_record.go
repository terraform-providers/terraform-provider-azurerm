package azurerm

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/dns/mgmt/2018-03-01-preview/dns"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmDnsPtrRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmDnsPtrRecordCreateOrUpdate,
		Read:   resourceArmDnsPtrRecordRead,
		Update: resourceArmDnsPtrRecordCreateOrUpdate,
		Delete: resourceArmDnsPtrRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
			},

			"resource_group_name": resourceGroupNameSchema(),

			"zone_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"records": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"ttl": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"tags": tagsSchema(),
		},
	}
}

func resourceArmDnsPtrRecordCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).dnsClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)
	resGroup := d.Get("resource_group_name").(string)
	zoneName := d.Get("zone_name").(string)
	ttl := int64(d.Get("ttl").(int))
	tags := d.Get("tags").(map[string]interface{})

	records, err := expandAzureRmDnsPtrRecords(d)
	if err != nil {
		return err
	}

	parameters := dns.RecordSet{
		RecordSetProperties: &dns.RecordSetProperties{
			Metadata:   expandTags(tags),
			TTL:        &ttl,
			PtrRecords: &records,
		},
	}

	eTag := ""
	ifNoneMatch := "" // set to empty to allow updates to records after creation
	resp, err := client.CreateOrUpdate(ctx, resGroup, zoneName, name, dns.PTR, parameters, eTag, ifNoneMatch)
	if err != nil {
		return err
	}

	if resp.ID == nil {
		return fmt.Errorf("Cannot read DNS PTR Record %s (resource group %s) ID", name, resGroup)
	}

	d.SetId(*resp.ID)

	return resourceArmDnsPtrRecordRead(d, meta)
}

func resourceArmDnsPtrRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient)
	dnsClient := client.dnsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resGroup := id.ResourceGroup
	name := id.Path["PTR"]
	zoneName := id.Path["dnszones"]

	resp, err := dnsClient.Get(ctx, resGroup, zoneName, name, dns.PTR)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading DNS PTR record %s: %+v", name, err)
	}

	d.Set("name", name)
	d.Set("resource_group_name", resGroup)
	d.Set("zone_name", zoneName)
	d.Set("ttl", resp.TTL)
	d.Set("etag", resp.Etag)

	if err := d.Set("records", flattenAzureRmDnsPtrRecords(resp.PtrRecords)); err != nil {
		return err
	}
	flattenAndSetTags(d, resp.Metadata)

	return nil
}

func resourceArmDnsPtrRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient)
	dnsClient := client.dnsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resGroup := id.ResourceGroup
	name := id.Path["PTR"]
	zoneName := id.Path["dnszones"]

	resp, err := dnsClient.Delete(ctx, resGroup, zoneName, name, dns.PTR, "")
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			return nil
		}

		return fmt.Errorf("Error deleting DNS PTR Record %s: %+v", name, err)
	}

	return nil
}

func flattenAzureRmDnsPtrRecords(records *[]dns.PtrRecord) []string {
	results := make([]string, 0, len(*records))

	if records != nil {
		for _, record := range *records {
			results = append(results, *record.Ptrdname)
		}
	}

	return results
}

func expandAzureRmDnsPtrRecords(d *schema.ResourceData) ([]dns.PtrRecord, error) {
	recordStrings := d.Get("records").(*schema.Set).List()
	records := make([]dns.PtrRecord, len(recordStrings))

	for i, v := range recordStrings {
		fqdn := v.(string)
		records[i] = dns.PtrRecord{
			Ptrdname: &fqdn,
		}
	}

	return records, nil
}

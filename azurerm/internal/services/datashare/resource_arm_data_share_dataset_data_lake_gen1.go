package datashare

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/datashare/mgmt/2019-11-01/datashare"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datashare/helper"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datashare/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datashare/validate"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmDataShareDataSetDataLakeGen1() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmDataShareDataSetDataLakeGen1Create,
		Read:   resourceArmDataShareDataSetDataLakeGen1Read,
		Delete: resourceArmDataShareDataSetDataLakeGen1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.DataShareDataSetID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DatashareDataSetName(),
			},

			"data_share_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataShareID,
			},

			"data_lake_store": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: azure.ValidateDataLakeAccountName(),
						},

						"resource_group_name": azure.SchemaResourceGroupName(),

						"subscription_id": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.IsUUID,
						},
					},
				},
			},

			"folder_path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"file_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceArmDataShareDataSetDataLakeGen1Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataShare.DataSetClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	shareId, err := parse.DataShareID(d.Get("data_share_id").(string))
	if err != nil {
		return err
	}

	existing, err := client.Get(ctx, shareId.ResourceGroup, shareId.AccountName, shareId.Name, name)
	if err != nil {
		if !utils.ResponseWasNotFound(existing.Response) {
			return fmt.Errorf("checking for present of existing  DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q): %+v", name, shareId.ResourceGroup, shareId.AccountName, shareId.Name, err)
		}
	}
	existingId := helper.GetAzurermDataShareDataSetId(existing.Value)
	if existingId != nil && *existingId != "" {
		return tf.ImportAsExistsError("azurerm_data_share_dataset_data_lake_gen1", *existingId)
	}

	var dataSet datashare.BasicDataSet

	if fileName, ok := d.GetOk("file_name"); ok {
		dataSet = datashare.ADLSGen1FileDataSet{
			Kind: datashare.KindAdlsGen1File,
			ADLSGen1FileProperties: &datashare.ADLSGen1FileProperties{
				AccountName:    utils.String(d.Get("data_lake_store.0.name").(string)),
				ResourceGroup:  utils.String(d.Get("data_lake_store.0.resource_group_name").(string)),
				SubscriptionID: utils.String(d.Get("data_lake_store.0.subscription_id").(string)),
				FolderPath:     utils.String(d.Get("folder_path").(string)),
				FileName:       utils.String(fileName.(string)),
			},
		}
	} else {
		dataSet = datashare.ADLSGen1FolderDataSet{
			Kind: datashare.KindAdlsGen1Folder,
			ADLSGen1FolderProperties: &datashare.ADLSGen1FolderProperties{
				AccountName:    utils.String(d.Get("data_lake_store.0.name").(string)),
				ResourceGroup:  utils.String(d.Get("data_lake_store.0.resource_group_name").(string)),
				SubscriptionID: utils.String(d.Get("data_lake_store.0.subscription_id").(string)),
				FolderPath:     utils.String(d.Get("folder_path").(string)),
			},
		}
	}

	if _, err := client.Create(ctx, shareId.ResourceGroup, shareId.AccountName, shareId.Name, name, dataSet); err != nil {
		return fmt.Errorf("creating/updating DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q): %+v", name, shareId.ResourceGroup, shareId.AccountName, shareId.Name, err)
	}

	resp, err := client.Get(ctx, shareId.ResourceGroup, shareId.AccountName, shareId.Name, name)
	if err != nil {
		return fmt.Errorf("retrieving DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q): %+v", name, shareId.ResourceGroup, shareId.AccountName, shareId.Name, err)
	}

	respId := helper.GetAzurermDataShareDataSetId(resp.Value)
	if respId == nil || *respId == "" {
		return fmt.Errorf("empty or nil ID returned for DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q)", name, shareId.ResourceGroup, shareId.AccountName, shareId.Name)
	}

	d.SetId(*respId)
	return resourceArmDataShareDataSetDataLakeGen1Read(d, meta)
}

func resourceArmDataShareDataSetDataLakeGen1Read(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataShare.DataSetClient
	shareClient := meta.(*clients.Client).DataShare.SharesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataShareDataSetID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.AccountName, id.ShareName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] DataShare %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, id.ShareName, err)
	}
	d.Set("name", id.Name)
	shareResp, err := shareClient.Get(ctx, id.ResourceGroup, id.AccountName, id.ShareName)
	if err != nil {
		return fmt.Errorf("retrieving DataShare %q (Resource Group %q / accountName %q): %+v", id.ShareName, id.ResourceGroup, id.AccountName, err)
	}
	if shareResp.ID == nil || *shareResp.ID == "" {
		return fmt.Errorf("reading ID of DataShare %q (Resource Group %q / accountName %q): ID is empt", id.ShareName, id.ResourceGroup, id.AccountName)
	}
	d.Set("data_share_id", shareResp.ID)

	switch resp := resp.Value.(type) {
	case datashare.ADLSGen1FileDataSet:
		if props := resp.ADLSGen1FileProperties; props != nil {
			if err := d.Set("data_lake_store", flattenAzureRmDataShareDataSetDataLakeStore(props.AccountName, props.ResourceGroup, props.SubscriptionID)); err != nil {
				return fmt.Errorf("setting `data_lake_store`: %+v", err)
			}
			d.Set("folder_path", props.FolderPath)
			d.Set("file_name", props.FileName)
			d.Set("display_name", props.DataSetID)
		}

	case datashare.ADLSGen1FolderDataSet:
		if props := resp.ADLSGen1FolderProperties; props != nil {
			if err := d.Set("data_lake_store", flattenAzureRmDataShareDataSetDataLakeStore(props.AccountName, props.ResourceGroup, props.SubscriptionID)); err != nil {
				return fmt.Errorf("setting `data_lake_store`: %+v", err)
			}
			d.Set("folder_path", props.FolderPath)
			d.Set("display_name", props.DataSetID)
		}

	default:
		return fmt.Errorf("data share dataset %q (Resource Group %q / accountName %q / shareName %q) is not a datalake store gen1 dataset", id.Name, id.ResourceGroup, id.AccountName, id.ShareName)
	}

	return nil
}

func resourceArmDataShareDataSetDataLakeGen1Delete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataShare.DataSetClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataShareDataSetID(d.Id())
	if err != nil {
		return err
	}

	if _, err := client.Delete(ctx, id.ResourceGroup, id.AccountName, id.ShareName, id.Name); err != nil {
		return fmt.Errorf("deleting DataShare DataSet %q (Resource Group %q / accountName %q / shareName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, id.ShareName, err)
	}
	return nil
}

func flattenAzureRmDataShareDataSetDataLakeStore(strName, strRG, strSubs *string) []interface{} {
	var name, rg, subs string
	if strName != nil {
		name = *strName
	}

	if strRG != nil {
		rg = *strRG
	}

	if strSubs != nil {
		subs = *strSubs
	}

	return []interface{}{
		map[string]interface{}{
			"name":                name,
			"resource_group_name": rg,
			"subscription_id":     subs,
		},
	}
}

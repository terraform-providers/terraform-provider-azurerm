package databasemigration

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/datamigration/mgmt/2018-04-19/datamigration"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmDatabaseMigrationProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmDatabaseMigrationProjectCreateUpdate,
		Read:   resourceArmDatabaseMigrationProjectRead,
		Update: resourceArmDatabaseMigrationProjectCreateUpdate,
		Delete: resourceArmDatabaseMigrationProjectDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateDatabaseMigrationProjectName,
			},

			"service_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateDatabaseMigrationServiceName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"source_platform": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					// Now that go sdk only export SQL as source platform type, we only allow it here.
					string(datamigration.ProjectSourcePlatformSQL),
				}, false),
			},

			"target_platform": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					// Now that go sdk only export SQL as source platform type, we only allow it here.
					string(datamigration.ProjectTargetPlatformSQLDB),
				}, false),
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmDatabaseMigrationProjectCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DatabaseMigration.ProjectsClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	serviceName := d.Get("service_name").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, serviceName, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for present of existing Database Migration Project (Project Name: %q / Service Name %q / Group Name %q): %+v", name, serviceName, resourceGroup, err)
			}
		}
		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_database_migration_project", *existing.ID)
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	sourcePlatform := d.Get("source_platform").(string)
	targetPlatform := d.Get("target_platform").(string)
	t := d.Get("tags").(map[string]interface{})

	parameters := datamigration.Project{
		Location: utils.String(location),
		ProjectProperties: &datamigration.ProjectProperties{
			SourcePlatform: datamigration.ProjectSourcePlatform(sourcePlatform),
			TargetPlatform: datamigration.ProjectTargetPlatform(targetPlatform),
		},
		Tags: tags.Expand(t),
	}

	if _, err := client.CreateOrUpdate(ctx, parameters, resourceGroup, serviceName, name); err != nil {
		return fmt.Errorf("Error creating Database Migration Project (Project Name %q / Service Name %q / Group Name %q): %+v", name, serviceName, resourceGroup, err)
	}

	resp, err := client.Get(ctx, resourceGroup, serviceName, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Database Migration Project (Project Name %q / Service Name %q / Group Name %q): %+v", name, serviceName, resourceGroup, err)
	}
	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("Cannot read Database Migration Project (Project Name %q / Service Name %q / Group Name %q) ID", name, serviceName, resourceGroup)
	}
	d.SetId(*resp.ID)

	return resourceArmDatabaseMigrationProjectRead(d, meta)
}

func resourceArmDatabaseMigrationProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DatabaseMigration.ProjectsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	name := id.Path["projects"]
	resourceGroup := id.ResourceGroup
	serviceName := id.Path["services"]

	resp, err := client.Get(ctx, resourceGroup, serviceName, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Database Migration Project %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error reading Database Migration Project (Project Name %q / Service Name %q / Group Name %q): %+v", name, serviceName, resourceGroup, err)
	}
	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("Cannot read Database Migration Project (Project Name %q / Service Name %q / Group Name %q) ID", name, serviceName, resourceGroup)
	}

	d.SetId(*resp.ID)

	d.Set("name", resp.Name)
	d.Set("service_name", serviceName)
	d.Set("resource_group_name", resourceGroup)

	location := ""
	if resp.Location != nil {
		location = azure.NormalizeLocation(*resp.Location)
	}
	d.Set("location", location)

	if prop := resp.ProjectProperties; prop != nil {
		d.Set("source_platform", string(prop.SourcePlatform))
		d.Set("target_platform", string(prop.TargetPlatform))
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDatabaseMigrationProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DatabaseMigration.ProjectsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	serviceName := id.Path["services"]
	name := id.Path["projects"]

	deleteRunningTasks := true
	if _, err := client.Delete(ctx, resourceGroup, serviceName, name, &deleteRunningTasks); err != nil {
		return fmt.Errorf("Error deleting Database Migration Project (Project Name %q / Service Name %q / Group Name %q): %+v", name, serviceName, resourceGroup, err)
	}

	return nil
}

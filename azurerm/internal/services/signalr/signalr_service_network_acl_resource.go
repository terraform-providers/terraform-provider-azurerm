package signalr

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/signalr/mgmt/2020-05-01/signalr"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/locks"
	networkValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/signalr/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/signalr/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmSignalRServiceNetworkACL() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceSignalRServiceNetworkACLCreateUpdate,
		Read:   resourceSignalRServiceNetworkACLRead,
		Update: resourceSignalRServiceNetworkACLCreateUpdate,
		Delete: resourceSignalRServiceNetworkACLDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Importer: pluginsdk.DefaultImporter(),

		Schema: map[string]*pluginsdk.Schema{
			"signalr_service_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.ServiceID,
			},

			"default_action": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(signalr.Allow),
					string(signalr.Deny),
				}, false),
			},

			"public_network": {
				Type:     pluginsdk.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						// API response includes the `Trace` type but it isn't in rest api client.
						// https://github.com/Azure/azure-rest-api-specs/issues/14923
						"allowed_request_types": {
							Type:          pluginsdk.TypeSet,
							Optional:      true,
							ConflictsWith: []string{"public_network.0.denied_request_types"},
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									string(signalr.ClientConnection),
									string(signalr.RESTAPI),
									string(signalr.ServerConnection),
									"Trace",
								}, false),
							},
						},

						"denied_request_types": {
							Type:          pluginsdk.TypeSet,
							Optional:      true,
							ConflictsWith: []string{"public_network.0.allowed_request_types"},
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									string(signalr.ClientConnection),
									string(signalr.RESTAPI),
									string(signalr.ServerConnection),
									"Trace",
								}, false),
							},
						},
					},
				},
			},

			"private_endpoint": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"id": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: networkValidate.PrivateEndpointID,
						},

						// API response includes the `Trace` type but it isn't in rest api client.
						// https://github.com/Azure/azure-rest-api-specs/issues/14923
						"allowed_request_types": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									string(signalr.ClientConnection),
									string(signalr.RESTAPI),
									string(signalr.ServerConnection),
									"Trace",
								}, false),
							},
						},

						"denied_request_types": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									string(signalr.ClientConnection),
									string(signalr.RESTAPI),
									string(signalr.ServerConnection),
									"Trace",
								}, false),
							},
						},
					},
				},
			},
		},
	}
}

func resourceSignalRServiceNetworkACLCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SignalR.Client
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ServiceID(d.Get("signalr_service_id").(string))
	if err != nil {
		return err
	}

	locks.ByName(id.SignalRName, "azurerm_signalr_service")
	defer locks.UnlockByName(id.SignalRName, "azurerm_signalr_service")

	resp, err := client.Get(ctx, id.ResourceGroup, id.SignalRName)
	if err != nil {
		return fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	parameters := resp

	if props := resp.Properties; props != nil {
		networkACL := &signalr.NetworkACLs{
			DefaultAction: signalr.ACLAction(d.Get("default_action").(string)),
			PublicNetwork: expandSignalRServicePublicNetwork(d.Get("public_network").([]interface{})),
		}

		if v, ok := d.GetOk("private_endpoint"); ok {
			networkACL.PrivateEndpoints = expandSignalRServicePrivateEndpoint(v.(*pluginsdk.Set).List(), props.PrivateEndpointConnections)
		}

		if networkACL.DefaultAction == signalr.Allow && len(*networkACL.PublicNetwork.Allow) != 0 {
			return fmt.Errorf("when `default_action` is `Allow` for `public_network`, `allowed_request_types` cannot be specified")
		} else if networkACL.DefaultAction == signalr.Deny && len(*networkACL.PublicNetwork.Deny) != 0 {
			return fmt.Errorf("when `default_action` is `Deny` for `public_network`, `denied_request_types` cannot be specified")
		}

		if networkACL.PrivateEndpoints != nil {
			for _, privateEndpoint := range *networkACL.PrivateEndpoints {
				if len(*privateEndpoint.Allow) != 0 && len(*privateEndpoint.Deny) != 0 {
					return fmt.Errorf("`allowed_request_types` and `denied_request_types` cannot be set together for `private_endpoint`")
				}

				if networkACL.DefaultAction == signalr.Allow && len(*privateEndpoint.Allow) != 0 {
					return fmt.Errorf("when `default_action` is `Allow` for `private_endpoint`, `allowed_request_types` cannot be specified")
				} else if networkACL.DefaultAction == signalr.Deny && len(*privateEndpoint.Deny) != 0 {
					return fmt.Errorf("when `default_action` is `Deny` for `private_endpoint`, `denied_request_types` cannot be specified")
				}
			}
		}

		parameters.Properties.NetworkACLs = networkACL
	}

	future, err := client.Update(ctx, id.ResourceGroup, id.SignalRName, &parameters)
	if err != nil {
		return fmt.Errorf("creating/updating NetworkACL for %s: %v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of creating/updating NetworkACL for %s: %v", id, err)
	}

	d.SetId(id.ID())

	return resourceSignalRServiceNetworkACLRead(d, meta)
}

func resourceSignalRServiceNetworkACLRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SignalR.Client
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ServiceID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.SignalRName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	d.Set("signalr_service_id", id.ID())

	if props := resp.Properties; props != nil {
		d.Set("default_action", props.NetworkACLs.DefaultAction)

		if err := d.Set("public_network", flattenSignalRServicePublicNetwork(props.NetworkACLs.PublicNetwork)); err != nil {
			return fmt.Errorf("setting `public_network`: %+v", err)
		}

		if err := d.Set("private_endpoint", flattenSignalRServicePrivateEndpoint(props.NetworkACLs.PrivateEndpoints, props.PrivateEndpointConnections)); err != nil {
			return fmt.Errorf("setting `private_endpoint`: %+v", err)
		}
	}

	return nil
}

func resourceSignalRServiceNetworkACLDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SignalR.Client
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ServiceID(d.Id())
	if err != nil {
		return err
	}

	locks.ByName(id.SignalRName, "azurerm_signalr_service")
	defer locks.UnlockByName(id.SignalRName, "azurerm_signalr_service")

	resp, err := client.Get(ctx, id.ResourceGroup, id.SignalRName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	// As this isn't a real object, so it has to update NetworkACL to default settings for the delete operation
	parameters := resp

	requestTypes := signalr.PossibleRequestTypeValues()
	requestTypes = append(requestTypes, "Trace")

	networkACL := &signalr.NetworkACLs{
		DefaultAction: signalr.Deny,
		PublicNetwork: &signalr.NetworkACL{
			Allow: &requestTypes,
		},
	}

	if resp.Properties != nil && resp.Properties.NetworkACLs != nil && resp.Properties.NetworkACLs.PrivateEndpoints != nil {
		privateEndpoints := make([]signalr.PrivateEndpointACL, 0)
		for _, item := range *resp.Properties.NetworkACLs.PrivateEndpoints {
			if item.Name != nil {
				privateEndpoints = append(privateEndpoints, signalr.PrivateEndpointACL{
					Allow: &requestTypes,
					Name:  item.Name,
				})
			}
		}
		networkACL.PrivateEndpoints = &privateEndpoints
	}

	if parameters.Properties != nil {
		parameters.Properties.NetworkACLs = networkACL
	}

	future, err := client.Update(ctx, id.ResourceGroup, id.SignalRName, &parameters)
	if err != nil {
		return fmt.Errorf("reverting NetworkACL to default settings for %s: %v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of reverting NetworkACL to default settings for %s: %v", id, err)
	}

	return nil
}

func expandSignalRServicePublicNetwork(input []interface{}) *signalr.NetworkACL {
	allowedRTs := make([]signalr.RequestType, 0)
	deniedRTs := make([]signalr.RequestType, 0)

	if len(input) != 0 && input[0] != nil {
		v := input[0].(map[string]interface{})

		for _, item := range *(utils.ExpandStringSlice(v["allowed_request_types"].(*pluginsdk.Set).List())) {
			allowedRTs = append(allowedRTs, (signalr.RequestType)(item))
		}

		for _, item := range *(utils.ExpandStringSlice(v["denied_request_types"].(*pluginsdk.Set).List())) {
			deniedRTs = append(deniedRTs, (signalr.RequestType)(item))
		}
	}

	return &signalr.NetworkACL{
		Allow: &allowedRTs,
		Deny:  &deniedRTs,
	}
}

func expandSignalRServicePrivateEndpoint(input []interface{}, privateEndpointConnections *[]signalr.PrivateEndpointConnection) *[]signalr.PrivateEndpointACL {
	results := make([]signalr.PrivateEndpointACL, 0)

	if privateEndpointConnections != nil {
		for _, privateEndpointConnection := range *privateEndpointConnections {
			result := signalr.PrivateEndpointACL{
				Allow: &[]signalr.RequestType{},
				Deny:  &[]signalr.RequestType{},
			}

			if privateEndpointConnection.Name != nil {
				result.Name = utils.String(*privateEndpointConnection.Name)
			}

			for _, item := range input {
				v := item.(map[string]interface{})
				privateEndpointId := v["id"].(string)

				if privateEndpointConnection.PrivateEndpointConnectionProperties != nil && privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint != nil && privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint.ID != nil && privateEndpointId == *privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint.ID {
					allowedRTs := make([]signalr.RequestType, 0)
					for _, item := range *(utils.ExpandStringSlice(v["allowed_request_types"].(*pluginsdk.Set).List())) {
						allowedRTs = append(allowedRTs, (signalr.RequestType)(item))
					}
					result.Allow = &allowedRTs

					deniedRTs := make([]signalr.RequestType, 0)
					for _, item := range *(utils.ExpandStringSlice(v["denied_request_types"].(*pluginsdk.Set).List())) {
						deniedRTs = append(deniedRTs, (signalr.RequestType)(item))
					}
					result.Deny = &deniedRTs

					break
				}
			}

			results = append(results, result)
		}
	}

	return &results
}

func flattenSignalRServicePublicNetwork(input *signalr.NetworkACL) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	allowRequestTypes := make([]string, 0)
	if input.Allow != nil {
		for _, item := range *input.Allow {
			allowRequestTypes = append(allowRequestTypes, (string)(item))
		}
	}
	allow := utils.FlattenStringSlice(&allowRequestTypes)

	deniedRequestTypes := make([]string, 0)
	if input.Deny != nil {
		for _, item := range *input.Deny {
			deniedRequestTypes = append(deniedRequestTypes, (string)(item))
		}
	}
	deny := utils.FlattenStringSlice(&deniedRequestTypes)

	return []interface{}{
		map[string]interface{}{
			"allowed_request_types": allow,
			"denied_request_types":  deny,
		},
	}
}

func flattenSignalRServicePrivateEndpoint(input *[]signalr.PrivateEndpointACL, privateEndpointConnections *[]signalr.PrivateEndpointConnection) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		if privateEndpointConnections != nil {
			for _, privateEndpointConnection := range *privateEndpointConnections {
				if item.Name != nil && privateEndpointConnection.Name != nil && *item.Name == *privateEndpointConnection.Name && privateEndpointConnection.PrivateEndpointConnectionProperties != nil && privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint != nil && privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint.ID != nil {
					allowedRequestTypes := make([]string, 0)
					if item.Allow != nil {
						for _, item := range *item.Allow {
							allowedRequestTypes = append(allowedRequestTypes, (string)(item))
						}
					}
					allow := utils.FlattenStringSlice(&allowedRequestTypes)

					deniedRequestTypes := make([]string, 0)
					if item.Deny != nil {
						for _, item := range *item.Deny {
							deniedRequestTypes = append(deniedRequestTypes, (string)(item))
						}
					}
					deny := utils.FlattenStringSlice(&deniedRequestTypes)

					results = append(results, map[string]interface{}{
						"id":                    *privateEndpointConnection.PrivateEndpointConnectionProperties.PrivateEndpoint.ID,
						"allowed_request_types": allow,
						"denied_request_types":  deny,
					})

					break
				}
			}
		}
	}

	return results
}

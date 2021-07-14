package batch

import (
	"context"
	"fmt"
	"time"

	batchDataplane "github.com/Azure/azure-sdk-for-go/services/batch/2020-03-01.11.0/batch"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/sdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/batch/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/batch/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type BatchJobResource struct{}

var _ sdk.ResourceWithUpdate = BatchJobResource{}

type BatchJobModel struct {
	Name        string `tfschema:"name"`
	BatchPoolId string `tfschema:"batch_pool_id"`
	DisplayName string `tfschema:"display_name"`
	Priority    int    `tfschema:"priority"`
	//MaxWallClockTime  string `tfschema:"max_wall_clock_time"`
	MaxTaskRetryCount        int               `tfschema:"max_task_retry_count"`
	CommonEnvironmentSetting map[string]string `tfschema:"common_environment_setting"`
}

func (r BatchJobResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validate.JobName,
		},
		"batch_pool_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validate.PoolID,
		},
		"display_name": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"common_environment_setting": {
			Type:     pluginsdk.TypeMap,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},
		"priority": {
			Type:         pluginsdk.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntBetween(-1000, 1000),
		},
		// TODO: identify how to represent "unlimited", it appears that a duration larger than some threshold is regarded as unlimited
		// Tracked in: https://github.com/Azure/azure-rest-api-specs/issues/15198
		//"max_wall_clock_time": {
		//	Type:         pluginsdk.TypeString,
		//	Optional:     true,
		//	Computed:     true,
		//	ValidateFunc: commonValidate.ISO8601Duration,
		//},
		"max_task_retry_count": {
			Type:         pluginsdk.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(-1),
		},
	}
}

func (r BatchJobResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r BatchJobResource) ResourceType() string {
	return "azurerm_batch_job"
}

func (r BatchJobResource) ModelObject() interface{} {
	return BatchJobModel{}
}

func (r BatchJobResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validate.JobID
}

func (r BatchJobResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {

			var model BatchJobModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}

			poolId, err := parse.PoolID(model.BatchPoolId)
			if err != nil {
				return err
			}

			accountId := parse.NewAccountID(poolId.SubscriptionId, poolId.ResourceGroup, poolId.BatchAccountName)
			client, err := metadata.Client.Batch.JobClient(ctx, accountId)
			if err != nil {
				return err
			}

			id := parse.NewJobID(poolId.SubscriptionId, poolId.ResourceGroup, poolId.BatchAccountName, poolId.Name, model.Name)

			if metadata.ResourceData.IsNewResource() {
				existing, err := r.getJob(ctx, client, id)
				if err != nil {
					if !utils.ResponseWasNotFound(existing.Response) {
						return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
					}
				}
				if !utils.ResponseWasNotFound(existing.Response) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			params := batchDataplane.JobAddParameter{
				ID:          &model.Name,
				DisplayName: &model.DisplayName,
				Priority:    utils.Int32(int32(model.Priority)),
				Constraints: &batchDataplane.JobConstraints{
					//MaxWallClockTime:  nil,
					MaxTaskRetryCount: utils.Int32(int32(model.MaxTaskRetryCount)),
				},
				CommonEnvironmentSettings: r.expandEnvironmentSettings(model.CommonEnvironmentSetting),
				PoolInfo: &batchDataplane.PoolInformation{
					PoolID: &poolId.Name,
				},
			}

			if err := r.addJob(ctx, client, id, params); err != nil {
				return err
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r BatchJobResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,

		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			id, err := parse.JobID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}
			accountId := parse.NewAccountID(id.SubscriptionId, id.ResourceGroup, id.BatchAccountName)
			client, err := metadata.Client.Batch.JobClient(ctx, accountId)
			if err != nil {
				return err
			}

			resp, err := r.getJob(ctx, client, *id)
			if err != nil {
				if utils.ResponseWasNotFound(resp.Response) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			model := BatchJobModel{
				Name:              id.Name,
				BatchPoolId:       parse.NewPoolID(id.SubscriptionId, id.ResourceGroup, id.BatchAccountName, id.PoolName).ID(),
				MaxTaskRetryCount: 0,
			}

			if resp.Priority != nil {
				model.Priority = int(*resp.Priority)
			}

			if resp.DisplayName != nil {
				model.DisplayName = *resp.DisplayName
			}

			if prop := resp.Constraints; prop != nil {
				if prop.MaxTaskRetryCount != nil {
					model.MaxTaskRetryCount = int(*prop.MaxTaskRetryCount)
				}
			}

			model.CommonEnvironmentSetting = r.flattenEnvironmentSettings(resp.CommonEnvironmentSettings)

			return metadata.Encode(&model)
		},
	}
}

func (r BatchJobResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			patch := batchDataplane.JobPatchParameter{}

			var model BatchJobModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}

			if metadata.ResourceData.HasChange("priority") {
				patch.Priority = utils.Int32(int32(model.Priority))
			}

			if metadata.ResourceData.HasChange("max_task_retry_count") {
				if patch.Constraints == nil {
					patch.Constraints = new(batchDataplane.JobConstraints)
				}
				patch.Constraints.MaxTaskRetryCount = utils.Int32(int32(model.MaxTaskRetryCount))
			}

			id, err := parse.JobID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}
			accountId := parse.NewAccountID(id.SubscriptionId, id.ResourceGroup, id.BatchAccountName)
			client, err := metadata.Client.Batch.JobClient(ctx, accountId)
			if err != nil {
				return err
			}

			if err := r.patchJob(ctx, client, *id, patch); err != nil {
				return err
			}
			return nil
		},
	}
}

func (r BatchJobResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			id, err := parse.JobID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}
			accountId := parse.NewAccountID(id.SubscriptionId, id.ResourceGroup, id.BatchAccountName)
			client, err := metadata.Client.Batch.JobClient(ctx, accountId)
			if err != nil {
				return err
			}
			if err := r.deleteJob(ctx, client, *id); err != nil {
				return err
			}
			return nil
		},
	}
}

func (r BatchJobResource) addJob(ctx context.Context, client *batchDataplane.JobClient, id parse.JobId, job batchDataplane.JobAddParameter) error {
	deadline, _ := ctx.Deadline()
	now := time.Now()
	timeout := deadline.Sub(now)
	_, err := client.Add(ctx, job, utils.Int32(int32(timeout.Seconds())), nil, nil, &date.TimeRFC1123{Time: now})
	if err != nil {
		return fmt.Errorf("creating %s: %v", id, err)
	}
	return nil
}

func (r BatchJobResource) getJob(ctx context.Context, client *batchDataplane.JobClient, id parse.JobId) (batchDataplane.CloudJob, error) {
	deadline, _ := ctx.Deadline()
	now := time.Now()
	timeout := deadline.Sub(now)
	return client.Get(ctx, id.Name, "", "", utils.Int32(int32(timeout.Seconds())), nil, nil, &date.TimeRFC1123{Time: now}, "", "", nil, nil)
}

func (r BatchJobResource) patchJob(ctx context.Context, client *batchDataplane.JobClient, id parse.JobId, job batchDataplane.JobPatchParameter) error {
	deadline, _ := ctx.Deadline()
	now := time.Now()
	timeout := deadline.Sub(now)
	_, err := client.Patch(ctx, id.Name, job, utils.Int32(int32(timeout.Seconds())), nil, nil, &date.TimeRFC1123{Time: now}, "", "", nil, nil)
	return err
}

func (r BatchJobResource) deleteJob(ctx context.Context, client *batchDataplane.JobClient, id parse.JobId) error {
	deadline, _ := ctx.Deadline()
	now := time.Now()
	timeout := deadline.Sub(now)
	_, err := client.Delete(ctx, id.Name, utils.Int32(int32(timeout.Seconds())), nil, nil, &date.TimeRFC1123{Time: now}, "", "", nil, nil)
	return err
}

func (r BatchJobResource) expandEnvironmentSettings(input map[string]string) *[]batchDataplane.EnvironmentSetting {
	if len(input) == 0 {
		return nil
	}
	m := make([]batchDataplane.EnvironmentSetting, 0, len(input))
	for k, v := range input {
		m = append(m, batchDataplane.EnvironmentSetting{
			Name:  &k,
			Value: &v,
		})
	}
	return &m
}

func (r BatchJobResource) flattenEnvironmentSettings(input *[]batchDataplane.EnvironmentSetting) map[string]string {
	if input == nil {
		return nil
	}

	m := make(map[string]string)
	for _, setting := range *input {
		if setting.Name == nil {
			continue
		}
		if setting.Value == nil {
			continue
		}
		m[*setting.Name] = *setting.Value
	}
	return m
}

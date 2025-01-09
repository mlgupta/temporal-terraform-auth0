package workflows

import (
	"time"

	"github.com/dbsentry/temporal-terraform-auth0/activities"
	"go.temporal.io/sdk/workflow"
)

const TerraformTaskQueue = "TERRAFORM_TASK_QUEUE"

func TerraformWorkflow(ctx workflow.Context) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 300,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	logger := workflow.GetLogger(ctx)

	var result string
	err := workflow.ExecuteActivity(ctx, activities.TerraformInitAuth0Activity).Get(ctx, &result)
	if err != nil {
		logger.Error("Error in TerraformInitAuth0Activity", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplyAuth0Activity).Get(ctx, &result)
	if err != nil {
		logger.Error("Error in TerraformApplyAuth0Activity", err)
		return "", err
	}

	return result, err
}

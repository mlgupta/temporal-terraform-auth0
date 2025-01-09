package activities

import (
	"context"

	"github.com/dbsentry/temporal-terraform-auth0/terraform"
	"go.temporal.io/sdk/activity"
)

func TerraformInitAuth0Activity(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)

	tf, err := terraform.NewTerraform(terraform.AUTH0_TF_DIRECTORY)
	if err != nil {
		logger.Error("Error creating terraform: ", err)
		return "", err
	}

	result, err := tf.TFInit(terraform.AUTH0_TF_DIRECTORY)
	if err != nil {
		logger.Error("Error initializing terraform: ", err)
		return "", err
	}
	return result, nil
}

func TerraformApplyAuth0Activity(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)

	tf, err := terraform.NewTerraform(terraform.AUTH0_TF_DIRECTORY)
	if err != nil {
		logger.Error("Error creating terraform: ", err)
		return "", err
	}

	result, err := tf.TFApply(terraform.AUTH0_TF_DIRECTORY)
	if err != nil {
		logger.Error("Error applying terraform: ", err)
		return "", err
	}
	return result, nil
}

func TerraformOutputAuth0Activity(ctx context.Context) (map[string]string, error) {
	return nil, nil
}

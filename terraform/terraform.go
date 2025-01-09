package terraform

//go:generate mockgen -destination=mocks/terraform.go . TF

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

const (
	EXEC_PATH          = "/opt/local/bin/terraform"
	AUTH0_TF_DIRECTORY = "./terraform/auth0"
)

type (
	Terraform struct {
		tf *tfexec.Terraform
	}

	TF interface {
		TFInit(directory string) (string, error)
		TFApply(directory string) (string, error)
	}
)

func NewTerraform(directory string) (TF, error) {
	tf, err := tfexec.NewTerraform(directory, EXEC_PATH)
	if err != nil {
		return Terraform{}, err
	}

	terraform := Terraform{
		tf: tf,
	}

	return terraform, nil
}

func (tf Terraform) TFInit(directory string) (string, error) {
	err := tf.tf.Init(context.Background())
	if err != nil {
		return "", err
	}
	return "Terraform Initialized successfully", nil
}

func (tf Terraform) TFApply(directory string) (string, error) {
	err := tf.tf.Apply(context.Background())
	if err != nil {
		return "", err
	}
	return "Terraform applied successfully", nil
}

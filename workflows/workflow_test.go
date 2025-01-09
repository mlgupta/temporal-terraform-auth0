package workflows

import (
	"fmt"
	"testing"

	"github.com/dbsentry/temporal-terraform-auth0/activities"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestTerraformWorkflow(t *testing.T) {

	testVars := []struct {
		name             string
		expectError      bool
		activityResponse []struct {
			result string
			error  error
		}
		result string
	}{
		{
			name:        "Error Terraform Init Activity",
			expectError: true,
			activityResponse: []struct {
				result string
				error  error
			}{
				{
					result: "",
					error:  fmt.Errorf("TerraformInitAuth0Activity Error"),
				},
				{
					result: "",
					error:  nil,
				},
			},
			result: "",
		},
		{
			name:        "Error Terraform Apply Activity",
			expectError: true,
			activityResponse: []struct {
				result string
				error  error
			}{
				{
					result: "result1",
					error:  nil,
				},
				{
					result: "",
					error:  fmt.Errorf("TerraformApplyAuth0Activity Error"),
				},
			},
			result: "",
		},
		{
			name:        "Success",
			expectError: false,
			activityResponse: []struct {
				result string
				error  error
			}{
				{
					result: "result1",
					error:  nil,
				},
				{
					result: "result2",
					error:  nil,
				},
			},
			result: "result2",
		},
	}

	for _, m := range testVars {

		testSuite := &testsuite.WorkflowTestSuite{}
		env := testSuite.NewTestWorkflowEnvironment()

		env.OnActivity(activities.TerraformInitAuth0Activity, mock.Anything).Return(m.activityResponse[0].result, m.activityResponse[0].error)
		env.OnActivity(activities.TerraformApplyAuth0Activity, mock.Anything).Return(m.activityResponse[1].result, m.activityResponse[1].error)

		env.ExecuteWorkflow(TerraformWorkflow)

		if m.expectError {
			require.Error(t, env.GetWorkflowError())
		} else {
			require.NoError(t, env.GetWorkflowError())
			var result string
			require.NoError(t, env.GetWorkflowResult(&result))
			require.Equal(t, m.result, result)
		}
	}
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsSpotFleetRequestInvalidAllocationStrategyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_spot_fleet_request" "foo" {
	allocation_strategy = "highestPrice"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsSpotFleetRequestInvalidAllocationStrategyRule(),
					Message: `allocation_strategy is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_spot_fleet_request" "foo" {
	allocation_strategy = "lowestPrice"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsSpotFleetRequestInvalidAllocationStrategyRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}

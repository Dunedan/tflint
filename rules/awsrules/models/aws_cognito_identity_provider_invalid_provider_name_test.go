// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCognitoIdentityProviderInvalidProviderNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	provider_name = "	"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCognitoIdentityProviderInvalidProviderNameRule(),
					Message: `provider_name does not match valid pattern ^[\p{L}\p{M}\p{S}\p{N}\p{P}]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	provider_name = "Google"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityProviderInvalidProviderNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}

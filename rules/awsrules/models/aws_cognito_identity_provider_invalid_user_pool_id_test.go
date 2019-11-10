// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCognitoIdentityProviderInvalidUserPoolIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	user_pool_id = "foobar"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCognitoIdentityProviderInvalidUserPoolIDRule(),
					Message: `user_pool_id does not match valid pattern ^[\w-]+_[0-9a-zA-Z]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	user_pool_id = "foo_bar"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityProviderInvalidUserPoolIDRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}

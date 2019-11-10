// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayGatewayResponseInvalidStatusCodeRule checks the pattern is valid
type AwsAPIGatewayGatewayResponseInvalidStatusCodeRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAPIGatewayGatewayResponseInvalidStatusCodeRule returns new rule with default attributes
func NewAwsAPIGatewayGatewayResponseInvalidStatusCodeRule() *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule {
	return &AwsAPIGatewayGatewayResponseInvalidStatusCodeRule{
		resourceType:  "aws_api_gateway_gateway_response",
		attributeName: "status_code",
		pattern:       regexp.MustCompile(`^[1-5]\d\d$`),
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Name() string {
	return "aws_api_gateway_gateway_response_invalid_status_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`status_code does not match valid pattern ^[1-5]\d\d$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

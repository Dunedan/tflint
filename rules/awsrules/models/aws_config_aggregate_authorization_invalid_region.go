// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsConfigAggregateAuthorizationInvalidRegionRule checks the pattern is valid
type AwsConfigAggregateAuthorizationInvalidRegionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigAggregateAuthorizationInvalidRegionRule returns new rule with default attributes
func NewAwsConfigAggregateAuthorizationInvalidRegionRule() *AwsConfigAggregateAuthorizationInvalidRegionRule {
	return &AwsConfigAggregateAuthorizationInvalidRegionRule{
		resourceType:  "aws_config_aggregate_authorization",
		attributeName: "region",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Name() string {
	return "aws_config_aggregate_authorization_invalid_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"region must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"region must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

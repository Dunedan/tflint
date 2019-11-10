// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSwfDomainInvalidNameRule checks the pattern is valid
type AwsSwfDomainInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSwfDomainInvalidNameRule returns new rule with default attributes
func NewAwsSwfDomainInvalidNameRule() *AwsSwfDomainInvalidNameRule {
	return &AwsSwfDomainInvalidNameRule{
		resourceType:  "aws_swf_domain",
		attributeName: "name",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSwfDomainInvalidNameRule) Name() string {
	return "aws_swf_domain_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSwfDomainInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSwfDomainInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSwfDomainInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSwfDomainInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

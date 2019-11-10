// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationCustomRuleInvalidTagValueScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidTagValueScopeRule() *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule {
	return &AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "tag_value_scope",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_tag_value_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidTagValueScopeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"tag_value_scope must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"tag_value_scope must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmActivationInvalidNameRule checks the pattern is valid
type AwsSsmActivationInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSsmActivationInvalidNameRule returns new rule with default attributes
func NewAwsSsmActivationInvalidNameRule() *AwsSsmActivationInvalidNameRule {
	return &AwsSsmActivationInvalidNameRule{
		resourceType:  "aws_ssm_activation",
		attributeName: "name",
		max:           256,
		pattern:       regexp.MustCompile(`^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmActivationInvalidNameRule) Name() string {
	return "aws_ssm_activation_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmActivationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmActivationInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmActivationInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmActivationInvalidNameRule) Check(runner *tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

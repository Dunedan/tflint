// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLightsailKeyPairInvalidNameRule checks the pattern is valid
type AwsLightsailKeyPairInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailKeyPairInvalidNameRule returns new rule with default attributes
func NewAwsLightsailKeyPairInvalidNameRule() *AwsLightsailKeyPairInvalidNameRule {
	return &AwsLightsailKeyPairInvalidNameRule{
		resourceType:  "aws_lightsail_key_pair",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^\w[\w\-]*\w$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailKeyPairInvalidNameRule) Name() string {
	return "aws_lightsail_key_pair_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailKeyPairInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailKeyPairInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailKeyPairInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailKeyPairInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^\w[\w\-]*\w$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

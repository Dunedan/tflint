// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMRolePolicyAttachmentInvalidRoleRule checks the pattern is valid
type AwsIAMRolePolicyAttachmentInvalidRoleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMRolePolicyAttachmentInvalidRoleRule returns new rule with default attributes
func NewAwsIAMRolePolicyAttachmentInvalidRoleRule() *AwsIAMRolePolicyAttachmentInvalidRoleRule {
	return &AwsIAMRolePolicyAttachmentInvalidRoleRule{
		resourceType:  "aws_iam_role_policy_attachment",
		attributeName: "role",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMRolePolicyAttachmentInvalidRoleRule) Name() string {
	return "aws_iam_role_policy_attachment_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRolePolicyAttachmentInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRolePolicyAttachmentInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRolePolicyAttachmentInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRolePolicyAttachmentInvalidRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`role does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

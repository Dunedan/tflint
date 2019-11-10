// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsBackupSelectionInvalidNameRule checks the pattern is valid
type AwsBackupSelectionInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsBackupSelectionInvalidNameRule returns new rule with default attributes
func NewAwsBackupSelectionInvalidNameRule() *AwsBackupSelectionInvalidNameRule {
	return &AwsBackupSelectionInvalidNameRule{
		resourceType:  "aws_backup_selection",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9\-\_\.]{1,50}$`),
	}
}

// Name returns the rule name
func (r *AwsBackupSelectionInvalidNameRule) Name() string {
	return "aws_backup_selection_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBackupSelectionInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBackupSelectionInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBackupSelectionInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBackupSelectionInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-zA-Z0-9\-\_\.]{1,50}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

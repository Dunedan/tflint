// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFsxLustreFileSystemInvalidImportPathRule checks the pattern is valid
type AwsFsxLustreFileSystemInvalidImportPathRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsFsxLustreFileSystemInvalidImportPathRule returns new rule with default attributes
func NewAwsFsxLustreFileSystemInvalidImportPathRule() *AwsFsxLustreFileSystemInvalidImportPathRule {
	return &AwsFsxLustreFileSystemInvalidImportPathRule{
		resourceType:  "aws_fsx_lustre_file_system",
		attributeName: "import_path",
		max:           900,
		min:           3,
	}
}

// Name returns the rule name
func (r *AwsFsxLustreFileSystemInvalidImportPathRule) Name() string {
	return "aws_fsx_lustre_file_system_invalid_import_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxLustreFileSystemInvalidImportPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxLustreFileSystemInvalidImportPathRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxLustreFileSystemInvalidImportPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxLustreFileSystemInvalidImportPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"import_path must be 900 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"import_path must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

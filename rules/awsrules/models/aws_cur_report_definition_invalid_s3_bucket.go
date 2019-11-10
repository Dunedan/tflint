// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCurReportDefinitionInvalidS3BucketRule checks the pattern is valid
type AwsCurReportDefinitionInvalidS3BucketRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCurReportDefinitionInvalidS3BucketRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidS3BucketRule() *AwsCurReportDefinitionInvalidS3BucketRule {
	return &AwsCurReportDefinitionInvalidS3BucketRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "s3_bucket",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidS3BucketRule) Name() string {
	return "aws_cur_report_definition_invalid_s3_bucket"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidS3BucketRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidS3BucketRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidS3BucketRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidS3BucketRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"s3_bucket must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

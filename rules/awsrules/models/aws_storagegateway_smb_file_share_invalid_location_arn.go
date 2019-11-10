// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidLocationArnRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidLocationArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewaySmbFileShareInvalidLocationArnRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidLocationArnRule() *AwsStoragegatewaySmbFileShareInvalidLocationArnRule {
	return &AwsStoragegatewaySmbFileShareInvalidLocationArnRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "location_arn",
		max:           310,
		min:           16,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidLocationArnRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_location_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidLocationArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidLocationArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidLocationArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidLocationArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"location_arn must be 310 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"location_arn must be 16 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

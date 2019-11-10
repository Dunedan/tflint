// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsMacieS3BucketAssociationInvalidBucketNameRule checks the pattern is valid
type AwsMacieS3BucketAssociationInvalidBucketNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsMacieS3BucketAssociationInvalidBucketNameRule returns new rule with default attributes
func NewAwsMacieS3BucketAssociationInvalidBucketNameRule() *AwsMacieS3BucketAssociationInvalidBucketNameRule {
	return &AwsMacieS3BucketAssociationInvalidBucketNameRule{
		resourceType:  "aws_macie_s3_bucket_association",
		attributeName: "bucket_name",
		max:           500,
	}
}

// Name returns the rule name
func (r *AwsMacieS3BucketAssociationInvalidBucketNameRule) Name() string {
	return "aws_macie_s3_bucket_association_invalid_bucket_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacieS3BucketAssociationInvalidBucketNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacieS3BucketAssociationInvalidBucketNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacieS3BucketAssociationInvalidBucketNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacieS3BucketAssociationInvalidBucketNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"bucket_name must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

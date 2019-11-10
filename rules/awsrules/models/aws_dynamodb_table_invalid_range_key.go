// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDynamoDBTableInvalidRangeKeyRule checks the pattern is valid
type AwsDynamoDBTableInvalidRangeKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBTableInvalidRangeKeyRule returns new rule with default attributes
func NewAwsDynamoDBTableInvalidRangeKeyRule() *AwsDynamoDBTableInvalidRangeKeyRule {
	return &AwsDynamoDBTableInvalidRangeKeyRule{
		resourceType:  "aws_dynamodb_table",
		attributeName: "range_key",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableInvalidRangeKeyRule) Name() string {
	return "aws_dynamodb_table_invalid_range_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableInvalidRangeKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableInvalidRangeKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableInvalidRangeKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableInvalidRangeKeyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"range_key must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"range_key must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFlowLogInvalidLogDestinationTypeRule checks the pattern is valid
type AwsFlowLogInvalidLogDestinationTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFlowLogInvalidLogDestinationTypeRule returns new rule with default attributes
func NewAwsFlowLogInvalidLogDestinationTypeRule() *AwsFlowLogInvalidLogDestinationTypeRule {
	return &AwsFlowLogInvalidLogDestinationTypeRule{
		resourceType:  "aws_flow_log",
		attributeName: "log_destination_type",
		enum: []string{
			"cloud-watch-logs",
			"s3",
		},
	}
}

// Name returns the rule name
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Name() string {
	return "aws_flow_log_invalid_log_destination_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`log_destination_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

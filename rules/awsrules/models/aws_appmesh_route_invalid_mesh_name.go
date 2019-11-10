// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAppmeshRouteInvalidMeshNameRule checks the pattern is valid
type AwsAppmeshRouteInvalidMeshNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshRouteInvalidMeshNameRule returns new rule with default attributes
func NewAwsAppmeshRouteInvalidMeshNameRule() *AwsAppmeshRouteInvalidMeshNameRule {
	return &AwsAppmeshRouteInvalidMeshNameRule{
		resourceType:  "aws_appmesh_route",
		attributeName: "mesh_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshRouteInvalidMeshNameRule) Name() string {
	return "aws_appmesh_route_invalid_mesh_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshRouteInvalidMeshNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshRouteInvalidMeshNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshRouteInvalidMeshNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshRouteInvalidMeshNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"mesh_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"mesh_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

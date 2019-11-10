// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule checks the pattern is valid
type AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule returns new rule with default attributes
func NewAwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule() *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule {
	return &AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule{
		resourceType:  "aws_ec2_transit_gateway",
		attributeName: "default_route_table_association",
		enum: []string{
			"enable",
			"disable",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Name() string {
	return "aws_ec2_transit_gateway_invalid_default_route_table_association"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Check(runner *tflint.Runner) error {
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
					`default_route_table_association is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

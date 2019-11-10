// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCodedeployDeploymentConfigInvalidComputePlatformRule checks the pattern is valid
type AwsCodedeployDeploymentConfigInvalidComputePlatformRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodedeployDeploymentConfigInvalidComputePlatformRule returns new rule with default attributes
func NewAwsCodedeployDeploymentConfigInvalidComputePlatformRule() *AwsCodedeployDeploymentConfigInvalidComputePlatformRule {
	return &AwsCodedeployDeploymentConfigInvalidComputePlatformRule{
		resourceType:  "aws_codedeploy_deployment_config",
		attributeName: "compute_platform",
		enum: []string{
			"Server",
			"Lambda",
			"ECS",
		},
	}
}

// Name returns the rule name
func (r *AwsCodedeployDeploymentConfigInvalidComputePlatformRule) Name() string {
	return "aws_codedeploy_deployment_config_invalid_compute_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployDeploymentConfigInvalidComputePlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployDeploymentConfigInvalidComputePlatformRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployDeploymentConfigInvalidComputePlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployDeploymentConfigInvalidComputePlatformRule) Check(runner *tflint.Runner) error {
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
					`compute_platform is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

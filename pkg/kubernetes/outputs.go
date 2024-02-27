package kubernetes

import (
	"context"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	"github.com/pkg/errors"
	code2cloudv1deploycepstackk8smodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/customendpoint/stack/kubernetes/model"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/org"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/stack/output/backend"
)

func Outputs(ctx context.Context, input *code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackInput) (
	*code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackOutputs, error) {
	pulumiOrgName, err := org.GetOrgName()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pulumi org name")
	}
	stackOutput, err := backend.StackOutput(pulumiOrgName, input.StackJob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stack output")
	}
	return OutputMapTransformer(stackOutput, input), nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackInput) *code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackOutputs{}
	}
	return &code2cloudv1deploycepstackk8smodel.CustomEndpointKubernetesStackOutputs{}
}

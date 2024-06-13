package occmundialoccecrpattern

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type IEcrProps interface {
	ImageName() *string
	Principals() *[]awsiam.IPrincipal
	ScanOnPush() *bool
}

// The jsii proxy for IEcrProps
type jsiiProxy_IEcrProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IEcrProps) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcrProps) Principals() *[]awsiam.IPrincipal {
	var returns *[]awsiam.IPrincipal
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcrProps) ScanOnPush() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanOnPush",
		&returns,
	)
	return returns
}


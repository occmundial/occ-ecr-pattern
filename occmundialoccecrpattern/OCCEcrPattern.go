package occmundialoccecrpattern

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/occmundial/occ-ecr-pattern/occmundialoccecrpattern/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/occmundial/occ-ecr-pattern/occmundialoccecrpattern/internal"
)

type OCCEcrPattern interface {
	constructs.Construct
	Ecr() awsecr.IRepository
	SetEcr(val awsecr.IRepository)
	EcrArn() *string
	SetEcrArn(val *string)
	EcrImageName() *string
	SetEcrImageName(val *string)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for OCCEcrPattern
type jsiiProxy_OCCEcrPattern struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_OCCEcrPattern) Ecr() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"ecr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OCCEcrPattern) EcrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OCCEcrPattern) EcrImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OCCEcrPattern) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewOCCEcrPattern(scope constructs.Construct, id *string, props IEcrProps) OCCEcrPattern {
	_init_.Initialize()

	if err := validateNewOCCEcrPatternParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OCCEcrPattern{}

	_jsii_.Create(
		"@occmundial/occ-ecr-pattern.OCCEcrPattern",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOCCEcrPattern_Override(o OCCEcrPattern, scope constructs.Construct, id *string, props IEcrProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@occmundial/occ-ecr-pattern.OCCEcrPattern",
		[]interface{}{scope, id, props},
		o,
	)
}

func (j *jsiiProxy_OCCEcrPattern)SetEcr(val awsecr.IRepository) {
	if err := j.validateSetEcrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecr",
		val,
	)
}

func (j *jsiiProxy_OCCEcrPattern)SetEcrArn(val *string) {
	if err := j.validateSetEcrArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecrArn",
		val,
	)
}

func (j *jsiiProxy_OCCEcrPattern)SetEcrImageName(val *string) {
	if err := j.validateSetEcrImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecrImageName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OCCEcrPattern_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOCCEcrPattern_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@occmundial/occ-ecr-pattern.OCCEcrPattern",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OCCEcrPattern) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


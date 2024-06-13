// OCC Pattern to create an AWS ECR based on OCC way
package occmundialoccecrpattern

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@occmundial/occ-ecr-pattern.IEcrProps",
		reflect.TypeOf((*IEcrProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "principals", GoGetter: "Principals"},
			_jsii_.MemberProperty{JsiiProperty: "scanOnPush", GoGetter: "ScanOnPush"},
		},
		func() interface{} {
			return &jsiiProxy_IEcrProps{}
		},
	)
	_jsii_.RegisterClass(
		"@occmundial/occ-ecr-pattern.OCCEcrPattern",
		reflect.TypeOf((*OCCEcrPattern)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ecr", GoGetter: "Ecr"},
			_jsii_.MemberProperty{JsiiProperty: "ecrArn", GoGetter: "EcrArn"},
			_jsii_.MemberProperty{JsiiProperty: "ecrImageName", GoGetter: "EcrImageName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OCCEcrPattern{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}

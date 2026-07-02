/*
https://medium.com/@kyodo-tech/designing-go-packages-effectively-00701545785d
*/

package Inspect

import (
	"fmt"
	"reflect"

	Error "github.com/go-composites/error/src"
	MethodNotImplemented "github.com/go-composites/error/src/method_not_implemented"
	NullError "github.com/go-composites/error/src/null"
	Null "github.com/go-composites/null/src"
	Result "github.com/go-composites/result/src"
	String "github.com/go-composites/string/src"
)

// firstable is the one-method structural view of any collection (e.g.
// Array.Interface) that Inspect needs from String.Split's payload. Asserting
// to this local interface instead of importing the array package breaks the
// array → boolean → inspect → array import cycle without changing behavior.
type firstable interface {
	First() Result.Interface
}

type Interface interface {
	HasError() bool
	Type() String.Interface
	Addr() String.Interface
	Data() Result.Interface
	ToGoString() string
}

type data struct {
	error   Error.Interface
	objType String.Interface
	objAddr String.Interface
	objData interface{}
}

func New(obj interface{}) Interface {
	var result = &data{
		error:   NullError.New(),
		objType: String.New(),
		objAddr: String.New(),
		objData: Null.New(),
	}

	// reflect.TypeOf(obj).String() is a dotted, package-qualified type name
	// (e.g. "Boolean.data"); its first "."-separated segment is the package
	// component. String.Split always yields at least one element and First()
	// always succeeds on that non-empty collection, so the unwrap is total.
	stringSplitResult := String.New(
		String.WithGoString(reflect.TypeOf(obj).String()),
	).Split(`.`)
	firstResult := stringSplitResult.Payload().(firstable).First()
	objType := firstResult.Payload().(String.Interface).ToGoString()
	result.objType = String.New(
		String.WithGoString(objType),
	)
	result.objAddr = String.New(
		String.WithGoString(
			fmt.Sprintf("%p", &obj),
		),
	)
	result.objData = obj
	return result
}

func (d data) HasError() bool {
	return !d.error.IsNull()
}

func (d data) Type() String.Interface {
	return d.objType
}
func (d data) Addr() String.Interface {
	return d.objAddr
}
func (d data) Data() Result.Interface {
	if d.HasError() {
		return Result.New(
			Result.WithError(
				d.error,
			),
		)
	} else {
		methodName := `InspectData`
		obj := reflect.ValueOf(d.objData)
		method := obj.MethodByName(methodName)
		if method.IsValid() {
			result := method.Call(nil)
			if len(result) > 0 {
				return Result.New(
					Result.WithPayload(
						result[0].String(),
					),
				)
			} else {
				return Result.New(
					Result.WithPayload(
						Null.New(),
					),
				)
			}
		} else {
			return Result.New(
				Result.WithError(
					MethodNotImplemented.New(methodName),
				),
			)
		}
	}
}
func (d data) ToGoString() string {
	result := d.Data()
	if result.HasError() {
		message := result.Error().Message()
		panic(message)
	} else {
		payload := result.Payload()
		return fmt.Sprintf(
			"<Object type=%s addr=%s data={ %s } >\n",
			d.Type().ToGoString(),
			d.Addr().ToGoString(),
			payload,
		)
	}

}

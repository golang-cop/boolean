/*
https://medium.com/@kyodo-tech/designing-go-packages-effectively-00701545785d
*/

package Inspect

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	Array "github.com/golang-cop/array/src"
	Error "github.com/golang-cop/error/src"
	MethodNotImplemented "github.com/golang-cop/error/src/method_not_implemented"
	NullError "github.com/golang-cop/error/src/null"
	Null "github.com/golang-cop/null/src"
	Result "github.com/golang-cop/result/src"
	String "github.com/golang-cop/string/src"
)

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

	stringSplitResult := String.New(
		String.WithGoString(reflect.TypeOf(obj).String()),
	).Split(`.`)
	//spew.Dump(stringSplitResult, stringSplitResult.HasError())
	//panic(`BAM`)
	if stringSplitResult.HasError() {
		result.error = stringSplitResult.Error()
	} else {
		interfaceTypeResult := stringSplitResult.Payload()
		firstResult := interfaceTypeResult.(Array.Interface).First()
		if firstResult.HasError() {
			result.error = firstResult.Error()
		} else {
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
		}
	}
	spew.Dump(result)
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
	//spew.Dump(result)
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

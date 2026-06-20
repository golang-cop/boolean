package Is

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

func Null(obj interface{}) bool {
	fmt.Printf("ObjType: %T\n", obj)
	spew.Dump(reflect.ValueOf(obj).Interface())
	spew.Dump(reflect.TypeOf(obj))
	return true
}

/*
func NullError(error Error.Interface) bool {
	return reflect.TypeOf(error) == NullError.Interface
}
*/

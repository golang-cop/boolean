package Boolean

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	Inspect "github.com/golang-oop/boolean/src/inspect"
)

type Interface interface {
	ToGoBool() bool
	ToGoString() string
	IsTrue() bool
	IsFalse() bool
	Equal(Interface) Interface
	Inspect() Inspect.Interface
}
type data struct {
	value bool
}

/*
Given a b Go boolean

Return a new Boolean.Interface with a data.value set to b.
*/
func New(b bool) Interface {
	return &data{
		value: b,
	}
}

/*
Return a new Boolean.Interface with a data.value set to true.
*/
func True() Interface {
	return New(true)
}

/*
Return a new Boolean.Interface with a data.value set to false.
*/
func False() Interface {
	return New(false)
}

/*
Return the data.value of a Boolean.Interface.
*/
func (d data) ToGoBool() bool {
	return d.value
}

/*
Return the string representing the data.value of a Boolean.Interface.
*/
func (d data) ToGoString() string {
	return fmt.Sprintf("\"%t\"", d.value)
}

/*
Test if the data.value is true.
*/
func (d data) IsTrue() bool {
	return d.value
}

/*
Test if the data.value is false.
*/
func (d data) IsFalse() bool {
	return !d.value
}

/*
Test if the data.value of the given Boolean.Interface
is equal to the receiver one.
*/
func (d data) Equal(b Interface) Interface {
	return New(d.ToGoBool() == b.ToGoBool())
}

/*
Return a string representing a Boolean.Interface
with its address and its value.
*/
type InspectStruct struct {
	Type string
	Addr string
	Data string
}

/*
Return an Inspect.Interface
*/
func (d data) Inspect() Inspect.Interface {
	//return fmt.Sprintf(
	//	"<Object type=Boolean addr=%p data={ value=%t }>", &d, d.value,
	//)
	inspect := Inspect.New(d)
	fmt.Println(`>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>`)
	spew.Dump(inspect)
	fmt.Println(`<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<`)
	return inspect
}

func (d data) InspectData() string {
	return fmt.Sprintf("value=%t", d.value)
}

/*
func (d data) Methods() string {
	return ``
}
*/

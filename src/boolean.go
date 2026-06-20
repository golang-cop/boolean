package Boolean

import (
	"fmt"

	Inspect "github.com/go-composites/boolean/src/inspect"
)

type Interface interface {
	ToGoBool() bool
	ToGoString() string
	IsTrue() bool
	IsFalse() bool
	Equal(Interface) Interface
	And(Interface) Interface
	Or(Interface) Interface
	Not() Interface
	Xor(Interface) Interface
	IsNull() bool
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
Return the logical AND of the receiver and the given Boolean.Interface.
*/
func (d data) And(b Interface) Interface {
	return New(d.ToGoBool() && b.ToGoBool())
}

/*
Return the logical OR of the receiver and the given Boolean.Interface.
*/
func (d data) Or(b Interface) Interface {
	return New(d.ToGoBool() || b.ToGoBool())
}

/*
Return the logical negation of the receiver.
*/
func (d data) Not() Interface {
	return New(!d.ToGoBool())
}

/*
Return the exclusive OR of the receiver and the given Boolean.Interface.
*/
func (d data) Xor(b Interface) Interface {
	return New(d.ToGoBool() != b.ToGoBool())
}

/*
IsNull reports that this is a real (non-null) Boolean.
*/
func (d data) IsNull() bool {
	return false
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
	return Inspect.New(d)
}

func (d data) InspectData() string {
	return fmt.Sprintf("value=%t", d.value)
}

// null is the Null-Object variant of a Boolean: a placeholder that honours the
// full Interface without ever being nil. A null boolean is neither true nor
// false; its logical operators all return the null boolean.
type null struct{}

/*
Null returns the Null-Object Boolean.
*/
func Null() Interface {
	return &null{}
}

// ToGoBool returns false for the null Boolean.
func (n null) ToGoBool() bool { return false }

// ToGoString renders the null Boolean as the literal "null".
func (n null) ToGoString() string { return "\"null\"" }

// IsTrue is false: a null boolean is not true.
func (n null) IsTrue() bool { return false }

// IsFalse is false: a null boolean is not false.
func (n null) IsFalse() bool { return false }

// Equal returns the null Boolean.
func (n null) Equal(b Interface) Interface { return Null() }

// And returns the null Boolean.
func (n null) And(b Interface) Interface { return Null() }

// Or returns the null Boolean.
func (n null) Or(b Interface) Interface { return Null() }

// Not returns the null Boolean.
func (n null) Not() Interface { return Null() }

// Xor returns the null Boolean.
func (n null) Xor(b Interface) Interface { return Null() }

// IsNull reports that this is the null Boolean.
func (n null) IsNull() bool { return true }

// InspectData renders the null Boolean's data.
func (n null) InspectData() string { return "value=null" }

// Inspect returns an Inspect.Interface for the null Boolean.
func (n null) Inspect() Inspect.Interface {
	return Inspect.New(n)
}

package Inspect

import (
	"strings"
	"testing"

	Error "github.com/go-composites/error/src"
	String "github.com/go-composites/string/src"
)

// withData is an inspectable value: it implements the InspectData() string
// method that Data() reflects for, exercising the len(result) > 0 branch.
type withData struct{}

func (withData) InspectData() string { return "value=42" }

// noReturn has an InspectData method that returns nothing, exercising the
// len(result) == 0 branch (payload becomes the Null object).
type noReturn struct{}

func (noReturn) InspectData() {}

// noMethod has no InspectData method at all, exercising the
// method-not-implemented branch (which makes Data() carry an error).
type noMethod struct{}

func TestNewInspectsType(t *testing.T) {
	insp := New(withData{})

	if insp.HasError() {
		t.Fatalf("HasError() = true, want false for a well-formed value")
	}
	// reflect.TypeOf(withData{}).String() == "Inspect.withData"; the first
	// dotted segment is the package component.
	if got := insp.Type().ToGoString(); got != "Inspect" {
		t.Fatalf("Type() = %q, want %q", got, "Inspect")
	}
	if got := insp.Addr().ToGoString(); !strings.HasPrefix(got, "0x") {
		t.Fatalf("Addr() = %q, want a 0x… pointer literal", got)
	}
}

func TestDataWithInspectDataPayload(t *testing.T) {
	res := New(withData{}).Data()

	if res.HasError() {
		t.Fatalf("Data().HasError() = true, want false")
	}
	if got := res.Payload(); got != "value=42" {
		t.Fatalf("Data().Payload() = %v, want %q", got, "value=42")
	}
}

func TestDataWithVoidInspectData(t *testing.T) {
	res := New(noReturn{}).Data()

	if res.HasError() {
		t.Fatalf("Data().HasError() = true, want false")
	}
	if res.Payload() == nil {
		t.Fatalf("Data().Payload() = nil, want the Null object")
	}
}

func TestDataWithoutInspectData(t *testing.T) {
	res := New(noMethod{}).Data()

	if !res.HasError() {
		t.Fatalf("Data().HasError() = false, want true for a value without InspectData")
	}
	if !strings.Contains(res.Error().Message(), "InspectData") {
		t.Fatalf("error message = %q, want it to mention InspectData", res.Error().Message())
	}
}

func TestDataWhenReceiverCarriesError(t *testing.T) {
	d := data{
		error:   Error.New("boom"),
		objType: String.New(),
		objAddr: String.New(),
		objData: nil,
	}

	if !d.HasError() {
		t.Fatalf("HasError() = false, want true when the receiver carries an error")
	}
	res := d.Data()
	if !res.HasError() {
		t.Fatalf("Data().HasError() = false, want true")
	}
	if got := res.Error().Message(); got != "boom" {
		t.Fatalf("Data().Error().Message() = %q, want %q", got, "boom")
	}
}

func TestToGoStringRendersObject(t *testing.T) {
	out := New(withData{}).ToGoString()

	for _, want := range []string{"type=Inspect", "addr=0x", "value=42"} {
		if !strings.Contains(out, want) {
			t.Fatalf("ToGoString() = %q, want it to contain %q", out, want)
		}
	}
}

func TestToGoStringPanicsOnError(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("ToGoString() did not panic on an errored inspection")
		}
		if msg, ok := r.(string); !ok || !strings.Contains(msg, "InspectData") {
			t.Fatalf("recovered %v, want a panic mentioning InspectData", r)
		}
	}()
	_ = New(noMethod{}).ToGoString()
}

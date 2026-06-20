package Is_test

import (
	"testing"

	Is "github.com/go-composites/boolean/src/is"
)

func TestNull(t *testing.T) {
	if !Is.Null(42) {
		t.Fatal("Null should report true")
	}
}

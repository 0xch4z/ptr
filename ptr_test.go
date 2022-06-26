package ptr

import (
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	var ref *int

	if ValueFrom(ref) != 0 {
		t.Error("expected value for nil int pointer to be 0")
	}

	value := 1
	ref = To(value)
	if ValueFrom(ref) != value {
		t.Errorf("expected value for int pointer to %d to be %[1]d", value)
	}
}

func TestTo(t *testing.T) {
	n := 5
	if !reflect.DeepEqual(To(n), &n) {
		t.Errorf("expected to point to %v", &n)
	}
}

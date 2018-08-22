package slicepointer

import (
	"reflect"
	"testing"
)

func TestAddOne(t *testing.T) {
	n := []int{1, 2, 3}
	AddOne(n)
	expected := []int{2, 3, 4}
	if reflect.DeepEqual(expected, n) == false {
		t.Log("expected:", expected)
		t.Error("found:", n)
	}
}

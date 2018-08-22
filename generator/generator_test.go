package generator

import "testing"

func Test_generator(t *testing.T) {
	expected := 4
	gen := Newintgenerator()
	gen()           //1
	gen()           //2
	gen()           //3
	result := gen() //4
	if expected != result {
		t.Error("expected:", expected)
		t.Log("found:", result)
	}
}

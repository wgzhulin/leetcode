package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestVerify(t *testing.T) {
	type Input struct {
		A int
		B int
	}
	i := Input{
		A: 1,
		B: 2,
	}

	type Except struct {
		E int
	}
	e := Except{
		E: 2,
	}
	Verify(t, Max, i, e)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Verify(t *testing.T, exec, input, except interface{}) {
	execValue := reflect.ValueOf(exec)
	inputValue := reflect.ValueOf(input)
	numField := inputValue.NumField()
	in := make([]reflect.Value, numField)
	for i := range in {
		in[i] = inputValue.Field(i)
	}

	callResult := execValue.Call(in)
	exceptValue := reflect.ValueOf(except)
	if len(callResult) != exceptValue.NumField() {
		fmt.Println("result num not equal")
		return
	}
	for i := range callResult {
		s1 := fmt.Sprint(callResult[i])
		s2 := fmt.Sprint(exceptValue.Field(i))
		if !strings.EqualFold(s1, s2) {
			t.Fatalf("not equal")
		}
	}
}

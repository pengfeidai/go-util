package main

import (
	"reflect"
	"slice/reverse"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []interface{}{1, 2, "a", 3, "b", "c"}
	reverse.Reverse(s)
	data := []interface{}{"c", "b", 3, "a", 2, 1}
	if !reflect.DeepEqual(s, data) {
		t.Error("failed")
	}
}

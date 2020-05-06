package common

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %+v does not equal to actual: %+v", expected, actual)
	}
}

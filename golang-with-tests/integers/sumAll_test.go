package integers

import (
	"reflect"
	"testing"
)

func TestAddAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}

package integers

import (
	"reflect"
	"testing"
)

func TestSumAllTais(t *testing.T) {
	checkSlices := func(result []int, expected []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, result)
		}
	}

	t.Run("sums two sized slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSlices(result, expected, t)
	})

	t.Run("sums various sized slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9}, []int{10, 11, 12})
		expected := []int{2, 9, 23}

		checkSlices(result, expected, t)
	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{2, 22, 222})
		expected := []int{0, 244}

		checkSlices(result, expected, t)
	})
}

package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("sum a slice", func(t *testing.T) {
		numbers := []int{8, 9, 5}

		sum := Add(numbers)
		expect := 22

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given %v", expect, sum, numbers)
		}
	})
}

func ExampleAdd() {
	sum := Add([]int{5, 8, 9, 7, 1})
	fmt.Print(sum)
	// Output: 30
}

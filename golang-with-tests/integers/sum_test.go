package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum a slice", func(t *testing.T) {
		numbers := []int{8, 9, 5}

		sum := Sum(numbers)
		expect := 22

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given %v", expect, sum, numbers)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{5, 8, 9, 7, 1})
	fmt.Print(sum)
	// Output: 30
}

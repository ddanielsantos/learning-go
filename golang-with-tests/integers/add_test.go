package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 4)
	expect := 6

	if sum != expect {
		t.Errorf("expected '%d' but got '%d'", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(5, 76)
	fmt.Print(sum)
	// Output: 81
}

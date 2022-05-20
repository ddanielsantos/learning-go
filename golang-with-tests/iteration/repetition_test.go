package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertNumber := func(t *testing.T, result, expected int8) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: '%d' Expected: '%d'", result, expected)
		}
	}

	t.Run("repeat a string 5 times", func(t *testing.T) {
		result := Repeat("a", 5)
		strLen := strings.Count(result, "a")

		assertNumber(t, int8(strLen), 5)
	})

	t.Run("repeat a string 10 times", func(t *testing.T) {
		result := Repeat("a", 10)
		strLen := strings.Count(result, "a")

		assertNumber(t, int8(strLen), 10)
	})
}

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Print(str)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

package iteration

import (
	"fmt"
	"testing"
)

func Test_Loop(t *testing.T) {
	t.Run("repeat character, given amount of time", func(t *testing.T) {

		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected %q, got %q", expected, repeated)
		}
	})

	t.Run("negative repeat count prints it 5 times", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q, got %q", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	fmt.Println(Repeat("x", 3))
	// Output: xxx
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", -1)
	}
}

package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 7)
	fmt.Println(repeat)
	// Output: bbbbbbb
}

package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// b.N: The Go testing framework automatically determines how many times to run the loop (b.N) to get a statistically significant measurement.
	for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	println(repeated)
	// Output: aaaaa
}
package math

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	p := math.Round(10.1)
	fmt.Printf("%.1f\n", p)

	got := int64(p)
	want := int64(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	p = math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	got = int64(p)
	want = int64(11)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}


	p = math.Round(-10.1)
	fmt.Printf("%.1f\n", p)

	want = int64(p)
	got = int64(-10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	p = math.Round(-10.5)
	fmt.Printf("%.1f\n", p)

	want = int64(p)
	got = int64(-11)

	if got != want {
		t.Errorf("got %d want %d", got , want)
	}
}

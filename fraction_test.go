package fraction

import (
	"fmt"
	"math"
	"testing"
)

var fractionTest = []struct {
	v   float64
	d   int64
	num int64
	den int64
}{
	{0.33, 10, 1, 3},
	{0.33, -1, 1, 3},
	{-0.33, 10, -1, 3},
}

var wholeFractionTest = []struct {
	v     float64
	d     int64
	whole int64
	num   int64
	den   int64
}{
	{2.33, 10, 2, 1, 3},
	{2.33, -1, 2, 1, 3},
	{-2.33, 10, -2, 1, 3},
	{-0.33, 10, 0, -1, 3},
}

func TestFraction(t *testing.T) {
	for i, tt := range fractionTest {
		num, den, e := Fraction(tt.v, tt.d)
		if num != tt.num {
			t.Errorf("%d. Fraction num = %v, want %v", i, num, tt.num)
		}
		if den != tt.den {
			t.Errorf("%d. Fraction den = %v, want %v", i, den, tt.den)
		}
		if abs := math.Abs(e); abs > 0.01 {
			t.Errorf("%d. Fraction math.Abs(%v) = %v > 0.01", i, e, abs)
		}
	}
}

func TestWholeFraction(t *testing.T) {
	for i, tt := range wholeFractionTest {
		whole, num, den, e := WholeFraction(tt.v, tt.d)
		if whole != tt.whole {
			t.Errorf("%d. WholeFraction whole = %v, want %v", i, whole, tt.whole)
		}
		if num != tt.num {
			t.Errorf("%d. WholeFraction num = %v, want %v", i, num, tt.num)
		}
		if den != tt.den {
			t.Errorf("%d. WholeFraction den = %v, want %v", i, den, tt.den)
		}
		if abs := math.Abs(e); abs > 0.01 {
			t.Errorf("%d. WholeFraction math.Abs(%v) = %v > 0.01", i, e, abs)
		}
	}
}

func ExampleFraction() {
	num, den, e := Fraction(0.33, -1)
	fmt.Println(num)
	fmt.Println(den)
	fmt.Println(e)
	// Output:
	// 1
	// 3
	// -0.0033333333333332993
}

func ExampleWholeFraction() {
	whole, num, den, e := WholeFraction(2.5, -1)
	fmt.Println(whole)
	fmt.Println(num)
	fmt.Println(den)
	fmt.Println(e)
	// Output:
	// 2
	// 1
	// 2
	// 0
}

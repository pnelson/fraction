// Package fraction approximates fractions from floats.
package fraction

// Fraction returns the numerator, denominator and error of v, limiting the
// denominator to a maximum of d. If the provided value of d is less than 1, it
// defaults to 10. If v is negative, the numerator will be negative.
func Fraction(v float64, d int64) (int64, int64, float64) {
	sign := int64(1)
	ai := int64(0)
	m := [2][2]int64{{1, 0}, {0, 1}}

	if v < 0 {
		sign = -1
		v *= -1.0
	}

	if d < 1 {
		d = 10
	}

	x := v

	// Find terms until denominator gets too large.
	for m[1][0]*ai+m[1][1] <= d {
		m[0][1], m[0][0] = m[0][0], m[0][0]*ai+m[0][1]
		m[1][1], m[1][0] = m[1][0], m[1][0]*ai+m[1][1]

		// Handle division by 0.
		if x == float64(ai) {
			break
		}

		x = 1 / (x - float64(ai))
		ai = int64(x)

		// Handle representation failure.
		if x > float64(0x7FFFFFFF) {
			break
		}
	}

	return m[0][0] * sign, m[1][0], v - (float64(m[0][0]) / float64(m[1][0]))
}

// WholeFraction returns the whole, numerator, denominator and error of v,
// limiting the denominator to a maximum of d. If the provided value of d is
// less than 1, it defaults to 10. The whole number will take the sign unless
// it is 0, in which case the numerator will.
func WholeFraction(v float64, d int64) (int64, int64, int64, float64) {
	whole := int64(v)
	v -= float64(whole)
	num, den, e := Fraction(v, d)
	if num < 0 && whole != 0 {
		num *= -1
	}
	return whole, num, den, e
}

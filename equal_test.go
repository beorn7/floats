package floats

import (
	"math"
	"testing"
)

// These tests have been copied from the reference implementation at:
// https://floating-point-gui.de/errors/NearlyEqualsTest.java

func TestAlmostEqual_big(t *testing.T) {
	// Regular large numbers - generally not problematic.
	testAlmostEqual(t, true, 1000000, 1000001, 0.00001)
	testAlmostEqual(t, true, 1000001, 1000000, 0.00001)
	testAlmostEqual(t, false, 10000, 10001, 0.00001)
	testAlmostEqual(t, false, 10001, 10000, 0.00001)
}

func TestAlmostEqual_bigNeg(t *testing.T) {
	// Negative large numbers.
	testAlmostEqual(t, true, -1000000, -1000001, 0.00001)
	testAlmostEqual(t, true, -1000001, -1000000, 0.00001)
	testAlmostEqual(t, false, -10000, -10001, 0.00001)
	testAlmostEqual(t, false, -10001, -10000, 0.00001)
}

func TestAlmostEqual_mid(t *testing.T) {
	// Numbers around 1.
	testAlmostEqual(t, true, 1.0000001, 1.0000002, 0.00001)
	testAlmostEqual(t, true, 1.0000002, 1.0000001, 0.00001)
	testAlmostEqual(t, false, 1.0002, 1.0001, 0.00001)
	testAlmostEqual(t, false, 1.0001, 1.0002, 0.00001)
}

func TestAlmostEqual_midNeg(t *testing.T) {
	// Numbers around -1.
	testAlmostEqual(t, true, -1.000001, -1.000002, 0.00001)
	testAlmostEqual(t, true, -1.000002, -1.000001, 0.00001)
	testAlmostEqual(t, false, -1.0001, -1.0002, 0.00001)
	testAlmostEqual(t, false, -1.0002, -1.0001, 0.00001)
}

func TestAlmostEqual_small(t *testing.T) {
	// Numbers between 1 and 0.
	testAlmostEqual(t, true, 0.000000001000001, 0.000000001000002, 0.00001)
	testAlmostEqual(t, true, 0.000000001000002, 0.000000001000001, 0.00001)
	testAlmostEqual(t, false, 0.000000000001002, 0.000000000001001, 0.00001)
	testAlmostEqual(t, false, 0.000000000001001, 0.000000000001002, 0.00001)
}

func TestAlmostEqual_smallNeg(t *testing.T) {
	// Numbers between -1 and 0.
	testAlmostEqual(t, true, -0.000000001000001, -0.000000001000002, 0.00001)
	testAlmostEqual(t, true, -0.000000001000002, -0.000000001000001, 0.00001)
	testAlmostEqual(t, false, -0.000000000001002, -0.000000000001001, 0.00001)
	testAlmostEqual(t, false, -0.000000000001001, -0.000000000001002, 0.00001)
}

func TestAlmostEqual_smallDiffs(t *testing.T) {
	// Small differences away from zero.
	testAlmostEqual(t, true, 0.3, 0.30000003, 0.00001)
	testAlmostEqual(t, true, -0.3, -0.30000003, 0.00001)
}

func TestAlmostEqual_zero(t *testing.T) {
	// Comparisons involving zero.
	testAlmostEqual(t, true, 0.0, 0.0, 0.00001)
	testAlmostEqual(t, true, 0.0, -0.0, 0.00001)
	testAlmostEqual(t, true, -0.0, -0.0, 0.00001)
	testAlmostEqual(t, false, 0.00000001, 0.0, 0.00001)
	testAlmostEqual(t, false, 0.0, 0.00000001, 0.00001)
	testAlmostEqual(t, false, -0.00000001, 0.0, 0.00001)
	testAlmostEqual(t, false, 0.0, -0.00000001, 0.00001)

	testAlmostEqual(t, true, 0.0, 1e-40, 0.01)
	testAlmostEqual(t, true, 1e-40, 0.0, 0.01)
	testAlmostEqual(t, false, 1e-40, 0.0, 0.000001)
	testAlmostEqual(t, false, 0.0, 1e-40, 0.000001)

	testAlmostEqual(t, true, 0.0, -1e-40, 0.1)
	testAlmostEqual(t, true, -1e-40, 0.0, 0.1)
	testAlmostEqual(t, false, -1e-40, 0.0, 0.00000001)
	testAlmostEqual(t, false, 0.0, -1e-40, 0.00000001)
}

func TestAlmostEqual_extremeMax(t *testing.T) {
	// Comparisons involving extreme values (overflow potential).
	testAlmostEqual(t, true, math.MaxFloat64, math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, math.MaxFloat64, -math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, -math.MaxFloat64, math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, math.MaxFloat64, math.MaxFloat64/2, 0.00001)
	testAlmostEqual(t, false, math.MaxFloat64, -math.MaxFloat64/2, 0.00001)
	testAlmostEqual(t, false, -math.MaxFloat64, math.MaxFloat64/2, 0.00001)
}

func TestAlmostEqual_infinities(t *testing.T) {
	// Comparisons involving infinities.
	testAlmostEqual(t, true, math.Inf(1), math.Inf(1), 0.00001)
	testAlmostEqual(t, true, math.Inf(-1), math.Inf(-1), 0.00001)
	testAlmostEqual(t, false, math.Inf(-1), math.Inf(1), 0.00001)
	testAlmostEqual(t, false, math.Inf(1), math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, math.Inf(-1), -math.MaxFloat64, 0.00001)
}

func TestAlmostEqual_nan(t *testing.T) {
	// Comparisons involving NaN values.
	testAlmostEqual(t, false, math.NaN(), math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), 0.0, 0.00001)
	testAlmostEqual(t, false, -0.0, math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), -0.0, 0.00001)
	testAlmostEqual(t, false, 0.0, math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), math.Inf(1), 0.00001)
	testAlmostEqual(t, false, math.Inf(1), math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), math.Inf(-1), 0.00001)
	testAlmostEqual(t, false, math.Inf(-1), math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, math.MaxFloat64, math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), -math.MaxFloat64, 0.00001)
	testAlmostEqual(t, false, -math.MaxFloat64, math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, false, math.SmallestNonzeroFloat64, math.NaN(), 0.00001)
	testAlmostEqual(t, false, math.NaN(), -math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, false, -math.SmallestNonzeroFloat64, math.NaN(), 0.00001)
}

func TestAlmostEqual_opposite(t *testing.T) {
	// Comparisons of numbers on opposite sides of 0
	testAlmostEqual(t, false, 1.000000001, -1.0, 0.00001)
	testAlmostEqual(t, false, -1.0, 1.000000001, 0.00001)
	testAlmostEqual(t, false, -1.000000001, 1.0, 0.00001)
	testAlmostEqual(t, false, 1.0, -1.000000001, 0.00001)
	testAlmostEqual(t, true, 10*math.SmallestNonzeroFloat64, 10*-math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, false, 10000*math.SmallestNonzeroFloat64, 10000*-math.SmallestNonzeroFloat64, 0.00001)
}

func TestAlmostEqual_ulp(t *testing.T) {
	// The really tricky part - comparisons of numbers very close to zero.
	testAlmostEqual(t, true, math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, true, math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, true, -math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, true, math.SmallestNonzeroFloat64, 0, 0.00001)
	testAlmostEqual(t, true, 0, math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, true, -math.SmallestNonzeroFloat64, 0, 0.00001)
	testAlmostEqual(t, true, 0, -math.SmallestNonzeroFloat64, 0.00001)

	testAlmostEqual(t, false, 0.000000001, -math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, false, 0.000000001, math.SmallestNonzeroFloat64, 0.00001)
	testAlmostEqual(t, false, math.SmallestNonzeroFloat64, 0.000000001, 0.00001)
	testAlmostEqual(t, false, -math.SmallestNonzeroFloat64, 0.000000001, 0.00001)
}

func testAlmostEqual(t *testing.T, want bool, a, b, ε float64) {
	if AlmostEqual(a, b, ε) != want {
		t.Errorf("expected `AlmostEqual(%f, %f, %f)` to return %t", a, b, ε, want)
	}
}

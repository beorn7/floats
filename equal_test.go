package floats

import (
	"math"
	"testing"
)

// These tests have been copied from the reference implementation at:
// https://floating-point-gui.de/errors/NearlyEqualsTest.java

func TestAlmostEqual_big(t *testing.T) {
	// Regular large numbers - generally not problematic.
	assertTrue(t, testAlmostEqual(1000000, 1000001))
	assertTrue(t, testAlmostEqual(1000001, 1000000))
	assertFalse(t, testAlmostEqual(10000, 10001))
	assertFalse(t, testAlmostEqual(10001, 10000))
}

func TestAlmostEqual_bigNeg(t *testing.T) {
	// Negative large numbers.
	assertTrue(t, testAlmostEqual(-1000000, -1000001))
	assertTrue(t, testAlmostEqual(-1000001, -1000000))
	assertFalse(t, testAlmostEqual(-10000, -10001))
	assertFalse(t, testAlmostEqual(-10001, -10000))
}

func TestAlmostEqual_mid(t *testing.T) {
	// Numbers around 1.
	assertTrue(t, testAlmostEqual(1.0000001, 1.0000002))
	assertTrue(t, testAlmostEqual(1.0000002, 1.0000001))
	assertFalse(t, testAlmostEqual(1.0002, 1.0001))
	assertFalse(t, testAlmostEqual(1.0001, 1.0002))
}

func TestAlmostEqual_midNeg(t *testing.T) {
	// Numbers around -1.
	assertTrue(t, testAlmostEqual(-1.000001, -1.000002))
	assertTrue(t, testAlmostEqual(-1.000002, -1.000001))
	assertFalse(t, testAlmostEqual(-1.0001, -1.0002))
	assertFalse(t, testAlmostEqual(-1.0002, -1.0001))
}

func TestAlmostEqual_small(t *testing.T) {
	// Numbers between 1 and 0.
	assertTrue(t, testAlmostEqual(0.000000001000001, 0.000000001000002))
	assertTrue(t, testAlmostEqual(0.000000001000002, 0.000000001000001))
	assertFalse(t, testAlmostEqual(0.000000000001002, 0.000000000001001))
	assertFalse(t, testAlmostEqual(0.000000000001001, 0.000000000001002))
}

func TestAlmostEqual_smallNeg(t *testing.T) {
	// Numbers between -1 and 0.
	assertTrue(t, testAlmostEqual(-0.000000001000001, -0.000000001000002))
	assertTrue(t, testAlmostEqual(-0.000000001000002, -0.000000001000001))
	assertFalse(t, testAlmostEqual(-0.000000000001002, -0.000000000001001))
	assertFalse(t, testAlmostEqual(-0.000000000001001, -0.000000000001002))
}

func TestAlmostEqual_smallDiffs(t *testing.T) {
	// Small differences away from zero.
	assertTrue(t, testAlmostEqual(0.3, 0.30000003))
	assertTrue(t, testAlmostEqual(-0.3, -0.30000003))
}

func TestAlmostEqual_zero(t *testing.T) {
	// Comparisons involving zero.
	assertTrue(t, testAlmostEqual(0.0, 0.0))
	assertTrue(t, testAlmostEqual(0.0, -0.0))
	assertTrue(t, testAlmostEqual(-0.0, -0.0))
	assertFalse(t, testAlmostEqual(0.00000001, 0.0))
	assertFalse(t, testAlmostEqual(0.0, 0.00000001))
	assertFalse(t, testAlmostEqual(-0.00000001, 0.0))
	assertFalse(t, testAlmostEqual(0.0, -0.00000001))

	assertTrue(t, AlmostEqual(0.0, 1e-40, 0.01))
	assertTrue(t, AlmostEqual(1e-40, 0.0, 0.01))
	assertFalse(t, AlmostEqual(1e-40, 0.0, 0.000001))
	assertFalse(t, AlmostEqual(0.0, 1e-40, 0.000001))

	assertTrue(t, AlmostEqual(0.0, -1e-40, 0.1))
	assertTrue(t, AlmostEqual(-1e-40, 0.0, 0.1))
	assertFalse(t, AlmostEqual(-1e-40, 0.0, 0.00000001))
	assertFalse(t, AlmostEqual(0.0, -1e-40, 0.00000001))
}

func TestAlmostEqual_extremeMax(t *testing.T) {
	// Comparisons involving extreme values (overflow potential).
	assertTrue(t, testAlmostEqual(math.MaxFloat64, math.MaxFloat64))
	assertFalse(t, testAlmostEqual(math.MaxFloat64, -math.MaxFloat64))
	assertFalse(t, testAlmostEqual(-math.MaxFloat64, math.MaxFloat64))
	assertFalse(t, testAlmostEqual(math.MaxFloat64, math.MaxFloat64/2))
	assertFalse(t, testAlmostEqual(math.MaxFloat64, -math.MaxFloat64/2))
	assertFalse(t, testAlmostEqual(-math.MaxFloat64, math.MaxFloat64/2))
}

func TestAlmostEqual_infinities(t *testing.T) {
	// Comparisons involving infinities.
	assertTrue(t, testAlmostEqual(math.Inf(1), math.Inf(1)))
	assertTrue(t, testAlmostEqual(math.Inf(-1), math.Inf(-1)))
	assertFalse(t, testAlmostEqual(math.Inf(-1), math.Inf(1)))
	assertFalse(t, testAlmostEqual(math.Inf(1), math.MaxFloat64))
	assertFalse(t, testAlmostEqual(math.Inf(-1), -math.MaxFloat64))
}

func TestAlmostEqual_nan(t *testing.T) {
	// Comparisons involving NaN values.
	assertFalse(t, testAlmostEqual(math.NaN(), math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), 0.0))
	assertFalse(t, testAlmostEqual(-0.0, math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), -0.0))
	assertFalse(t, testAlmostEqual(0.0, math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), math.Inf(1)))
	assertFalse(t, testAlmostEqual(math.Inf(1), math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), math.Inf(-1)))
	assertFalse(t, testAlmostEqual(math.Inf(-1), math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), math.MaxFloat64))
	assertFalse(t, testAlmostEqual(math.MaxFloat64, math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), -math.MaxFloat64))
	assertFalse(t, testAlmostEqual(-math.MaxFloat64, math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), math.SmallestNonzeroFloat64))
	assertFalse(t, testAlmostEqual(math.SmallestNonzeroFloat64, math.NaN()))
	assertFalse(t, testAlmostEqual(math.NaN(), -math.SmallestNonzeroFloat64))
	assertFalse(t, testAlmostEqual(-math.SmallestNonzeroFloat64, math.NaN()))
}

func TestAlmostEqual_opposite(t *testing.T) {
	// Comparisons of numbers on opposite sides of 0
	assertFalse(t, testAlmostEqual(1.000000001, -1.0))
	assertFalse(t, testAlmostEqual(-1.0, 1.000000001))
	assertFalse(t, testAlmostEqual(-1.000000001, 1.0))
	assertFalse(t, testAlmostEqual(1.0, -1.000000001))
	assertTrue(t, testAlmostEqual(10*math.SmallestNonzeroFloat64, 10*-math.SmallestNonzeroFloat64))
	assertFalse(t, testAlmostEqual(10000*math.SmallestNonzeroFloat64, 10000*-math.SmallestNonzeroFloat64))
}

func TestAlmostEqual_ulp(t *testing.T) {
	// The really tricky part - comparisons of numbers very close to zero.
	assertTrue(t, testAlmostEqual(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))
	assertTrue(t, testAlmostEqual(math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64))
	assertTrue(t, testAlmostEqual(-math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))
	assertTrue(t, testAlmostEqual(math.SmallestNonzeroFloat64, 0))
	assertTrue(t, testAlmostEqual(0, math.SmallestNonzeroFloat64))
	assertTrue(t, testAlmostEqual(-math.SmallestNonzeroFloat64, 0))
	assertTrue(t, testAlmostEqual(0, -math.SmallestNonzeroFloat64))

	assertFalse(t, testAlmostEqual(0.000000001, -math.SmallestNonzeroFloat64))
	assertFalse(t, testAlmostEqual(0.000000001, math.SmallestNonzeroFloat64))
	assertFalse(t, testAlmostEqual(math.SmallestNonzeroFloat64, 0.000000001))
	assertFalse(t, testAlmostEqual(-math.SmallestNonzeroFloat64, 0.000000001))
}

func testAlmostEqual(a, b float64) bool {
	return AlmostEqual(a, b, 0.00001)
}

func assertTrue(t *testing.T, actual bool) {
	if actual != true {
		t.Fatal("expected: true")
	}
}

func assertFalse(t *testing.T, actual bool) {
	if actual != false {
		t.Fatal("expected: false")
	}
}

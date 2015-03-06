package floats

import "math"

var (
	// MinNormal is the smallest positive normal value of type float64.
	MinNormal = math.Float64frombits(0x0010000000000000)
	// MinNormal32 is the smallest positive normal value of type float32.
	MinNormal32 = math.Float32frombits(0x008000000)
)

// AlmostEqual returns true if a and b are equal within a relative error of
// ε. See http://floating-point-gui.de/errors/comparison/ for the details of the
// applied method.
func AlmostEqual(a, b, ε float64) bool {
	if a == b {
		return true
	}
	diff := math.Abs(a - b)
	if a == 0 || b == 0 || diff < MinNormal {
		return diff < ε*MinNormal
	}
	return diff/(math.Abs(a)+math.Abs(b)) < ε
}

// AlmostEqual32 returns true if a and b are equal within a relative error of
// ε. See http://floating-point-gui.de/errors/comparison/ for the details of the
// applied method.
func AlmostEqual32(a, b, ε float32) bool {
	if a == b {
		return true
	}
	diff := Abs32(a - b)
	if a == 0 || b == 0 || diff < MinNormal32 {
		return diff < ε*MinNormal32
	}
	return diff/(Abs32(a)+Abs32(b)) < ε
}

// Abs32 works like math.Abs, but for float32.
func Abs32(x float32) float32 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

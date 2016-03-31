package vec3

import (
	"fmt"
	"github.com/scritch007/gm/math32"
)

type Vec3 [3]float32

func New(x, y, z float32) *Vec3 {
	return &Vec3{x, y, z}
}

// Clone initializes a new Vec3 initialized with values from an existing one.
func (lhs *Vec3) Clone() *Vec3 {
	return &Vec3{lhs[0], lhs[1], lhs[2]}
}

// Cross calculates the vector cross product. Saves the result into the
// calling vector. Returns itself for function chaining.
func (lhs *Vec3) Cross(rhs *Vec3) *Vec3 {
	a, b, c := lhs[0], lhs[1], lhs[2]

	lhs[0] = b*rhs[2] - c*rhs[1]
	lhs[1] = c*rhs[0] - a*rhs[2]
	lhs[2] = a*rhs[1] - b*rhs[0]

	return lhs
}

// Div divides the the calling vector by the provided one. The result is
// saved back into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Div(rhs float32) *Vec3 {
	lhs[0] /= rhs
	lhs[1] /= rhs
	lhs[2] /= rhs
	return lhs
}

// Len returns the vector length.
func (lhs *Vec3) Len() float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

// Multiply the vector with a scalar. Returns itself.
func (lhs *Vec3) Mul(rhs float32) *Vec3 {
	lhs[0] *= rhs
	lhs[1] *= rhs
	lhs[2] *= rhs
	return lhs
}

// Norrmalize the vector. Returns itself for function chaining.
func (lhs *Vec3) Normalize() *Vec3 {
	lhs.Div(lhs.Len())
	return lhs
}

// Sub subtracts the provided vector from the calling one. The result is
// saved into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Sub(rhs *Vec3) *Vec3 {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
	return lhs
}

// Add adds the provided vector to the calling one. The result is
// saved into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Add(rhs *Vec3) *Vec3 {
	lhs[0] += rhs[0]
	lhs[1] += rhs[1]
	lhs[2] += rhs[2]
	return lhs
}

func (lhs *Vec3) MulInner(rhs *Vec3) float32 {
	var p float32
	p = 0.0
	for i := 0; i < 3; i++ {
		p += lhs[i] * rhs[i]
	}
	return p
}

func (lhs *Vec3) String() string {
	return fmt.Sprintf("%f %f %f", lhs[0], lhs[1], lhs[2])
}

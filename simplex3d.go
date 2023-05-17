package noise

import (
	"fmt"
	"math"
)

// Simplex3D is a 3D Simplex noise generator
type Simplex3D struct {
	Permutations
}

// Seed the noise functions. Only 65536 different seeds are supported.
// Use a float between 0 and 1 or an integer from 1 to 65536.
func NewSimplex3D(seed float64) *Simplex3D {
	s := &Simplex3D{}
	s.Seedf(seed)
	return s
}

// Noise function returns a value in the range of -1 to 1.
func (d *Simplex3D) Noise(x, y, z float64) float64 {
	var n0, n1, n2, n3 float64 // Noise contributions from the four corners

	// Skew the input space to determine which simplex cell we're in
	s := (x + y + z) * F3 // Hairy factor for 2D
	i := math.Floor(x + s)
	j := math.Floor(y + s)
	k := math.Floor(z + s)

	t := (i + j + k) * G3
	x0 := x - i + t // The x,y distances from the cell origin, unskewed.
	y0 := y - j + t
	z0 := z - k + t

	// For the 3D case, the simplex shape is a slightly irregular tetrahedron.
	// Determine which simplex we are in.
	var i1, j1, k1 uint8 // Offsets for second corner of simplex in (i,j,k) coords
	var i2, j2, k2 uint8 // Offsets for third corner of simplex in (i,j,k) coords
	if x0 >= y0 {
		if y0 >= z0 {
			i1, j1, k1, i2, j2, k2 = 1, 0, 0, 1, 1, 0
		} else if x0 >= z0 {
			i1, j1, k1, i2, j2, k2 = 1, 0, 0, 1, 0, 1
		} else {
			i1, j1, k1, i2, j2, k2 = 0, 0, 1, 1, 0, 1
		}
	} else {
		if y0 < z0 {
			i1, j1, k1, i2, j2, k2 = 0, 0, 1, 0, 1, 1
		} else if x0 < z0 {
			i1, j1, k1, i2, j2, k2 = 0, 1, 0, 0, 1, 1
		} else {
			i1, j1, k1, i2, j2, k2 = 0, 1, 0, 1, 1, 0
		}
	}

	// A step of (1,0,0) in (i,j,k) means a step of (1-c,-c,-c) in (x,y,z),
	// a step of (0,1,0) in (i,j,k) means a step of (-c,1-c,-c) in (x,y,z), and
	// a step of (0,0,1) in (i,j,k) means a step of (-c,-c,1-c) in (x,y,z), where
	// c = 1/6.
	x1 := x0 - float64(i1) + G3 // Offsets for second corner
	y1 := y0 - float64(j1) + G3
	z1 := z0 - float64(k1) + G3

	x2 := x0 - float64(i2) + 2*G3 // Offsets for third corner
	y2 := y0 - float64(j2) + 2*G3
	z2 := z0 - float64(k2) + 2*G3

	x3 := x0 - 1 + 3*G3 // Offsets for fourth corner
	y3 := y0 - 1 + 3*G3
	z3 := z0 - 1 + 3*G3

	// Work out the hashed gradient indices of the four simplex corners
	i0 := uint8(i)
	j0 := uint8(j)
	k0 := uint8(k)
	gi0 := d.gradP[i0+d.perm[j0+d.perm[k0]]]
	gi1 := d.gradP[i0+i1+d.perm[j0+j1+d.perm[k0+k1]]]
	gi2 := d.gradP[i0+i2+d.perm[j0+j2+d.perm[k0+k2]]]
	gi3 := d.gradP[i0+1+d.perm[j0+1+d.perm[k0+1]]]

	// Calculate the contribution from the four corners
	t0 := 0.6 - x0*x0 - y0*y0 - z0*z0
	if t0 < 0 {
		n0 = 0
	} else {
		t0 *= t0
		n0 = t0 * t0 * gi0.dot3(x0, y0, z0) // (x,y) of grad3 used for 2D gradient
	}
	t1 := 0.6 - x1*x1 - y1*y1 - z1*z1
	if t1 < 0 {
		n1 = 0
	} else {
		t1 *= t1
		n1 = t1 * t1 * gi1.dot3(x1, y1, z1)
	}
	t2 := 0.6 - x2*x2 - y2*y2 - z2*z2
	if t2 < 0 {
		n2 = 0
	} else {
		t2 *= t2
		n2 = t2 * t2 * gi2.dot3(x2, y2, z2)
	}
	t3 := 0.6 - x3*x3 - y3*y3 - z3*z3
	if t3 < 0 {
		n3 = 0
	} else {
		t3 *= t3
		n3 = t3 * t3 * gi3.dot3(x3, y3, z3)
	}

	result := 32 * (n0 + n1 + n2 + n3)
	if result < -1.0 || result > 1.0 {
		panic(fmt.Sprintf("Simplex3D: invalid noise value %f", result))
	}
	return result
}

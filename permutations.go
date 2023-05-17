package noise

import "math"

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

type grad struct {
	x, y, z float64
}

func (g grad) dot2(x, y float64) float64 {
	return g.x*x + g.y*y
}

func (g grad) dot3(x, y, z float64) float64 {
	return g.x*x + g.y*y + g.z*z
}

var grad3 = []grad{
	{1, 1, 0},
	{-1, 1, 0},
	{1, -1, 0},
	{-1, -1, 0},
	{1, 0, 1},
	{-1, 0, 1},
	{1, 0, -1},
	{-1, 0, -1},
	{0, 1, 1},
	{0, -1, 1},
	{0, 1, -1},
	{0, -1, -1},
}

// To remove the need for index wrapping, double the permutation table length
var PERMUTATIONS = [...]uint8{
	151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
}

var F2 = 0.5 * (math.Sqrt(3) - 1.0)
var G2 = (3.0 - math.Sqrt(3)) / 6.0

const F3 = 1.0 / 3.0
const G3 = 1.0 / 6.0

type Permutations struct {
	perm  [512]uint8
	gradP [512]grad
}

// Seed the noise functions. Only 65536 different seeds are supported.
// Use a float between 0 and 1 or an integer from 1 to 65536.
func (s *Permutations) Seedf(seed float64) {
	if seed > 0 && seed < 1 {
		seed *= 65536
	}
	s.Seed(uint16(math.Floor(seed)))
}

// Seed the noise functions. Only 65536 different seeds are supported.
// Use a an integer from 1 to 65536.
// This isn't a very good seeding function, but it works ok. It supports 2^16
// different seed values. Write something better if you need more seeds.
func (s *Permutations) Seed(seed uint16) {
	if seed < 256 {
		seed |= seed << 8
	}

	for i := 0; i < 256; i++ {
		var v uint8
		if i&1 != 0 {
			v = PERMUTATIONS[i] ^ uint8(seed&255)
		} else {
			v = PERMUTATIONS[i] ^ uint8((seed>>8)&255)
		}
		s.perm[i], s.perm[i+256] = v, v
		s.gradP[i], s.gradP[i+256] = grad3[v%12], grad3[v%12]
	}
}

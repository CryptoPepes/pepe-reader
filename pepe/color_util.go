package pepe

import "github.com/lucasb-eyer/go-colorful"

// Force colors to parse, panic on invalid hex code
func MustHexColor(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustHexColor: " + err.Error())
	}
	return c
}

// Gradient table, each color key-point has a position in range 0.0, 1.0
// Note: Gradient key-points should be sorted!
type GradientTable []struct {
	Col colorful.Color
	Pos float64
}

// Returns a HCL-blend between the two colors around input `t`.
func (gradient GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(gradient) - 1; i++ {
		c1 := gradient[i]
		c2 := gradient[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// t is now between key-points c1 and c2. Blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}

	// Nothing found? Then t is at or past the last gradient key-point.
	return gradient[len(gradient) - 1].Col
}

package objects

type Color struct {
    R, G, B float64
}

var (
    Black = Color{0.0,0.0,0.0}
    White = Color{1.0,1.0,1.0}
    Blue = Color{0.5,0.7,1.0}
)

func (c Color) Add(o Color) Color {
    return Color{c.R + o.R, c.G + c.G, c.B + o.B}
}

func (c Color) Multiply(o Color) Color {
	  return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c Color) AddScalar(a float64) Color {
	  return Color{c.R + a, c.G + a, c.B + a}
}

func (c Color) SubtractScalar(a float64) Color {
	  return Color{c.R - a, c.G - a, c.B - a}
}

func (c Color) MultiplyScalar(a float64) Color {
	  return Color{c.R * a, c.G * a, c.B * a}
}

func (c Color) DivideScalar(a float64) Color {
	  return Color{c.R/a, c.G/a, c.B/a}
}

func Gradient(a, b Color, f float64) Color {
	  // scale between 0.0 and 1.0
	  f = 0.5 * (f + 1.0)
	  // linear blend: blended_value = (1 - f) * a + f * b
	  return a.MultiplyScalar(1.0 - f).Add(b.MultiplyScalar(f))
}

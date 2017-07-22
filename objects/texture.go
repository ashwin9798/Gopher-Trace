package objects

//Main texture interface
type Texture interface {
    Value(u,v float64)(Vector, Vector)
}

////////////////////////////////////////////////
type ConstantTexture struct {
    C Color
}

func (c *ConstantTexture) Value(u,v float64)(Vector, Vector) {
    return Color, Vector{}
}

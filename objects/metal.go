package objects

type Metal struct {
    C Color
    Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit)(bool, Ray) {
    directionReflected := reflect(input.Direction, hit.Normal)
    scatteredRay := Ray{hit.Point, directionReflected.Add(VectorInUnitSphere().MultiplyScalar(m.Fuzz))}
    bounced := directionReflected.Dot(hit.Normal) > 0
    return bounced, scatteredRay
}

func (m Metal) Color() Color {
    return m.C
}

func reflect(v Vector, n Vector) Vector {
    product := (v.Dot(n))*2
    return v.Subtract(n.MultiplyScalar(product))
}

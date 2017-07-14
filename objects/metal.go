package objects

type Metal struct {
    C Color
    Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit)(bool, Ray) {
    directionReflected := input.Direction.Reflect(hit.Normal)
    scatteredRay := Ray{hit.Point, directionReflected.Add(VectorInUnitSphere().MultiplyScalar(m.Fuzz))}
    bounced := directionReflected.Dot(hit.Normal) > 0
    return bounced, scatteredRay
}

func (m Metal) Color() Color {
    return m.C
}

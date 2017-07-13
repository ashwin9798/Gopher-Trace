package objects

type Metal struct {
    C Vector
}

func (m Metal) Bounce(input Ray, hit Hit)(bool, Ray) {
    directionReflected := reflect(input.Direction, hit.Normal)
    scatteredRay := Ray{hit.Point, directionReflected}
    bounced := directionReflected.Dot(hit.Normal) > 0
    return bounced, scatteredRay
}

func (m Metal) Color() Vector {
    return m.C
}

func reflect(v Vector, n Vector) Vector {
    product := (v.Dot(n))*2
    return v.Subtract(n.MultiplyScalar(product))
}

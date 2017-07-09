package objects

import (
    "math"
)

type Sphere struct {
    Center Vector
    Radius float64
}

func (s *Sphere) Hit(r *Ray, tmin float64, tmax float64)(bool, HitRecord) {
    oc := r.Origin.Subtract(s.Center)
    a := r.Direction.Dot(r.Direction)
    b := 2.0 * oc.Dot(r.Direction)
    c := oc.Dot(oc) - s.Radius*s.Radius
    discriminant := b*b - 4*a*c

    record := HitRecord{}

    //TBC
    if discriminant > 0.0 {
        t := (-b - math.Sqrt(b*b-a*c))
    }
}

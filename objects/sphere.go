package objects

import (
    "math"
)

type Sphere struct {
    Center Vector
    Radius float64
}

func (s *Sphere) Hit(r *Ray, tMin float64, tMax float64)(bool, HitRecord) {
    oc := r.Origin.Subtract(s.Center)
    a := r.Direction.Dot(r.Direction)
    b := 2.0 * oc.Dot(r.Direction)
    c := oc.Dot(oc) - s.Radius*s.Radius
    discriminant := b*b - 4*a*c

    record := HitRecord{}

    if discriminant > 0.0 {
        temp := (-b - math.Sqrt(b*b-a*c))
        if temp < tMax && temp > tMin {
            record.T = temp
            record.P = r.Point(temp)
            record.Normal = (record.P.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
        temp = (-b + math.Sqrt(b*b-a*c))  //other root
        if temp < tMax && temp > tMin {
            record.T = temp
            record.P = r.Point(temp)
            record.Normal = (record.P.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
    }
    return false, record
}

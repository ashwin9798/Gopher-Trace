package objects

import (
    "math"
)

type Sphere struct {
    Center Vector
    Radius float64
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64)(bool, Hit) {
    oc := r.Origin.Subtract(s.Center)
    a := r.Direction.Dot(r.Direction)
    b := 2.0 * oc.Dot(r.Direction)
    c := oc.Dot(oc) - s.Radius*s.Radius
    discriminant := b*b - 4*a*c

    record := Hit{Material: s.material}

    if discriminant > 0.0 {
        temp := (-b - math.Sqrt(b*b-a*c))
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
        temp = (-b + math.Sqrt(b*b-a*c))  //other root
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
    }
    return false, Hit{}
}

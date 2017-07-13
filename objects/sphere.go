package objects

import (
    "math"
)

type Sphere struct {
    Center Vector
    Radius float64
    Material
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64)(bool, Hit) {
    oc := r.Origin.Subtract(s.Center)
    a := r.Direction.Dot(r.Direction)
    b := oc.Dot(r.Direction)
    c := oc.Dot(oc) - s.Radius*s.Radius
    discriminant := b*b - a*c

    record := Hit{Material: s.Material}

    if discriminant > 0 {
        temp := (-b - math.Sqrt(discriminant)) / a
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
        temp = (-b + math.Sqrt(discriminant)) / a  //other root
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(s.Center)).DivideScalar(s.Radius)
            return true, record
        }
    }
    return false, Hit{}
}

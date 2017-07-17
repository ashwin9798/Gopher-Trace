package objects

import (
    "math"
)

type MovingSphere struct {
    Center0, Center1 Vector
    time0, time1 float64
    Radius float64
    Material
}

func NewMovingSphere(center0, center1 Vector, t0, t1, radius float64, m Material) *Sphere {
	return &Sphere{center0, center1, t0, t1, radius, m}
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

func (m MovingSphere) Center(time float64) Vector {
    return m.Center0.Add((m.Center1.Subtract(m.Center0)).MultiplyScalar((time - m.time0) / (m.time1 - m.time0)))
}

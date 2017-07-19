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

func NewMovingSphere(center0, center1 Vector, t0, t1, radius float64, m Material) *MovingSphere {
	return &MovingSphere{center0, center1, t0, t1, radius, m}
}

func (m *MovingSphere) Hit(r Ray, tMin float64, tMax float64)(bool, Hit) {
    oc := r.Origin.Subtract(m.Center(r.Time))
    a := r.Direction.Dot(r.Direction)
    b := oc.Dot(r.Direction)
    c := oc.Dot(oc) - m.Radius*m.Radius
    discriminant := b*b - a*c

    record := Hit{Material: m.Material}

    if discriminant > 0 {
        temp := (-b - math.Sqrt(discriminant)) / a
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(m.Center(r.Time))).DivideScalar(m.Radius)
            return true, record
        }
        temp = (-b + math.Sqrt(discriminant)) / a  //other root
        if temp < tMax && temp > tMin {
            record.T = temp
            record.Point = r.Point(temp)
            record.Normal = (record.Point.Subtract(m.Center(r.Time))).DivideScalar(m.Radius)
            return true, record
        }
    }
    return false, Hit{}
}

func (m MovingSphere) Center(time float64) Vector {
    return m.Center0.Add((m.Center1.Subtract(m.Center0)).MultiplyScalar((time - m.time0) / (m.time1 - m.time0)))
}

func (m *MovingSphere) BoundingBox(t0, t1 float64)(bool, AABB) {
      box0 := AABB{m.Center(t0) - Vector{m.Radius,m.Radius,m.Radius}, m.Center(t0) + Vector{m.Radius,m.Radius,m.Radius}}
      box1 := AABB{m.Center(t1) - Vector{m.Radius,m.Radius,m.Radius}, m.Center(t1) + Vector{m.Radius,m.Radius,m.Radius}}
      box := surroundingBox(box0, box1);
      return true, box;
}

func surroundingBox(box0, box1 AABB) AABB{
      small := Vector{ffmin(box0.min.X, box1.min.X), ffmin(box0.min.Y, box1.min.Y), ffmin(box0.min.Z, box1.min.Z)}
      big := Vector{ffmax(box0.max.X, box1.max.X), ffmin(box0.max.Y, box1.max.Y), ffmin(box0.max.Z, box1.max.Z)}
      return AABB{small,big}
}

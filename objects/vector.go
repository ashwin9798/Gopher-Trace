package objects

import (
    "math"
    "math/rand"
)

type Vector struct {
    X, Y, Z float64
}

var UnitVector = Vector{1,1,1}

func VectorInUnitSphere() Vector {
    for {
      r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
      p := r.MultiplyScalar(2.0).Subtract(UnitVector)
      if p.SquaredLength() >= 1.0 {
          return p
      }
    }
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Length() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Dot(o Vector) float64 {
    return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Add(o Vector) Vector {
    return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector) Subtract(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vector) Divide(o Vector) Vector {
	return Vector{v.X / o.X, v.Y / o.Y, v.Z / o.Z}
}

func (v Vector) Multiply(o Vector) Vector {
	return Vector{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

func (v Vector) AddScalar(a float64) Vector {
	return Vector{v.X + a, v.Y + a, v.Z + a}
}

func (v Vector) SubtractScalar(a float64) Vector {
	return Vector{v.X - a, v.Y - a, v.Z - a}
}

func (v Vector) MultiplyScalar(a float64) Vector {
	return Vector{v.X * a, v.Y * a, v.Z * a}
}

func (v Vector) DivideScalar(a float64) Vector {
	return Vector{v.X/a, v.Y/a, v.Z/a}
}

func (v Vector) Normalize() Vector {
    l := v.Length()
    return Vector{v.X/l, v.Y/l, v.Z/l}
}

func (v Vector) Refract(n Vector, ni_over_nt float64) (bool, Vector) {
    unitVec := v.Normalize()
    dt := unitVec.Dot(n.Normalize())
    discriminant := 1.0 - ni_over_nt*ni_over_nt*(1-dt*dt)
    if discriminant > 0 {
        a := unitVec.Subtract(n.MultiplyScalar(dt)).MultiplyScalar(ni_over_nt)
        b := n.MultiplyScalar(math.Sqrt(discriminant))
        return true, a.Subtract(b)
    }
    return false, Vector{}
}


func (v Vector) Reflect(n Vector) Vector {
    product := (v.Dot(n))*2
    return v.Subtract(n.MultiplyScalar(product))
}

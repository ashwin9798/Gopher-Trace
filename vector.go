package main

import (
    "math"
)

type Vector struct {
    X, Y, Z float64
}

func (v Vector) Length() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Dot(o Vector) float64 {
    return v.X*o*X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Add(o Vector) Vector {
    return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector) Subtract(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
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

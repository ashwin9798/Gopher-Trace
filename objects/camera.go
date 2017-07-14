package objects

import (
    "math"
)

type Camera struct {
    lowerLeft, horizontal, vertical, origin Vector
}

func NewCamera(lookFrom, lookAt, vUp Vector, vfov, aspect float64) Camera {
    c := Camera{}

    theta := vfov * math.Pi/180
    halfHeight := math.Tan(theta/2)
    halfWidth := aspect * halfHeight
    w := lookFrom.Subtract(lookAt).Normalize()
    u := vUp.Cross(w).Normalize()
    v := w.Cross(u)

    c.origin = lookFrom
    c.lowerLeft =  c.origin.Subtract(u.MultiplyScalar(halfWidth)).Subtract(v.MultiplyScalar(halfHeight)).Subtract(w)
	  c.horizontal = u.MultiplyScalar(2 * halfWidth)
	  c.vertical = v.MultiplyScalar(2 * halfHeight)

	   return c
}

func (c *Camera) RayAt(u float64, v float64) Ray {
    horizontal := c.horizontal.MultiplyScalar(u)
    vertical := c.vertical.MultiplyScalar(v)

    direction := c.lowerLeft.Add(horizontal).Add(vertical).Subtract(c.origin)
    return Ray{c.origin, direction}
}

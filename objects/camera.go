package objects

import (
    "math"
    "math/rand"
)

type Camera struct {
    lowerLeft, horizontal, vertical, origin, u, v, w Vector
    lensRadius float64
    shutterOpenTime float64
    shutterCloseTime float64
}

func NewCamera(lookFrom, lookAt, vUp Vector, vfov, aspect, aperture, focusDist float64) Camera {
    c := Camera{}

    theta := vfov * math.Pi/180
    halfHeight := math.Tan(theta/2)
    halfWidth := aspect * halfHeight
    w := lookFrom.Subtract(lookAt).Normalize()
    u := vUp.Cross(w).Normalize()
    v := w.Cross(u)

    x := u.MultiplyScalar(halfWidth * focusDist)
	  y := v.MultiplyScalar(halfHeight * focusDist)

    c.lensRadius = aperture
    c.origin = lookFrom
    c.lowerLeft =  c.origin.Subtract(x).Subtract(y).Subtract(w.MultiplyScalar(focusDist))
	  c.horizontal = u.MultiplyScalar(2 * halfWidth * focusDist)
	  c.vertical = v.MultiplyScalar(2 * halfHeight * focusDist)

    c.w = w
    c.u = u
    c.v = v

    return c
}

func (c *Camera) RayAt(u float64, v float64) Ray {
    rd := randomInUnitDisc().MultiplyScalar(c.lensRadius)
    offset := c.u.MultiplyScalar(rd.X).Add(c.v.MultiplyScalar(rd.Y))

    horizontal := c.horizontal.MultiplyScalar(u)
    vertical := c.vertical.MultiplyScalar(v)
    origin := c.origin.Add(offset)

    direction := c.lowerLeft.Add(horizontal).Add(vertical).Subtract(c.origin).Subtract(offset)
    time := c.shutterOpenTime + rand.Float64()*(c.shutterCloseTime - c.shutterOpenTime)

    return Ray{c.origin, direction, c.Time}
}

func randomInUnitDisc() Vector {
	  var p Vector
	  for {
		    p = Vector{rand.Float64(), rand.Float64(), 0}.MultiplyScalar(2).Subtract(Vector{1, 1, 0})
		    if p.Dot(p) < 1.0 {
			       return p
		    }
	  }
}

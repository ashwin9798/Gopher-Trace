package objects

import (
    "math"
    "math/rand"
)

type Dielectric struct {
    RefIndex float64
}

func (d Dielectric) Color() Color {
	return Color{1.0, 1.0, 1.0}
}

func (d Dielectric) Bounce(input Ray, hit Hit) (bool, Ray) {
    var outwardNormal Vector

    var ni_over_nt float64
    var refracted Vector
    var cosine float64

    if input.Direction.Dot(hit.Normal) > 0 {
        outwardNormal = hit.Normal.MultiplyScalar(-1)
        ni_over_nt = d.RefIndex
        cosine = (d.RefIndex*(input.Direction.Dot(hit.Normal))/input.Direction.Length())
    } else {
        outwardNormal = hit.Normal
        ni_over_nt = 1.0/d.RefIndex
        cosine = -((d.RefIndex*(input.Direction.Dot(hit.Normal))/input.Direction.Length()))
    }
    var success bool
	  var reflectProbability float64

	  if success, refracted = input.Direction.Refract(outwardNormal, ni_over_nt); success {
		    reflectProbability = d.schlick(cosine)
    } else {
        reflectProbability = 1.0
    }
    if rand.Float64() < reflectProbability {
		    reflected := input.Direction.Reflect(hit.Normal)
		    return true, Ray{hit.Point, reflected, 0.0}
	  }
    return true, Ray{hit.Point, refracted, 0.0}
}

func (d Dielectric) schlick(cosine float64) float64 {
    r0 := (1-d.RefIndex)/(1+d.RefIndex)
    r0 = r0*r0
    return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

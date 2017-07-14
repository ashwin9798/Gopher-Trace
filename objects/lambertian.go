package objects

type Lambertian struct {
    C Color
}

func (l Lambertian) Bounce(input Ray, hit Hit) (bool, Ray) {
	  direction := hit.Normal.Add(VectorInUnitSphere())
	  return true, Ray{hit.Point, direction}
}

func (l Lambertian) Color() Color {
	 return l.C
}

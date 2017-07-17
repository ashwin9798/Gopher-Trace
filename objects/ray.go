package objects

type Ray struct {
    Origin, Direction Vector
    Time float64
}

func (r Ray) Point(t float64) Vector {
    b := r.Direction.MultiplyScalar(t)
    a := r.Origin
    return a.Add(b)
}

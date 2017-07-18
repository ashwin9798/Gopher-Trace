package objects

type AABB struct {  //axis aligned bounded box
    max, min Vector
}

func ffmin(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func ffmax(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}


func (a *AABB) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
    for i := 0; i < 3; a++ {
        var t0, t1 float64
        if i == 1 {
            t0 = ffmin(((a.min.X - r.Origin.X)/r.Direction.X),((a.max.X - r.Origin.X)/r.Direction.X))
            t1 = ffmax(((a.min.X - r.Origin.X)/r.Direction.X),((a.max.X - r.Origin.X)/r.Direction.X))
        }
        if i == 2 {
            t0 = ffmin(((a.min.Y - r.Origin.Y)/r.Direction.Y),((a.max.Y - r.Origin.Y)/r.Direction.Y))
            t1 = ffmax(((a.min.Y - r.Origin.X)/r.Direction.Y),((a.max.Y - r.Origin.Y)/r.Direction.Y))
        }
        if i == 3 {
            t0 = ffmin(((a.min.Z - r.Origin.Z)/r.Direction.Z),((a.max.Z - r.Origin.Z)/r.Direction.Z))
            t1 = ffmax(((a.min.Z - r.Origin.Z)/r.Direction.Z),((a.max.Z - r.Origin.Z)/r.Direction.Z))
        }
        tMin = ffmax(t0, tMin)
        tMax = ffmin(t1, tMax)
        if tMax <= tMin {
            return false, Hit{}
        }
        return true, Hit{}
    }
}

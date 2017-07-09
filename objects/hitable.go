package objects

type HitRecord struct {
    T float64
    P, normal Vector
}

type Hitable interface {
    Hit(r *ray, tMin float64, tMax float64)(bool, HitRecord)
}

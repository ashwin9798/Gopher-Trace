package objects

type Material interface {
    Bounce(input Ray, hit Hit) (bool, Ray)
    Color() Color
}

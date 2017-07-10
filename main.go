package main

import (
  "fmt"
  "os"

  obj "objects"
)

func check(e error, s string) {
    if e != nil {
        fmt.Fprintf(os.Stderr, s, e)
        os.Exit(1)
    }
}


const (
	dimensionsX = 400 // size of x
	dimensionsY = 200 // size of y
	numSamples = 100 // number of samples for aa
	col  = 255.99
)

func gradient(v *obj.Vector) obj.Vector {
    t := 0.5 * (v.Y + 1.0)
    // linear blend: blended_value = (1 - t) * white + t * blue
    return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func color(r *obj.Ray, h obj.Hitable) obj.Vector {
    hit, record := h.Hit(r, 0.0, math.MaxFloat64)

    if hit {
        return record.Normal.AddScalar(1.0).MultiplyScalar(0.5)
    }
    unitDirection := r.Direction.Normalize()
}

func main() {

    f, err := os.Create("out.ppm")
    defer f.Close()

    check(err, "Error opening file: %v\n")

    _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", dimensionsX, dimensionsY)

    check(err, "Error writting to file: %v\n")

    lowerLeft := obj.Vector{-2.0, -1.0, -1.0}
	  horizontal := obj.Vector{4.0, 0.0, 0.0}
	  vertical := obj.Vector{0.0, 2.0, 0.0}
	  origin := obj.Vector{0.0, 0.0, 0.0}

    //loop through all the pixels from top left to bottom right
    //write rgb values for each
    for j := dimensionsY-1; j>=0; j-- {
      for i := 0; i<dimensionsX; i++ {
          u := float64(i) / float64(dimensionsX)
          v := float64(j) / float64(dimensionsY)

          position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

          // direction = lowerLeft + (u * horizontal) + (v * vertical)
          direction := lowerLeft.Add(position)

          rgb := obj.Ray{origin, direction}.Color()

          // get intensity of colors
          ir := int(color * rgb.X)
          ig := int(color * rgb.Y)
          ib := int(color * rgb.Z)

          _, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

          check(err, "Error writing to file: %v\n")
      }
    }

}

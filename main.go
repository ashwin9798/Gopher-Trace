package main

import (
  "fmt"
  "os"
  "math"
  "math/rand"

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

var (
	  white = obj.Vector{1.0, 1.0, 1.0}
	  blue  = obj.Vector{0.5, 0.7, 1.0}
	  camera = obj.NewCamera()
	  sphere = obj.Sphere{obj.Vector{0, 0, -1}, 0.5}
	  floor  = obj.Sphere{obj.Vector{0, -100.5, -1}, 100}
    world = obj.World{[]obj.Hitable{&sphere, &floor}}
)

func gradient(r obj.Ray) obj.Vector {
    v := r.Direction.Normalize()
    t := 0.5 * (v.Y + 1.0)
    // linear blend: blended_value = (1 - t) * white + t * blue
    return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func color(r obj.Ray, world obj.Hitable, depth int) obj.Vector {
    hit, record := world.Hit(r, 0.001, math.MaxFloat64)

    if hit {
      if depth < 50 {
        bounced, bouncedRay := record.Bounce(r, record)
        if bounced {
          newColor := color(bouncedRay, world, depth+1)
          return record.Material.Color().Multiply(newColor)
        }
      }
      return obj.Vector{}
    }
    return gradient(r)
}

func main() {

    f, err := os.Create("out.ppm")
    defer f.Close()

    check(err, "Error opening file: %v\n")

    _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", dimensionsX, dimensionsY)

    check(err, "Error writting to file: %v\n")

    //loop through all the pixels from top left to bottom right
    //write rgb values for each
    for j := dimensionsY-1; j>=0; j-- {
      for i := 0; i<dimensionsX; i++ {
        rgb := obj.Vector{}

          for s := 0; s < numSamples; s++ {
              u := (float64(i)+ rand.Float64())/ float64(dimensionsX)
              v := (float64(j) + rand.Float64())/ float64(dimensionsY)
              r := camera.RayAt(u,v)
              color := color(&r, &world)
				      rgb = rgb.Add(color)
          }
          // average
			    rgb = rgb.DivideScalar(float64(numSamples))
          // get intensity of colors
          ir := int(col * rgb.X)
          ig := int(col * rgb.Y)
          ib := int(col * rgb.Z)

          _, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
          check(err, "Error writing to file: %v\n")
      }
    }

}

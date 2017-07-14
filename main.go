package main

import (
  "fmt"
  "os"
  "math"
  "math/rand"
  "time"

  obj "objects"
)

const (
	 dimensionsX = 400 // size of x
	 dimensionsY = 200 // size of y
	 numSamples = 100 // number of samples for aa
	 col  = 255.99
)

func check(e error, s string) {
    if e != nil {
        fmt.Fprintf(os.Stderr, s, e)
        os.Exit(1)
    }
}

func color(r obj.Ray, world obj.Hitable, depth int) obj.Color {
    hit, record := world.Hit(r, 0.001, math.MaxFloat64)

    if hit {
      if depth < 50 {
        bounced, bouncedRay := record.Bounce(r, record)
        if bounced {
          newColor := color(bouncedRay, world, depth+1)
          return record.Material.Color().Multiply(newColor)
        }
      }
      return obj.Black
    }
    return obj.Gradient(obj.White, obj.Blue, r.Direction.Normalize().Y)
}

func createFile() *os.File {
    f, err := os.Create("out.ppm")
    check(err, "Error opening file: %v\n")

    _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", dimensionsX, dimensionsY)
    check(err, "Error writting to file: %v\n")
    return f
}

func writeFile(f *os.File, rgb obj.Color) {
    ir := int(col * math.Sqrt(rgb.R))
    ig := int(col * math.Sqrt(rgb.G))
    ib := int(col * math.Sqrt(rgb.B))

    _, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
    check(err, "Error writing to file: %v\n")
}

func sample(world *obj.World, camera *obj.Camera, i, j int) obj.Color {
    rgb := obj.Color{}
    for s := 0; s < numSamples; s++ {
        u := (float64(i)+ rand.Float64())/ float64(dimensionsX)
        v := (float64(j) + rand.Float64())/ float64(dimensionsY)
        r := camera.RayAt(u,v)
        color := color(r, world, 0)
        rgb = rgb.Add(color)
    }
    // average
    return rgb.DivideScalar(float64(numSamples))
}

func render(world *obj.World, camera *obj.Camera) {
	   ticker := time.NewTicker(time.Millisecond * 100)

	    go func() {
		      for {
			         <-ticker.C
			         fmt.Print(".")
		      }
	    }()

	    f := createFile()
	    defer f.Close()

	    start := time.Now()

      //loop through all the pixels from top left to bottom right
      //write rgb values for each
	    for j := dimensionsY - 1; j >= 0; j-- {
		      for i := 0; i < dimensionsX; i++ {
			         rgb := sample(world, camera, i, j)
			         writeFile(f, rgb)
		      }
	    }
	    ticker.Stop()
	    fmt.Printf("\nDone.\nElapsed: %v\n", time.Since(start))
}


func main() {
    //objects in the world
    lookFrom := obj.Vector{-2, 2, 1}
	  lookAt := obj.Vector{0, 0, -1}
	  vUp := obj.Vector{0, 1, 0}

	  camera := obj.NewCamera(lookFrom, lookAt, vUp, 90, float64(dimensionsX)/float64(dimensionsY))

	  world := obj.World{}

	  radius := math.Cos(math.Pi / 4)

	  blue := obj.NewSphere(-radius, 0, -1, radius, obj.Lambertian{obj.Color{0, 0, 1}})
	  red := obj.NewSphere(radius, 0, -1, radius, obj.Lambertian{obj.Color{1, 0, 0}})

	  world.AddAll(&blue, &red)

	  render(&world, &camera)
}

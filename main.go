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
	 dimensionsX = 600 // size of x
	 dimensionsY = 500 // size of y
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

func createRandomWorld() *obj.World {
    myWorld := &obj.World{}
    myWorld.Add(obj.NewSphere(0,-1000,0, 1000, obj.Lambertian{obj.Color{0.5, 0.5, 0.5}}))
    for a := -11; a < 11; a++ {
        for b := -11; b < 11; b++ {
            chooseMaterial := rand.Float64()
            center := obj.Vector{float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64()}
            if (center.Subtract(obj.Vector{4,0.2,0}).Length()) > 0.9 {
                if chooseMaterial < 0.8 {
                    myWorld.Add(obj.NewSphere(center.X, center.Y, center.Z, 0.2, obj.Lambertian{obj.Color{rand.Float64()*rand.Float64(),rand.Float64()*rand.Float64(),rand.Float64()*rand.Float64()}}))
                } else if chooseMaterial < 0.95 {
                    myWorld.Add(obj.NewSphere(center.X, center.Y, center.Z, 0.2, obj.Lambertian{obj.Color{0.5*(1+rand.Float64()),0.5*(1+rand.Float64()),0.5*(1+rand.Float64())}}))
                } else {
                    myWorld.Add(obj.NewSphere(center.X, center.Y, center.Z, 0.2, obj.Dielectric{1.5}))
                }
            }
        }
    }
    glass := obj.NewSphere(0, 1, 0, 1.0, obj.Dielectric{1.5})
  	lambertian := obj.NewSphere(-4, 1, 0, 1.0, obj.Lambertian{C: obj.Color{R: 0.4, G: 0.0, B: 0.1}})
	  metal := obj.NewSphere(4, 1, 0, 1.0, obj.Metal{C: obj.Color{R: 0.7, G: 0.6, B: 0.5}, Fuzz: 0.0})
    myWorld.AddAll(glass, lambertian, metal)

    return myWorld
}

func main() {
    //objects in the world
    lookFrom := obj.Vector{3, 3, 2}
	  lookAt := obj.Vector{0, 0, -1}
	  vUp := obj.Vector{0, 1, 0}

    focusDist := lookFrom.Subtract(lookAt).Length()
    aperture := 0.1

	  camera := obj.NewCamera(lookFrom, lookAt, vUp, 90, float64(dimensionsX)/float64(dimensionsY), aperture, focusDist)

    world := createRandomWorld()
	  render(world, &camera)
}

package main

import (
  "fmt"
  "os"
)

func check(e error, s string) {
    if e != nil {
        fmt.Fprintf(os.Stderr, s, e)
        os.Exit(1)
    }
}

func main() {
    dimensionsX := 400
    dimensionsY := 200

    const color = 255.99

    f, err := os.Create("out.ppm")
    defer f.Close()

    check(err, "Error opening file: %v\n")

    _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", dimensionsX, dimensionsY)

    check(err, "Error writting to file: %v\n")

    lowerLeft := Vector{-2.0, -1.0, -1.0}
	  horizontal := Vector{4.0, 0.0, 0.0}
	  vertical := Vector{0.0, 2.0, 0.0}
	  origin := Vector{0.0, 0.0, 0.0}

    //loop through all the pixels from top left to bottom right
    //write rgb values for each
    for j:= dimensionsY-1; j>=0; j-- {
      for i := 0; i<dimensionsX; i++ {
          u := float64(i) / float64(dimensionsX)
          v := float64(j) / float64(dimensionsY)

          position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

          // direction = lowerLeft + (u * horizontal) + (v * vertical)
          direction := lowerLeft.Add(position)

          rgb := Ray{origin, direction}.Color()

          // get intensity of colors
          ir := int(color * rgb.X)
          ig := int(color * rgb.Y)
          ib := int(color * rgb.Z)

          _, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

          check(err, "Error writing to file: %v\n")
      }
    }

}

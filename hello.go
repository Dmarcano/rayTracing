package main

import (
	"fmt"
	"math/rand"

	"github.com/Dmarcan/rayTracing/ppmImage"
	_ "github.com/Dmarcan/rayTracing/ppmImage"
	"github.com/Dmarcan/rayTracing/vector"
)

func main() {
	fmt.Printf("My favorite number is %d\n", rand.Int())

	color := vector.Color{R: 0, G: 0, B: 0}
	_ = color
	img := ppmImage.NewPPmImage(256, 256, 255)
}

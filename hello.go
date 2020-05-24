package main

import (
	"fmt"

	"github.com/Dmarcan/rayTracing/ppmImage"
	_ "github.com/Dmarcan/rayTracing/ppmImage"
)

func main() {

	img := ppmImage.NewPPmImage(256, 256, 255)
	ppmImage.SamplePpm(img)
	// writePpm(img)
	img.BufferedWrite("out/test.ppm")
	fmt.Printf("Done!")
}

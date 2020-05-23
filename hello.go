package main

import (
	"fmt"
	"os"

	"github.com/Dmarcan/rayTracing/ppmImage"
	_ "github.com/Dmarcan/rayTracing/ppmImage"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writePpm(img *ppmImage.PpmImage) {
	f, err := os.Create("out/sample.ppm")
	check(err)
	defer f.Close()

	for j := range img.Content {
		row := img.Content[j]
		for i := range row {
			color := row[i]
			line := fmt.Sprintf("%d %d %d\n", color.R, color.G, color.B)
			_, err := f.WriteString(line)
			check(err)
		}
		f.Sync()
	}
}

func main() {

	img := ppmImage.NewPPmImage(256, 256, 255)
	ppmImage.SamplePpm(img)
	// writePpm(img)
	img.BufferedWrite("out/test.ppm")
	fmt.Printf("Done!")
}

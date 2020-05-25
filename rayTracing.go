package main

import (
	"github.com/Dmarcan/rayTracing/ppmImage"
	"github.com/Dmarcan/rayTracing/ray"
	"github.com/Dmarcan/rayTracing/vector"
	"github.com/Dmarcan/rayTracing/vector/utils"
)

func main() {

	img := ppmImage.NewPPmImage(256, 256, 255)
	ppmImage.SamplePpm(img)

	sampleGradient(img)
}

func rayColor(r ray.Ray) vector.Vector3 {
	unitDirection := utils.UnitVector(r.Dir)
	t := 0.5*unitDirection.Y + 1.0
	magnitude := 1.0 - t

	return *utils.Add(*utils.Mult(vector.Vector3{1.0, 1.0, 1.0}, magnitude), *utils.Mult(vector.Vector3{0.5, 0.7, 1.0}, t))
}

func sampleGradient(img *ppmImage.PpmImage) {
	aspectRatio := 16.0 / 9.0
	imgWidth := 384
	imgHeight := int(float64(imgWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := vector.Vector3{0.0, 0.0, 0.0}
	horizontalAxis := vector.Vector3{0.0, viewportWidth, 0.0}
	verticalAxis := vector.Vector3{0.0, 0.0, viewportHeight}
	lowerLeftCorner := utils.Add(origin, *utils.Div(horizontalAxis, 2.0))
	lowerLeftCorner = utils.Subtract(*utils.Subtract(*lowerLeftCorner, *utils.Div(verticalAxis, 2.0)), vector.Vector3{0.0, 0.0, focalLength})

}

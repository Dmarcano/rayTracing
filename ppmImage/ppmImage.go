/*

A top level file that creates ppm images

*/

package ppmImage

import "github.com/Dmarcan/rayTracing/vector"

type ppmImage struct {
	width    int
	height   int
	maxPixel int
	content  [][]vector.Color
}

// func init(img *PpmImage, width, height, max_pixel int) {
// 	img.width = width
// 	img.height = height
// 	img.max_pixel = max_pixel
// }

func NewPPmImage(width, height, max_pixel int) *ppmImage {
	img := new(ppmImage)
	img.width = width
	img.height = height
	img.maxPixel = max_pixel

	content := make([][]vector.Color, height)

	for range content {
		row := make([]vector.Color, height)
		content = append(content, row)
	}

	img.content = content

	return img
}

func Sample_ppm(img *ppmImage) {

}

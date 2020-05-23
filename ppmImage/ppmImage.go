/*

A top level file that creates ppm images

*/

package ppmImage

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Dmarcan/rayTracing/vector"
)

type PpmImage struct {
	Width    int
	Height   int
	maxPixel int
	Content  [][]vector.Color
}

// func init(img *PpmImage, width, height, max_pixel int) {
// 	img.Width = width
// 	img.Height = height
// 	img.max_pixel = max_pixel
// }

/*
	Constructor for a PPM format image.
	It creates an image of specified dimensions

	PPM images have fields that store the dimensions and a pointer to a 2-D
	Color matrix that stores the content
*/
func NewPPmImage(width, height, maxPixel int) *PpmImage {
	img := new(PpmImage)
	img.Width = width
	img.Height = height
	img.maxPixel = maxPixel

	content := make([][]vector.Color, 0, height)

	for i := 0; i < height; i++ {
		row := make([]vector.Color, width)
		content = append(content, row)
	}

	img.Content = content

	return img
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Writes the contents of the file onto the given path using a write buffer
*/
func (img *PpmImage) BufferedWrite(path string) {
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	header := fmt.Sprintf("P3\n%d %d\n%d\n", img.Width, img.Height, img.maxPixel)

	writer := bufio.NewWriter(f)
	writer.WriteString(header)
	for j := range img.Content {
		row := img.Content[j]
		for i := range row {
			color := row[i]
			line := fmt.Sprintf("%d %d %d\n", color.R, color.G, color.B)
			_, err := writer.WriteString(line)
			check(err)
		}
	}
	writer.Flush()
}

/*
 Creates a basic gradient PPM image
*/
func SamplePpm(img *PpmImage) {
	// iterate the 2-D Image array first in x then y
	for i := 0; i < img.Width; i++ {
		for j := img.Height - 1; j >= 0; j-- {
			r := float64(float64(i) / float64(img.Width-1))
			g := float64(float64(j) / float64(img.Height-1))
			b := 0.25

			const multiplier = 255.99

			ir := uint8(multiplier * r)
			ig := uint8(multiplier * g)
			ib := uint8(multiplier * b)

			img.Content[j][i].SetColor(ir, ig, ib)

		}
	}
}

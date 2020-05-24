package vector

/*
	Converts a float64 3-D vector to an RGB uint8 Pixel Color
*/
func ToColor(vec Vector3) Color {
	r := uint8(vec.x)
	g := uint8(vec.y)
	b := uint8(vec.z)
	return Color{r, g, b}
}

/*
Multiplies every vector dimension by 255.999.

This is meant to convert a vector that ranges from [0,1] to [0, 255]
*/
func (vec1 *Vector3) VectorMagnitudeMultiplication() {
	magnitude_const := float64(255.999)
	vec1.Mult(magnitude_const)
}

/*
A container for float valued 3 Dimensional Coordinates

*/
type Vector3 struct {
	x float64
	y float64
	z float64
}

/*
	Adds two vectors together and mutates the calling vector to match RHS of equation
*/
func (vec1 *Vector3) Add(vec2 Vector3) {

}

/*
	Adds two vectors together and mutates the calling vector to match RHS of equation
*/
func (vec1 *Vector3) Subtract(vec2 Vector3) {

}

func (vec1 *Vector3) Mult(t float64) {

}

func (vec1 *Vector3) Div(t float64) {

}

func (vec1 Vector3) UnitVector() {

}

func (vec1 Vector3) Length() {

}

/*
A container for 1 byte color values
*/
type Color struct {
	R uint8
	G uint8
	B uint8
}

/*
 Quickly Changes a color to match another
*/
func (color *Color) SetColor(red, green, blue uint8) {
	color.R = red
	color.G = green
	color.B = blue
}

/*
	Adds the RGB color values of one color to another
*/
func (color *Color) Add(otherColor Color) {
	color.R += otherColor.R
	color.G += otherColor.G
	color.B += otherColor.B
}

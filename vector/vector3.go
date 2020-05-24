package vector

import (
	"math"
)

/*
	Converts a float64 3-D vector to an RGB uint8 Pixel Color
*/
func ToColor(vec Vector3) Color {
	r := uint8(vec.X)
	g := uint8(vec.Y)
	b := uint8(vec.Z)
	return Color{r, g, b}
}

/*
Multiplies every vector dimension by 255.999.

This is meant to convert a vector that ranges from [0,1] to [0, 255]
*/
func (vec1 *Vector3) VectorMagnitudeMultiplication() {
	magnitudeConst := float64(255.999)
	vec1.SelfMult(magnitudeConst)
}

/*
A container for float valued 3 Dimensional Coordinates

*/
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

/*
	Mutates the vector and subtracts the second vector compoennts to it
*/
func (vec1 *Vector3) SelfSubtract(vec2 Vector3) {
	vec1.X -= vec2.X
	vec1.Y -= vec2.Y
	vec1.Z -= vec2.Z
}

/*
	Mutates the vector and adds the second vector compoennts to it
*/
func (vec1 *Vector3) SelfAdd(vec2 Vector3) {
	vec1.X += vec2.X
	vec1.Y += vec2.Y
	vec1.Z += vec2.Z
}

/*
	Mutates the vector and multiplies the components by the  constant t
*/
func (vec1 *Vector3) SelfMult(t float64) {
	vec1.X *= t
	vec1.Y *= t
	vec1.Z *= t
}

/*
	Mutates the vector and divides the components by the  constant t
*/
func (vec1 *Vector3) SelfDiv(t float64) {
	vec1.SelfMult(1.0 / t)
}

/*
	Gets the eucledian length of the vectors

	Follows the equation of

	sqrt(  x^2 + y^2 + z^2)
*/
func (vec1 Vector3) Length() float64 {
	return math.Sqrt(vec1.LengthSquared())
}

/*
	Gets the squared length

	just x^2 + y^2 + z^2

	saves the use of using sqrt function to find the squared length or find easy magnitude
*/
func (vec1 Vector3) LengthSquared() float64 {
	return vec1.X*vec1.X + vec1.Y*vec1.Y + vec1.Z*vec1.Z
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

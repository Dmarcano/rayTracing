package utils

import (
	. "github.com/Dmarcan/rayTracing/vector"
)

/*
Calculates the dot product of two vectors
*/

/*
	returns a new vector which is equal to the component subtraction of two vectors
*/
func Subtract(vec1, vec2 Vector3) *Vector3 {
	newVec := new(Vector3)

	newVec.X = vec1.X - vec2.X
	newVec.Y = vec1.Y - vec2.Y
	newVec.Z = vec1.Z - vec2.Z

	return newVec
}

/*
	returns a new vector which is equal to the component addition of two vectors
*/
func Add(vec1, vec2 Vector3) *Vector3 {
	newVec := new(Vector3)

	newVec.X = vec1.X + vec2.X
	newVec.Y = vec1.Y + vec2.Y
	newVec.Z = vec1.Z + vec2.Z

	return newVec
}

/*

Adds a constant double to the vector components
*/
func AddOffset(vec1 Vector3, t float64) *Vector3 {
	newVec := new(Vector3)

	newVec.X = vec1.X + t
	newVec.Y = vec1.Y + t
	newVec.Z = vec1.Z + t
	return newVec
}

/*
	returns a new vector which is equal to the component multiplication of the vector and a double
*/
func Mult(vec Vector3, t float64) *Vector3 {
	newVec := new(Vector3)
	newVec.X = vec.X * t
	newVec.Y = vec.Y * t
	newVec.Z = vec.Z * t

	return newVec
}

/*
	returns a new vector which is equal to the component division of the vector and a double
*/
func Div(vec Vector3, t float64) *Vector3 {
	return Mult(vec, 1.0/t)
}

/*
this function takes the vector dot product

*/
func VectorDot(vec1, vec2 Vector3) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y + vec1.Z*vec2.Z
}

/*
	Finds the length 1 vevtor that points the same direction as the current vector
*/
func UnitVector(vec Vector3) *Vector3 {
	return Div(vec, vec.Length())
}

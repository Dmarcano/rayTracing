package vector

func ToColor(vec Vector3) Color {
	return Color{vec.x, vec.y, vec.z}
}

type Vector3 struct {
	x uint8
	y uint8
	z uint8
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func (color *Color) SetColor(red, green, blue uint8) {
	color.R = red
	color.G = green
	color.B = blue
}

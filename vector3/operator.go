package vector3

import "math"

func New(x, y, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func Add(v1, v2 Vector3) Vector3 {
	return v1.Add(v2)
}

func Sub(v1, v2 Vector3) Vector3 {
	return v1.Sub(v2)
}

func Mul(v1, v2 Vector3) Vector3 {
	return v1.Mul(v2)
}

func Div(v1, v2 Vector3) Vector3 {
	return v1.Div(v2)
}

func Distance(v1, v2 Vector3) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2)+math.Pow(v1.Y - v2.Y, 2)+math.Pow(v1.Z-v2.Z, 2))
}

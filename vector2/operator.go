package Vector

import "math"

func New(x, y float64) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func NewFloat32(x, y float32) Vector2 {
	return Vector2{
		X: float64(x),
		Y: float64(y),
	}
}

func NewInt32(x, y int32) Vector2 {
	return Vector2{
		X: float64(x),
		Y: float64(y),
	}
}

func NewInt64(x, y int64) Vector2 {
	return Vector2{
		X: float64(x),
		Y: float64(y),
	}
}

func Add(v1, v2 Vector2) Vector2 {
	return v1.Add(v2)
}

func Sub(v1, v2 Vector2) Vector2 {
	return v1.Sub(v2)
}

func Mul(v1, v2 Vector2) Vector2 {
	return v1.Mul(v2)
}

func Div(v1, v2 Vector2) Vector2 {
	return v1.Div(v2)
}

func Equal(v1, v2 Vector2) bool {
	return (math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2)) < 0.000000001
}

func Distance(v1, v2 Vector2) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}

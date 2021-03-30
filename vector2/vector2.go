package Vector

import (
	"math"
	"strconv"
)

type Vector2 struct {
	X, Y float64
}

var (
	Zero = Vector2{X: 0, Y: 0}
	One  = Vector2{X: 1, Y: 1}
)

func (this Vector2) String() string {
	strX := strconv.FormatFloat(this.X, 'f', 3, 64)
	strY := strconv.FormatFloat(this.Y, 'f', 3, 64)
	return "[x:" + strX + " y:" + strY + "]"
}

func (this *Vector2) Set(x, y float64) {
	this.X = x
	this.Y = y
}

func (this *Vector2) Get() (float64, float64) {
	return this.X, this.Y
}

func (this *Vector2) GetInt32() (int32, int32) {
	return int32(this.X), int32(this.Y)
}

func (this *Vector2) GetFloat32() (float32, float32) {
	return float32(this.X), float32(this.Y)
}

func (this *Vector2) Clone() Vector2 {
	return Vector2{
		X: this.X,
		Y: this.Y,
	}
}

// 取模的平方
func (this *Vector2) SqrMagnitude() float64 {
	return this.X*this.X + this.Y*this.Y
}

// 取模
func (this *Vector2) Magnitude() float64 {
	return math.Sqrt(this.SqrMagnitude())
}

// 归一化
func (this *Vector2) Normalize() {
	magnitude := this.Magnitude()
	v := this.Div(New(magnitude, magnitude))
	this.X = v.X
	this.Y = v.Y
}

func (this *Vector2) Add(v Vector2) Vector2 {
	return Vector2{
		X: this.X + v.X,
		Y: this.Y + v.Y,
	}
}

func (this *Vector2) Sub(v Vector2) Vector2 {
	return Vector2{
		X: this.X - v.X,
		Y: this.Y - v.Y,
	}
}

func (this *Vector2) Mul(v Vector2) Vector2 {
	return Vector2{
		X: this.X * v.X,
		Y: this.Y * v.Y,
	}
}

func (this *Vector2) Div(v Vector2) Vector2 {
	var x, y float64
	if v.X != 0 {
		x = this.X / v.X
	}
	if v.Y != 0 {
		x = this.Y / v.Y
	}
	return Vector2{
		X: x,
		Y: y,
	}
}

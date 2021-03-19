package vector3

import (
	"math"
	"strconv"
)

type Vector3 struct {
	X, Y, Z float64
}

var (
	Zero = Vector3{X: 0, Y: 0, Z: 0}
	One  = Vector3{X: 1, Y: 1, Z: 1}
)

func (this Vector3) String() string {
	strX := strconv.FormatFloat(this.X, 'f', 3, 64)
	strY := strconv.FormatFloat(this.Y, 'f', 3, 64)
	strZ := strconv.FormatFloat(this.Z, 'f', 3, 64)
	return "[x:" + strX + " y:" + strY + " z:" + strZ + "]"
}

func (this *Vector3) Set(x, y float64) {
	this.X = x
	this.Y = y
}

func (this *Vector3) Get() (float64, float64) {
	return this.X, this.Y
}

func (this *Vector3) Clone() Vector3 {
	return Vector3{
		X: this.X,
		Y: this.Y,
	}
}

// 取模的平方
func (this *Vector3) SqrMagnitude() float64 {
	return this.X*this.X + this.Y*this.Y
}

// 取模
func (this *Vector3) Magnitude() float64 {
	return math.Sqrt(this.SqrMagnitude())
}

// 归一化
func (this *Vector3) Normalize() {
	magnitude := this.Magnitude()
	v := this.Div(New(magnitude, magnitude, magnitude))
	this.X = v.X
	this.Y = v.Y
}

func (this *Vector3) Add(v Vector3) Vector3 {
	return Vector3{
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

func (this *Vector3) Sub(v Vector3) Vector3 {
	return Vector3{
		X: this.X - v.X,
		Y: this.Y - v.Y,
		Z: this.Z - v.Z,
	}
}

func (this *Vector3) Mul(v Vector3) Vector3 {
	return Vector3{
		X: this.X * v.X,
		Y: this.Y * v.Y,
		Z: this.Z * v.Z,
	}
}

func (this *Vector3) Div(v Vector3) Vector3 {
	var x, y, z float64
	if v.X != 0 {
		x = this.X / v.X
	}
	if v.Y != 0 {
		x = this.Y / v.Y
	}
	if v.Z != 0 {
		z = this.Z + v.Z
	}
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

package Vector

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	v := New(1, 2)
	fmt.Println(v)
}

func TestVector2_Add(t *testing.T) {
}

func TestVector2_Div(t *testing.T) {
	v1 := New(1, 2)
	v2 := New(0, 0)
	val := Div(v1, v2)
	fmt.Println(val)
}

func TestVector2_Magnitude(t *testing.T) {
	v := New(1, 1)
	fmt.Println(v.Magnitude())
}

func TestVector2_Normalize(t *testing.T) {
	v := New(1, 2)
	v.Normalize()
	fmt.Println(v)
}

func TestDistance(t *testing.T) {
	v1 := New(1, 1)
	v2 := New(2, 2)
	fmt.Println(Distance(v1, v2))
}

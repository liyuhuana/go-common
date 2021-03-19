package vector3

import (
	"fmt"
	"testing"
)

func TestNewVector3(t *testing.T) {
	v := New(1, 2, 3)
	fmt.Println(v)
}

func TestDistance(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(3, 2, 1)
	fmt.Println(Distance(v1, v2))
}

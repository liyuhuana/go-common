package math_util

import (
	"fmt"
	"math/rand"
	"strconv"
)

func RandRange(min int, max int) int {
	if min <= 0 || max <= 0 {
		return 0
	}
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min)
}

//int32转到string
func Int32ToStr(src int32) string {
	dst := strconv.Itoa(int(src))
	return dst
}

//int64转到string
func Int64ToStr(src int64) string {
	fmt.Sprintf("%v",src)
	dst := strconv.FormatInt(src, 10)
	return dst
}

package int_algo

import (
	"math"
)

//7. 整数反转

func Reverse(x int) int {
	res := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		temp := x % 10
		if res > (math.MaxInt32-temp)/10 {
			return 0
		}
		x /= 10
		res = res*10 + temp
	}
	res *= sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

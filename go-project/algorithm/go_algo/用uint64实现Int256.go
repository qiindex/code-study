package go_algo

import (
	"math"
	"math/bits"
)

type Int256 struct {
	data []uint64
}

func (int256 *Int256) init() *Int256 {

	return &Int256{
		data: []uint64{0, 1, 2, 3},
	}
}

func (a *Int256) add(b Int256) Int256 {
	carry := uint64(0)
	//carryBit := uint(0)

	sums := Int256{data: make([]uint64, len(a.data))}
	for i := 3; i >= 0; i-- {
		sum := carry + a.data[i] + b.data[i]

		if sum < b.data[i] || sum < a.data[i] {
			carry = 1
		} else {
			carry = 0
		}
		sum = sum & math.MaxUint64
		sums.data[i] = sum
		//sums.data[i]), carryBit = bits.Add(uint(a.data[i]), uint(b.data[i]), uint(carry))更优雅
	}

	//
	if carry == 1 {
		panic("Int256 overflow")
		//return Int256{data: []uint64{0, 0, 0, 0}}
	}
	return sums
}

// Add 实现256位加法（使用bits.Add处理进位）

func (a *Int256) sub(b Int256) Int256 {
	carry := uint64(0)
	sums := Int256{data: make([]uint64, len(a.data))}
	for i := 3; i >= 0; i-- {
		if a.data[i] >= b.data[i]+carry {
			sums.data[i] = a.data[i] - b.data[i] - carry
			carry = 0
		} else {
			sums.data[i], carry = bits.Sub64(a.data[i], b.data[i], carry) //这种方式更优雅
			//sums.data[i] = a.data[i] + (math.MaxUint64) - b.data[i] - carry //这样容易理解；直接依赖 uint64 的回绕特性，相当于0-1：1<<64 - 1
			//carry = 1
		}
	}

	if carry == 1 {
		// 这里应该返回一个表示下溢的值，而不是全0
		// 可以panic或者返回一个特殊值
		panic("Int256 underflow")
	}
	return sums
}

func main() {
	a := &Int256{}
	a = a.init()

	b := Int256{data: []uint64{0, 0, 0, 1}} // 1

	// 测试加法
	sum := a.add(b)
	println("sum:", sum.data[0], sum.data[1], sum.data[2], sum.data[3])

	// 测试减法
	diff := a.sub(b)
	println("diff:", diff.data[0], diff.data[1], diff.data[2], diff.data[3])
}

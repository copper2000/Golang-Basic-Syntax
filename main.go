package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Bool
	var myBool bool = true // false
	fmt.Println(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	// string
	var myString string = "hello"
	fmt.Println(myString)

	// int
	var myInt int = 456
	fmt.Println(myInt)

	// int 8, 16, 32, 64
	// 1. Range

	// Range int8
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)
	// -128 -> 127

	// Range int16
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)
	// -32768 -> 32767

	// Range int32
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)
	// -2147483648 -> 2147483647

	// Range int64
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)
	// -9223372036854775808 -> 9223372036854775807

	// 2. Bits
	fmt.Println("================")

	fmt.Println(bits.OnesCount(math.MaxUint8))
	// 8 bits

	fmt.Println(bits.OnesCount(math.MaxUint16))
	// 16 bits

	fmt.Println(bits.OnesCount(math.MaxUint32))
	// 32 bits

	fmt.Println(bits.OnesCount(math.MaxUint64))
	// 64 bits

}

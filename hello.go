package main

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	m := message.NewPrinter(language.Russian)

	fmt.Println("Integers:")

	m.Printf("MinInt: %d\n", math.MinInt)
	m.Printf("MaxInt: %d\n", math.MaxInt)

	m.Printf("MinInt8: %d\n", math.MinInt8)
	m.Printf("MaxInt8: %d\n", math.MaxInt8)

	m.Printf("MinInt16: %d\n", math.MinInt16)
	m.Printf("MaxInt16: %d\n", math.MaxInt16)

	m.Printf("MinInt32: %d\n", math.MinInt32)
	m.Printf("MaxInt32: %d\n", math.MaxInt32)

	m.Printf("MinInt64: %d\n", math.MinInt64)
	m.Printf("MaxInt64: %d\n", math.MaxInt64)

	fmt.Println("Unsigned Integers:")

	var maxUint uint64 = math.MaxUint //без указания типа не инициализируется, пишет overflow
	m.Printf("MaxUint: %d\n", maxUint)
	m.Printf("MaxUint8: %d\n", math.MaxUint8)
	m.Printf("MaxUint16: %d\n", math.MaxUint16)
	m.Printf("MaxUint32: %d\n", math.MaxUint32)
	m.Printf("MaxUint64: %d\n", uint64(math.MaxUint64))

}

package main

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	printer := message.NewPrinter(language.Russian)

	fmt.Println("Integers:")

	printer.Printf("MinInt: %d\n", math.MinInt)
	printer.Printf("MaxInt: %d\n", math.MaxInt)

	printer.Printf("MinInt8: %d\n", math.MinInt8)
	printer.Printf("MaxInt8: %d\n", math.MaxInt8)

	printer.Printf("MinInt16: %d\n", math.MinInt16)
	printer.Printf("MaxInt16: %d\n", math.MaxInt16)

	printer.Printf("MinInt32: %d\n", math.MinInt32)
	printer.Printf("MaxInt32: %d\n", math.MaxInt32)

	printer.Printf("MinInt64: %d\n", math.MinInt64)
	printer.Printf("MaxInt64: %d\n", math.MaxInt64)

	fmt.Println("Unsigned Integers:")

	printer.Printf("MaxUint: %d\n", uint64(math.MaxUint))
	printer.Printf("MaxUint8: %d\n", math.MaxUint8)
	printer.Printf("MaxUint16: %d\n", math.MaxUint16)
	printer.Printf("MaxUint32: %d\n", math.MaxUint32)
	printer.Printf("MaxUint64: %d\n", uint64(math.MaxUint64))

	printer.Printf("byte (alias for uint8): %d\n", math.MaxUint8)

	printer.Printf("min none zero float32: %f\n", math.SmallestNonzeroFloat32)
	printer.Printf("Max float32: %f\n", math.MaxFloat32)

	printer.Printf("min none zero float64: %f\n", math.SmallestNonzeroFloat64)
	printer.Printf("float64: %f\n", math.MaxFloat64)

}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(4))
	fmt.Printf("%b\n", 1000)         // 1111101000
	fmt.Printf("%o\n", 1000)         // 1750
	fmt.Printf("%O\n", 1000)         // 0o1750
	fmt.Printf("%d\n", 1000)         // 1000
	fmt.Printf("%x\n", 1000)         // 3e8
	fmt.Printf("%X\n", 1000)         // 3E8
	fmt.Printf("%g\n", 123.456789)   // 123.456789
	fmt.Printf("%s\n", "str")        // str
	fmt.Printf("%t\n", true)         // true
	fmt.Printf("%e\n", 4+1i)         // 4.000000e+00+1.000000e+00i
	fmt.Printf("%.2f\n", 123.456789) // 123.45
}

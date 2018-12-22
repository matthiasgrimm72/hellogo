package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x)
	zn := z
	i := 0
	delta := 1.0

	for i = 0; delta > 0.000001; i++ {
		zn = z - (z*z-float64(x))/(2*z)
		delta = 1 - (zn / z)
		z = zn
		fmt.Println(delta)
	}

	fmt.Printf("%v Iterationen\n", i)

	return z
}

func main() {
	x := 214132311122.0
	fmt.Printf("Meine Wurzel: %v\n", Sqrt(x))
	fmt.Printf("Mathe Wurzel: %v", math.Sqrt(x))
}

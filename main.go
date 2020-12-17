package main

import (
	"fmt"
	"math"
)

func inputABC() (a, b, c float64) {
	fmt.Printf("input A: ")
	fmt.Scanf("%f\n", &a)
	fmt.Printf("input B: ")
	fmt.Scanf("%f\n", &b)
	fmt.Printf("input C: ")
	fmt.Scanf("%f\n", &c)
	return
}

func discriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - (4 * a * c)
}

func roots(a, b, d float64) (x1, x2 float64) {
	x1 = (-b + math.Sqrt(d)) / (2 * a)
	x2 = (-b - math.Sqrt(d)) / (2 * a)
	return
}

func main() {
	a, b, c := inputABC()
	d := discriminant(a, b, c)

	if d < 0 {
		fmt.Println("non roots")
	} else if d == 0 {
		x, _ := roots(a, b, d)
		fmt.Printf("1 root: %f\n", x)
	} else {
		x1, x2 := roots(a, b, c)
		fmt.Println("have 2 roots:")
		fmt.Printf("root number 1: %f\n", x1)
		fmt.Printf("root number 2: %f\n", x2)
	}
}

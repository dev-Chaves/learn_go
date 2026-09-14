package main

import (
	"fmt"
	"math/rand"

	"rsc.io/quote"
)

func sum(x int, y int) int {
	return x + y
}

func main() {

	peso := 12.7

	var a float32 = 0.2

	var b float32 = 0.3

	var c float64 = 0.2

	var d float64 = 0.3

	fmt.Println("------------------")

	if (a * b) == 0.6 {
		fmt.Println("verdade")
	} else {
		fmt.Println("Falso!!!")
	}

	fmt.Println("------------------")

	if (c * d) == 0.6 {
		fmt.Println("verdade")
	} else {
		fmt.Println("Falso!!!")
	}

	fmt.Println("------------------")

	//if total.Mul(e, f) == 0.6 {
	//	fmt.Println("verdade")
	//} else {
	//	fmt.Println("Falso!!!")
	//}

	fmt.Println("------------------")

	fmt.Println(a * b)

	fmt.Println("------------------")

	fmt.Printf("Tipo da váriavel: %T\n", peso)

	fmt.Println(rand.Intn(5))

	fmt.Println(c)

	fmt.Println(sum(10, 10))

	fmt.Println(quote.Go())

}

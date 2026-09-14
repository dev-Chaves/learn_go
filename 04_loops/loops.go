package main

import (
	"fmt"
)

func main() {

	var teste []int = []int{}

	teste = append(teste, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	for i := 0; i < len(teste); i++ {

		fmt.Println(teste[i])

	}

	i:=0

	for i < len(teste){
		fmt.Println(teste[i])

		i++
	}

	fmt.Println("Final do slice")

}

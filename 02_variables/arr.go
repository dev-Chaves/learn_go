package main

import "fmt"

func main() {

	fmt.Println("---------")

	var teste [3]int = [3]int{1, 2, 3}

	fmt.Println(teste)

	fmt.Println("---------")

	fmt.Printf("%T", teste)

	fmt.Println()

	fmt.Println("---------")

	var slice []string = []string{"string"}

	fmt.Println("Tamanho: ", len(slice), "Capacidade: ", cap(slice))

	slice = append(slice, "kakau", "chaves")

	slice = append(slice, "lua", "casa")

	slice = append(slice, "lua")

	slice = append(slice, "lua")

	fmt.Println("Tamanho: ", len(slice), "Capacidade: ", cap(slice))


}

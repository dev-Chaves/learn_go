package main

import "fmt"
import "errors"

func main() {

	var teste int = 10

	switch teste {
	case 100, 10, 1000:
		fmt.Println("caiu na regra do 1")
	case 200, 20, 2000:
		fmt.Println("caiu na regra do 2")
	default:
		fmt.Println("inválido")
	}

	fmt.Println("-------")

	if teste < 5 {
		fmt.Println("Menor que 5")
	} else {
		fmt.Println("Maior que 5")
	}

	fmt.Println("-----------")

	if err := test(-1); err != nil {
		fmt.Println("Erro ao multiplicar: ", err)
		return
	}

	switch {
	case teste >= 10:
		fmt.Println("Maior ou igual a 10")
	case teste <= 10:
		fmt.Println("Menor ou igual 10")
	}

}

func test(x int) any {
	if x < 0 {
		return errors.New("O valor deve ser maior que 0")
	}

	return x * 10
}

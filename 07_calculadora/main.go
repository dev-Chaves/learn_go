package main

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.Float64("a", 0, "Valor A")

	b := flag.Float64("b", 0, "Valor B")

	operacao := flag.String("op", "soma", "Calcular valores")

	flag.Parse()

	switch *operacao {
	case "soma":
		Executar(Adicao{}, *a, *b)
	case "sub":
		Executar(Subtracao{}, *a, *b)
	case "mult":
		Executar(Multiplicacao{}, *a, *b)
	case "div":
		Executar(Divisao{}, *a, *b)
	default:
		fmt.Println("Operação não encontrada")
	}
}

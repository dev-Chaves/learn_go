package main

import (
	"errors"
	"fmt"
)

type Operacao interface {
	Calcular(a, b float64) (float64, error)
}

type Adicao struct {
}

type Subtracao struct {
}

type Divisao struct {
}

type Multiplicacao struct {
}

func (Adicao) Calcular(a, b float64) (float64, error) {
	return a + b, nil
}

func (Subtracao) Calcular(a, b float64) (float64, error) {
	return a - b, nil
}

func (Divisao) Calcular(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Divisor não pode ser 0")
	}

	return a / b, nil

}

func (Multiplicacao) Calcular(a, b float64) (float64, error) {

	return a * b, nil

}

func Executar(operacao Operacao, a, b float64) {

	resultado, err := operacao.Calcular(a, b)

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	fmt.Printf("Resultado da operação: %.2f\n", resultado)

}

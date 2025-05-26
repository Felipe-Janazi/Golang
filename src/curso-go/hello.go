package main

import "fmt"

type Carro struct {
	Nome string
}

func(c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "Azera",
	}

	carro.andar()

	// Uma função que devolve uma função que devolve um inteiro
	resultado := func(x ...int)  func() int {

		res := 0

		// Um looping de valores V, dentro do range informado de X
		// Sendo um array, assim somando todos os seus valores
		for _, v := range x {
			res += v
		}
		return func() int {
			return res + res
		}
	}

	// Chamado da função, devolvendo uma função e devolvendo um inteiro
	fmt.Println(resultado(10,10,10,10,10)())
}

// Retorno por nome da variáveis
func Soma(x int, y int) (result int) {

	result = x + y
	return
}

// Soma de um array de inteiros sendo passados
func somaTudo(x ...int) int {
	resultado := 0

	// Um looping de valores V, dentro do range informado de X
	// Sendo um array, assim somando todos os seus valores
	for _, v := range x {
		resultado += v
	}

	return resultado
}

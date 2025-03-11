package main

import (
	"fmt"
)


func main() {
	// helloWorld()

	// fmt.Println("")

	// fmt.Println("Informe um número impar ou par:")
	// var numero int
	// fmt.Scan(&numero)
	// fmt.Println(parOuImpar(numero))

	// fmt.Println("")

	// var numero1 int
	// var numero2 int
	// fmt.Println("Informe um número para a soma:")
	// fmt.Scan(&numero1)
	// fmt.Println("Informe outro número para a soma:")
	// fmt.Scan(&numero2)
	// fmt.Println(soma(numero1, numero2))

	// fmt.Println("")

	// var fatorial int
	// fmt.Println("Informe um número para cálculo de fatorial:")
	// fmt.Scan(&fatorial)
	// fmt.Println(calculoFatorial(fatorial))

	// fmt.Println("")

	var numeroTabuada int
	fmt.Println("Informe um número para realizarmos a tabuada até o 10:")
	fmt.Scan(&numeroTabuada)
	fmt.Println(tabuada(numeroTabuada))
}
func helloWorld() {
	fmt.Println("Hello, World!")
}

func parOuImpar(numero int) string{

	if (numero % 2 == 0){
		frase := fmt.Sprintf("O número %d é um número par", numero)

		return frase
	}

	frase := fmt.Sprintf("O número %d é um número impar", numero)

		return frase
}

func soma(numero1 int, numero2 int) string {

	soma := numero1 + numero2
	fraseSoma :=  fmt.Sprintf("A soma dos números é: %d + %d = %d", numero1, numero2, soma)
	return fraseSoma
}

func calculoFatorial(fatorial int) string {

	calculo := fatorial
	for i := fatorial - 1; i > 0; i-- {
		calculo *= i
	}

	frase := fmt.Sprintf("O fatorial do número %d é de %d.", fatorial, calculo)
	return frase
}

func tabuada(numero int) string {

	var frase string
	for i := 0; i <= 10; i++ {
		
		calculo := numero * i;

		frase += fmt.Sprintf("%d * %d = %d \n", numero, i, calculo)
	}

	return frase
}
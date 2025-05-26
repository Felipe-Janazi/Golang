package main

import "fmt"

type Carro struct {
	Name string
}

// Sem o ponteiro mudamos o nome somente durante a execução dessa função
// Quando apontamos o mesmo local de memória o valor de *Carro mudamos o nome diretamente na struct
func (c *Carro) andou() {
	c.Name = "Teste"
	fmt.Println(c.Name, "andou")
}

func main() {
	carro := Carro{
		Name: "Azera",
	}

	carro.andou()
	fmt.Println(carro.Name)



	variavel := 10

	abc(&variavel)

	fmt.Println(variavel)
	// a := 10 

	// // Atribuindo o endereço de memória
	// var ponteiro *int = &a
	// fmt.Println(*ponteiro)

	// // & Apontando para o endereço em memória da variável informada
	//  * Apontando para o valor dessa variável
	// Alterando o valor do endereço de memória tudo que aponte para o mesmo endereço é alterado
}

func abc( a *int){
	*a = 200
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome, Email, CPF string
}

// Dessa forma tornamos a func um método atrelando ele ao cliente pelo "c Cliente"
func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método", c.Nome)
}

type ClienteInternacional struct {
	// Assim podemos pegar tudo do cliente anterior e adicionar somente o desejado neste novo cliente
	Cliente
	// Dessa forma numa exibição de Json podemos mudar o nome do campo que será apresentado
	Pais string `json:"pais"`
}

func main() {

	cliente := Cliente{
		Nome:  "Felipe",
		Email: "felipe@gmail",
		CPF:   "12312312312",
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"Gabriel", "gabriel@gmail", "12345678900"}
	fmt.Printf("Nome: %s.\nEmail: %s.\nCPF: %s.\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		// Declarando o cliente usando uma struct que utiliza uma outra struct dentro
		Cliente: Cliente{"Janazi", "janazi@gmail", "11111111111"},
		Pais: "Eua",
	}
	// Ao inves de usarmos Cliente3.Cliente.Nome podemos usarm diretamente Nome, pois o go interpreta como parte do cliente
	fmt.Printf("Nome: %s.\nEmail: %s.\nPais: %s.\n", cliente3.Nome, cliente3.Email, cliente3.Pais)

	// Podemos usar a função "Exibe" mesmo que ela pertença a outra struct, afinal a struct do cliente 3 está utilizando a strcut Cliente
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	// Dessa forma transformamos o cliente3 em json
	clienteJson, err := json.Marshal(cliente3)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Para exibir corretamente o json precisamos transforma-lo em uma string
	fmt.Println(string(clienteJson))

	// Para adicionarmos os valores de json após fazemos:
	jsonCliente4 := `{"Nome":"Ferreira","Email":"ferreira@gmail","CPF":"11111111111","Pais":"jurandir"}`
	cliente4 := ClienteInternacional{}
	// Vamos transformar o json em bytes e colocar no endereço de memória que o cliente4 aponta, pois se não colocarmos iremos mudar os valores somente para este caso de ação
	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}

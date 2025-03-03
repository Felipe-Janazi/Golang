package main

import (
	"api_rest/database"
	// "api_rest/models"
	"api_rest/routes"
	"fmt"
)



func main() {

	// Criando personalidades direto dentro da lista que está nas Models
	// models.Personalidades = []models.Personalidade {
	// 	{Nome: "Nome 1", Historia: "Historia 1", Id: 1},
	// 	{Nome: "Nome 2", Historia: "Historia 2", Id: 2},
	// }

	database.ConectaComBancoDeDados()
	
	fmt.Println("Iniciando o servidor rest com o Go")

	// Pegando todas as rotas
	routes.HandleRequest()
}
package main

import (
	"api-rest-gin/database"
	"api-rest-gin/routes"
)

// import "github.com/gin-gonic/gin"

func main() {
	database.ConectaComBancoDeDados()
	// Temos a lista no models que é igual a uma lista de alunos

	routes.HandleRequests()
}
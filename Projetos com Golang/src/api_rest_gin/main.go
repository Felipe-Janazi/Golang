package main

import (
	"api-rest-gin/models"
	"api-rest-gin/routes"
)

// import "github.com/gin-gonic/gin"

func main() {
	// Temos a lista no models que é igual a uma lista de alunos
	models.Alunos = []models.Aluno{
		{Nome: "Felipe", CPF: "12312312312", RG: "12312312312"},
		{Nome: "Ana", CPF: "12312312313", RG: "12312312313"},
	}
	routes.HandleRequests()
}
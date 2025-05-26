package controllers

import (
	"api-rest-gin/database"
	"api-rest-gin/models"

	"github.com/gin-gonic/gin"
)

// Com essa função conseguimos inserir na tela um usuário
func ExibeTodosOsAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	// Para pegarmos parametros que estão na routes usamos a nomenclatura abaixo
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + " tudo beleza?",
	})
}
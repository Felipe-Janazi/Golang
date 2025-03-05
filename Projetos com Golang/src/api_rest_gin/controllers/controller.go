package controllers

import (
	"api-rest-gin/models"

	"github.com/gin-gonic/gin"
)

// Com essa função conseguimos inserir na tela um usuário
func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}


func Saudacao(c *gin.Context){
// Para pegarmos parametros que estão na routes usamos a nomenclatura abaixo 
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:" : "E ai " + nome + " tudo beleza?",
	})
}
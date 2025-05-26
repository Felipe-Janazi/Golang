package controllers

import (
	"api-rest-gin/database"
	"api-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	// Pegar todo o corpo da requisição e empacotar para json
	if err := c.ShouldBindJSON(&aluno); err != nil {
		// Caso tenha ocorrido algum erro entrará neste caso, que devolve um BadRequest
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Caso de tudo certo criamos o aluno e devolvemos uma mensagem de criação
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}







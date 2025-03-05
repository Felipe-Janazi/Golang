package controllers

import (
	"api-rest-gin/database"
	"api-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletaAluno(c *gin.Context) {

	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso"})
}

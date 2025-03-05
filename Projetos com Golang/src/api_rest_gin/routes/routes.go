package routes

import (
	"api-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	// Rota com método e leva para uma função
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.Run(":8000")
}
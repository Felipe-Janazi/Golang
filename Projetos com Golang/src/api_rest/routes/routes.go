package routes

import (
	"api_rest/controllers"
	"api_rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	// Utilizando a rota do pacote gorilla
	// Instalação: go get -u github.com/gorilla/mux
	r := mux.NewRouter() 

	// Desta forma usamos o middleware para executar uma função sempre que usamos o mux = r
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("GET")
	// Informamos o nome do campo que será passado = {id}
	r.HandleFunc("/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")

	// Expondo a porta 8000 e interligando junto ao frontend em react
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
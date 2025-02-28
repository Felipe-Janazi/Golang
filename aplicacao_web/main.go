package main

import (
	"net/http"
	"aplicacao_web/routes"
)


func main () {

	routes.CarregaRotas()
	
	// Damos a porta 8000 para ser executada
	http.ListenAndServe(":8000", nil)
}


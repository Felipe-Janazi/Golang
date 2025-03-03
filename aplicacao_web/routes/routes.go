package routes

import (
	"aplicacao_web/controllers"
	"net/http"
)

func CarregaRotas() {
	// Sempre que tiver uma requisição para / quem irá atender é o index
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}

// localhost:8000/new
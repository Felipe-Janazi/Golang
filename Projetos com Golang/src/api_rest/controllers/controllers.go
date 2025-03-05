package controllers

import (
	"api_rest/database"
	"api_rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
)



func Home(w http.ResponseWriter, r *http.Request) {
	// Printando na página raiz
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	// Quero ter um novo encoder do tipo json, falando que são as personalidades
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// Informamos que queremos as informações que vierem da requisição
	vars := mux.Vars(r)

	//  Qual o valor que queremos da requisição que está definido nas routes
	id := vars["id"]

	// Criamos um objeto personalidade com base na struct
	var personalidade models.Personalidade

	// Desta forma encontramos a personalidade que queremos editar
	database.DB.First(&personalidade, id)

	// Encodamos na tela a personalidade escolhida
	json.NewEncoder(w).Encode(personalidade)

	// Cada personalidade na variavel, tendo um range de toda a lista de Personalidades
	// for _, personalidade := range models.Personalidades {

	// 	// Converte o Id da struct para String e compara
	// 	if strconv.Itoa(personalidade.Id) == id {

	// 		// Se for igual vamos exibir Encodando na tela
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
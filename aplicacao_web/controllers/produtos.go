package controllers

import (
	"aplicacao_web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Aqui estamos criando uma varivel que vai buscar todos os templates que possuimos na pasta templates e que terminem com html
var temp = template.Must(template.ParseGlob("templates/*.html"))

// Toda vez que tiver uma requisição iremos escrever para a pessoa e apontando para o request
func Index (w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()

	// Informamos quem consegue escrever, quem será mostrado, afinal informamos no HTML onde começa e termina o Index e informamos o que vamos enviar para a página HTML, neste caso estamos enviando uma lista de produtos para que seja utilizada no HTML. 
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	//  Se o método for post
	if r.Method == "POST" {

		// Pegamos os valores do campo name
		nome := r.FormValue("nome")	
		descricao := r.FormValue("descricao")	
		preco := r.FormValue("preco")	
		quantidade := r.FormValue("quantidade")	

		//  Convertemos, pois eles vem como string
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		//  Convertemos 
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		//  E agora que esta convertido enviamos para as models que irão inserir no BD
		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	//  Informamos que deu tudo certo e vamos retornar a página raiz do meu projeto e informamos o código que significa isso 
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request){

	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id para Int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		//  Convertemos 
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.AtualizaProduto(idConvertidoParaInt, quantidadeConvertidaParaInt, nome,  descricao, precoConvertidoParaFloat, )

		http.Redirect(w, r, "/", 301)
	}
}
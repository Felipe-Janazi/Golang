package db

import (
	"database/sql"

	// Como usamos o postgre durante a execução precisamos usar o _ mostrando que mesmo que não estamos a utilizando ela irá continuar existindo no nosso código, diferente das demais bibliotecas que precisam já estar ligadas a alguma parte do projeto
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// Passmaos todas as informações que podem ser obtidas acessando com botão direito o PostreSQL > connection > e informamos aqui
	conexao := "user=postgres dbname=alura_loja password=fg060306 host=localhost sslmode=disable"

	// Informamos que queremos abris o postgres passando a variavel com as informações do Banco
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	// Retornamos o resultado da abertura ao Banco de Dados
	return db
}

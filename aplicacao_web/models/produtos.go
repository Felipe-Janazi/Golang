package models

import (
	"aplicacao_web/db"
)


type Produto struct {
	Id, Quantidade int
	Nome, Descricao string
	Preco float64
}

func BuscaTodosOsProdutos () []Produto{

	db := db.ConectaComBancoDeDados()

	// Fazemos uma query para puxar todos os dados do Banco
	selectDeTodosOsProdutos, err := db.Query("Select * from produtos order by id asc")

	if err != nil {
		panic(err.Error)
	}

	// Instanciamos um produto e uma lista de produtos que vamos usar
	p := Produto{}
	produtos := []Produto{}

	//  Fazemos um for que vai até a última linhda do select
	for selectDeTodosOsProdutos.Next(){
		// Criamos as variaveis para usar no select
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//  Usamos o scan para escanear cada linha do select e salvar as informações em variaveis criadas acima
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error)
		}

		// Colocamos cada informação no produto que instanciamos
		p.Id = id 
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		// Salvamos esse produto dentro da nossa lista de produtos 
		produtos = append(produtos, p)

	}
	
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {

	// Fazendo a conexão com o Banco de Dados para inserir 
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string ) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	// Executa esse script com o seguinte id/campo
	deletarOProduto.Exec(id)

	defer db.Close()
}

func EditaProduto(id string ) Produto {

	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64){

	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
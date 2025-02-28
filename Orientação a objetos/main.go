package main

import (
	"fmt"

	// Para criarmos um pacote próprio usamos "go mod init (nome da pasta que possui todos os pacotes)"
	// Desta forma criamos um go.mod e depois colocamos o nome da pasta raiz seguido pela pasta do nosso novo package.
	// O package precisa ter as variaveis da estrutura todas em letras maisculas
	// Podemos apelidar esse pacote colocando alguma letra ou palavra antes dele
	// c "orientacao_a_objetos/contas"
	"orientacao_a_objetos/contas"
	// "orientacao_a_objetos/clientes"
)
  
// Aqui nos chamamos a interface para fazer a validação com base na conta do cliente, se a conta for corrente e tiver exatamente o método desejado ela passa e executa o saque na conta especifica
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// Aqui nos usamos a interface para verificar a conta do cliente, se ela possui ou não o método desejado
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	
	contaLuiza := contas.ContaCorrente{}
	contaLuiza.Depositar(100)

	// Aqui passamos o endereço da conta que desejamos realizar o pagamento 
	PagarBoleto(&contaDenis, 60)
	PagarBoleto(&contaLuiza, 60)

	fmt.Println(contaDenis)
	fmt.Println(contaLuiza)
	







// DUAS FORMAS DE CRIAR CONTAS E USUARIOS E PREVENINDO QUE O SALDO CRIADO NÃO SEJA NEGATIVO E PODENDO CONSULTAR E DEPOSITAR VALORES

	// clienteBruno := clientes.Titular{
	// 	Nome: "Bruno",
	// 	CPF: "123.123.123.12",
	// 	Profissao: "Desenvolvedor"}

	// 	contaDoBruno := contas.ContaCorrente {
	// 		Titular: clienteBruno,
	// 		NumeroAgencia: 123,
	// 		NumeroConta: 123456,
	// 	}

	// 	contaDoBruno.Depositar(100)
	// 	fmt.Println(contaDoBruno.ObterSaldo())

		
		// contaDoBruno := contas.ContaCorrente{
			// 	Titular: clientes.Titular{
				// 		Nome: "Bruno",
				// 		CPF: "123.111.123.12",
				// 		Profissao: "Desenvolvedor"},
				// 	NumeroAgencia: 123,
				// 	NumeroConta: 123456,
				// 	Saldo: 100,
				// }
				
				
				
		// fmt.Println(contaDoBruno)





	//  EXEMPLO PARA TRANAFERÊNCIA DE VALORES ENTRE OBJETOS 

	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// // Para transferir precisamos informar que etsamos enviando para o endereço da conta do Gustavo usando um &
	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)






	// EXEMPLO DE DEPOSITO E SAQUES

//  contaDaSilvia := ContaCorrente{}
//  contaDaSilvia.titular = "Silvia"
//  contaDaSilvia.saldo = 500

// fmt.Println(contaDaSilvia.saldo)

// fmt.Println(contaDaSilvia.Sacar(-100))
// fmt.Println(contaDaSilvia.saldo)

// //  Criação rápida de variáveis com cada uma sendo responsável por um retorno
// status, valor := contaDaSilvia.Depositar(500)
// fmt.Println(status, valor)






	//  EXEMPLO DE CRIAÇÃO DE OBJETOS 

// Passamos os nomes dos campos quando não precisamos preencher todos
// contaDoGuilherme := ContaCorrente{titular: "Guilherme",
// 	saldo: 125.5}	

// 	// E passamos sem o nome dos campos quando queremos preencher os campos em ordem
// 	contaDaBruna := ContaCorrente{"Bruna", 589, 123456, 125.5}

// 	fmt.Println(contaDaBruna)

// 	// Aqui apontamos que estamos alocando um espaço na ContaCorrente, por isso usamos *
// 	var contaDaCris *ContaCorrente
// 	contaDaCris = new(ContaCorrente)
// 	contaDaCris.titular = "Cris"
// 	contaDaCris.saldo = 500

// Com este comando estamos vendo o conteúdo da conta, se tirarmos o * estamos verificando o local dessa conta e depois o conteúdo da mesma, o que afeta em uma comparação de contas
// 	fmt.Println(*contaDaCris)
}

package main

import "fmt"
  
// Criação de uma estrutura que objetos irão seguir
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Informamos que vai apontar para a conta que chamou a função
// O valor do saque e o retorno para essa função
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente."
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito > 0 {
	c.saldo += valorDoDeposito;
	return "Deposito realizado com sucesso!"
	}

	return "Valor do depósito menor que 0"
}

func main() {
 contaDaSilvia := ContaCorrente{}
 contaDaSilvia.titular = "Silvia"
 contaDaSilvia.saldo = 500

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(-100))
fmt.Println(contaDaSilvia.saldo)

contaDaSilvia.Depositar(500)
fmt.Println(contaDaSilvia.saldo)



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

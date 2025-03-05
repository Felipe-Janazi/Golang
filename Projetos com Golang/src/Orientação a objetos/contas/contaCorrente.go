package contas

import "orientacao_a_objetos/clientes"

// Criação de uma estrutura que objetos irão seguir
type ContaCorrente struct {
	Titular       clientes.Titular
	// Se as variáveis são de um mesmo tipo podemos apenas colocar uma virgula por nome e colocar seu tipo no final
	NumeroAgencia, NumeroConta int
	// Quando colocamos uma variavel do struct minuscula é como se colocassemos o private, somente esta própria página pode mexer com este elemento
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

	return "saldo insuficiente."
}

// Aqui informamos que vamos lidar com mais de um retorno desta função
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do depósito menor que 0", c.saldo
}

// Para apontarmos a conta que vai receber a transferência usamos *, desta forma apontamos o destino desta conta que recebe a transferência
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia

		contaDestino.Depositar(valorDaTransferencia)

		return true
	}

	return false

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
package contas

import "orientacao_a_objetos/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "saldo insuficiente."
}

// Aqui informamos que vamos lidar com mais de um retorno desta função
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do depósito menor que 0", c.saldo
}

// Para apontarmos a conta que vai receber a transferência usamos *, desta forma apontamos o destino desta conta que recebe a transferência
func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {

	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia

		contaDestino.Depositar(valorDaTransferencia)

		return true
	}

	return false

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}


package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	// Enquanto esse for estiver lendo o outro fica parado esperando que seja lido para continuar
	for res := range msg {
		fmt.Println("Worker: ", workerId, "Msg: ", res) 
		time.Sleep(time.Second)
	}
}

func main() {

	msg := make(chan int)

	// Dessa forma enquanto o worker um ta pegando uma mensagem e ta aguardando ser lido o worker dois também pegou uma mensagem e já enviou, assim tornando mais rápido, sem ter que aguardar o work um ser lido e enviar, como um escalonamento horizontal
	go worker(1,msg)
	go worker(2,msg)
	// Dessa forma podemos criar vários para que eles lidem com filas de requisições, afinal cada um usa somente 2k de memória
	go worker(2,msg)
	go worker(2,msg)

	// Esse looping só pode continuar quando a msg tiver sido lida
	for i := 0; i < 10; i++{
		msg <- i
	}
}
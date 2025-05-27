package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string)

	go func() {
		// Neste caso ele vai esperar para atribuir na variavel hello e por isso irá cair no default no select
		time.Sleep(time.Second * 2)
		hello <- "Hello world"
	}()


	// time.Sleep(time.Second * 3)

	select {
		// Informa que caso chegue algum valor para o meu X ele irá printar
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("Default")
	}
}
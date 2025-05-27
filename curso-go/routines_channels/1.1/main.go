package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Começou")

	// Go trabalha de forma cooperativa porém ele também consegue trabalhar de forma preempitiva
	go func () {
		for {

		}
	} ()

	time.Sleep(time.Second)
	fmt.Println("Terminou")
}
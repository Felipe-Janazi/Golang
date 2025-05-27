package main

import "fmt"

// Thred 1
func main() {

	// T1 <-> T2
	hello := make(chan string)

	// Thred 2
	go func() {
		hello <- "Hello world"
	}()

	// Ao hello receber valor a result teve seu valor alterado para igual ao hello
	result := <- hello

	fmt.Println(result)
}

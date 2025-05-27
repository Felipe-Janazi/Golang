package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				continue
			}
		}
	}()

	fmt.Sprintln("Aguardando")

	// Enquanto esse forever não receber um valor ele vai travar a execução
	<- forever
}
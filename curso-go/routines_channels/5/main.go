package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan int)

	go func() {

		i := 0

		// Esse for só vai preencer o queue quando o queue for rodado 
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()
	
	// Enquanto esse não ler o queue o for fica parado aguardando
	for x := range queue {
		fmt.Println(x)
	}
}
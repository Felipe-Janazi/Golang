package main

import (
	"errors"
	"fmt"
)

func main() {
	
	res, _ := Soma(10, 10)

	fmt.Println(res)
}

func Soma(x int, y int) (int, error){
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}



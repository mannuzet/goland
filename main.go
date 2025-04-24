package main

import (
	"fmt"
)

func main() {
	// string
	estoque := map[string]int{
		"coxinha":       10,
		"pão de queijo": 15,
		"refrigerante":  20,
	}

	// string
	estoque["coxinha"] -= 2
	estoque["pão de queijo"] -= 1

	// string
	fmt.Println("Estoque restante:")
	for item, quantidade := range estoque {
		fmt.Printf("%s: %d\n", item, quantidade)
	}
}
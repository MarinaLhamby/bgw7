package main

import "fmt"

func main() {
	palavra := "MELI"
	fmt.Printf("Número de caracteres: %d\n", len(palavra))
	for _, r := range palavra {
		fmt.Printf("Letra: %c\n", r)
	}
}

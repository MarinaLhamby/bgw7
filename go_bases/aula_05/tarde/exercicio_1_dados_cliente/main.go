package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Começando programa para ler o arquivo customers.txt")

	readFile()

	fmt.Println("Execução concluída")

}

func readFile() error {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	return nil
}

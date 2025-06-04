package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Começando programa para ler o arquivo customers.txt")

	data := readFile()

	fmt.Println(string(data))
	fmt.Println("Execução concluída")

}

func readFile() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var fileText string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic("it was not possible to read this file")
		}

		fileText += line
		if err == io.EOF {
			break
		}
	}

	return fileText
}

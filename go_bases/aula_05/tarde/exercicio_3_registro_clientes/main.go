package main

import (
	"exercicio_3_registro_clientes/customer"
	"exercicio_3_registro_clientes/file"
	"fmt"
)

func main() {
	fmt.Println("Começando programa para ler o arquivo customers.txt")
	customerInput := customer.Customer{
		File:        "customers.txt",
		Name:        "John Doe",
		ID:          12345,
		PhoneNumber: "5199999997",
		Address:     "Rua das Flores 123",
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Several errors were detected at runtime")
		}
		fmt.Println("End of execution")

	}()
	data := file.ReadFile()
	customers := file.FileToCustomer(data)

	exists := file.CheckCustomerExistence(customerInput, customers)
	if exists {
		return
	}

	err := customerInput.IsZero()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file.InsertInfoToFile(customerInput)
}

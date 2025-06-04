package file

import (
	"bufio"
	"exercicio_3_registro_clientes/customer"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const fileName = "customers.txt"

func ReadFile() []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic("it was not possible to read this file")
		}

		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}

	return lines
}

func FileToCustomer(data []string) []customer.Customer {
	var customers []customer.Customer

	for _, customerString := range data {
		splittedCustomer := strings.Split(customerString, ";")
		id, err := strconv.Atoi(splittedCustomer[2])
		if err != nil {
			panic("Error converting ID to integer: " + err.Error())
		}

		customers = append(customers, customer.Customer{
			File:        splittedCustomer[0],
			Name:        splittedCustomer[1],
			ID:          int64(id),
			PhoneNumber: splittedCustomer[3],
			Address:     splittedCustomer[4],
		})
	}

	return customers
}

func CheckCustomerExistence(newCustomer customer.Customer, customers []customer.Customer) bool {
	for _, customer := range customers {
		if customer.Compare(newCustomer) {
			panic("client already exists")
		}
	}
	return false
}

func InsertInfoToFile(newCustomer customer.Customer) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer file.Close()

	_, err = file.WriteString(
		"\n" + newCustomer.File + ";" + newCustomer.Name + ";" + strconv.FormatInt(newCustomer.ID, 10) + ";" + newCustomer.PhoneNumber + ";" + newCustomer.Address)
	if err != nil {
		panic("Error writing to file: " + err.Error())
	}

}

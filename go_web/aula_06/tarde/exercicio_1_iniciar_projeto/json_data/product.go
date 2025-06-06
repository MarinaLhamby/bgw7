package json_data

import (
	"encoding/json"
	"os"

	"github.com/bgw7/exercicio_1_iniciar_projeto/product"
)

type ProductData struct {
	Products []product.Product
}

func GetJsonDataFromFile() []product.Product {
	f, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var products []product.Product

	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		panic(err)
	}
	return products
}

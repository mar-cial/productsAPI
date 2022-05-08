package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sku         string  `json:"sku"`
	StoreID     string  `json:"store_id"`
	Inventory   int64   `json:"inventory"`
	Category    string  `json:"category"`
}

func LoadProducts() ([]Product, error) {
	var products []Product
	content, err := ioutil.ReadFile("./data/products.json")
	if err != nil {
		fmt.Println("error reading file")
		return nil, err
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		fmt.Println("error umarshalling file")
		return nil, err
	}

	return products, nil
}

package tutorialjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestEncodeTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Laptop",
		Price:    1000,
		ImageURL: "https://imagesource.com/laptop",
	}

	result, _ := json.Marshal(product)
	fmt.Println(string(result))
}

func TestDecodeTag(t *testing.T) {
	jsonObj := `{"id":"P001","name":"Laptop","price":1000,"image_url":"https://imagesource.com/laptop"}`
	jsonBytes := []byte(jsonObj)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

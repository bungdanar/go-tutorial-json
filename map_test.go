package tutorialjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonObj := `{"id":"P001","name":"Laptop","price":1000,"image_url":"https://imagesource.com/laptop"}`
	jsonBytes := []byte(jsonObj)

	var product map[string]interface{}

	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P001",
		"name":      "Laptop",
		"price":     1000,
		"image_url": "https://imagesource.com/laptop",
	}

	result, _ := json.Marshal(product)
	fmt.Println(string(result))
}

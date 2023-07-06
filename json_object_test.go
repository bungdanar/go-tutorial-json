package tutorialjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Adresses   []Address
}

func TestEncodeObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Danar",
		MiddleName: "Gumilang",
		LastName:   "Putera",
	}

	result, _ := json.Marshal(customer)
	fmt.Println(string(result))
}

func TestDecodeObject(t *testing.T) {
	jsonObj := `{"FirstName":"Danar","MiddleName":"Gumilang","LastName":"Putera"}`
	jsonBytes := []byte(jsonObj)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

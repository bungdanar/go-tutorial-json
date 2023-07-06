package tutorialjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Danar",
		MiddleName: "Gumilang",
		LastName:   "Putera",
		Hobbies:    []string{"Fishing", "Running", "Swimming"},
	}

	result, _ := json.Marshal(customer)
	fmt.Println(string(result))
}

func TestDecodeArray(t *testing.T) {
	jsonArray := `{"FirstName":"Danar","MiddleName":"Gumilang","LastName":"Putera","Hobbies":["Fishing","Running","Swimming"]}`
	jsonBytes := []byte(jsonArray)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestEncodeArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Danar",
		MiddleName: "Gumilang",
		LastName:   "Putera",
		Hobbies:    []string{"Fishing", "Running", "Swimming"},
		Adresses: []Address{
			{
				Street:     "Kapuk",
				Country:    "Jaktim",
				PostalCode: "101",
			},
			{
				Street:     "Harpmul",
				Country:    "Jakpus",
				PostalCode: "100",
			},
		},
	}

	result, _ := json.Marshal(customer)
	fmt.Println(string(result))
}

func TestDecodeArrayComplex(t *testing.T) {
	jsonArray := `{"FirstName":"Danar","MiddleName":"Gumilang","LastName":"Putera","Hobbies":["Fishing","Running","Swimming"],"Adresses":[{"Street":"Kapuk","Country":"Jaktim","PostalCode":"101"},{"Street":"Harpmul","Country":"Jakpus","PostalCode":"100"}]}`
	jsonBytes := []byte(jsonArray)
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestEncodeDirectylyFromArray(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Kapuk",
			Country:    "Jaktim",
			PostalCode: "101",
		},
		{
			Street:     "Harpmul",
			Country:    "Jakpus",
			PostalCode: "100",
		},
	}

	result, _ := json.Marshal(addresses)
	fmt.Println(string(result))
}

func TestDecodeDirectlyToArray(t *testing.T) {
	jsonArray := `[{"Street":"Kapuk","Country":"Jaktim","PostalCode":"101"},{"Street":"Harpmul","Country":"Jakpus","PostalCode":"100"}]`
	jsonBytes := []byte(jsonArray)

	addreses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addreses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addreses)
}

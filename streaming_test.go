package tutorialjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := Customer{}
	decoder.Decode(&customer)

	fmt.Println(customer)
}

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

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

	encoder.Encode(&customer)

	fmt.Println(customer)
}

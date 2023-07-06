package tutorialjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}

func TestEncode(t *testing.T) {
	logJson("Bung")
	logJson(1)
	logJson(true)
	logJson([]string{"Danar", "Gumilang", "Putera"})
}

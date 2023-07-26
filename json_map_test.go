package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONmap(t *testing.T) {
	// suppose ada request data json seperti ini
	jsonRequest := `{"id":"111","name":"contoh produk 1","price":"2000000"}`

	// decode
	jsonBytes := []byte(jsonRequest)
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

	// encode
	jsonEncodeBytes, error := json.Marshal(result)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(jsonEncodeBytes))
}

package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONdecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Contoh First Name","MiddleName":"Contoh Middle Name","LastName":"Contoh Last Name","Umur":17,"Married":false,"NamaAnak":["ayu","sandi","paraswati"]}`
	jsonBytes := []byte(jsonRequest)

	pasien := Pasien{} // ada di file jsn_encode_test.go
	json.Unmarshal(jsonBytes, &pasien)

	fmt.Println(pasien)
	fmt.Println(pasien.FirstName)
	fmt.Println(pasien.MiddleName)
	fmt.Println(pasien.LastName)
	fmt.Println(pasien.Umur)
	fmt.Println(pasien.Married)
}

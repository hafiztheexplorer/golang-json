package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func InputanArray() {

}

/*
================================================================
JSON Array encode decode
================================================================
*/
func TestJSONarrayencode(t *testing.T) {

	pasien := Pasien{
		FirstName:  "Contoh First Name",
		MiddleName: "Contoh Middle Name",
		LastName:   "Contoh Last Name",
		Umur:       17,
		Married:    false,
		NamaAnak:   []string{"ayu", "sandi", "paraswati"}, // JSON array
	}

	byte, error := json.Marshal(pasien)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(byte))

}

func TestJSONarraydecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Contoh First Name","MiddleName":"Contoh Middle Name","LastName":"Contoh Last Name","Umur":17,"Married":false,"NamaAnak":["ayu","sandi","paraswati"]}`
	jsonBytes := []byte(jsonRequest)

	pasien := Pasien{}
	error := json.Unmarshal(jsonBytes, &pasien)
	if error != nil {
		panic(error)
	}

	fmt.Println(pasien)
	fmt.Println("First Name : ", pasien.FirstName)
	fmt.Println("Middle Name : ", pasien.MiddleName)
	fmt.Println("Last Name : ", pasien.LastName)
	fmt.Println("Umur : ", pasien.Umur)
	fmt.Println("Married : ", pasien.Married)
	fmt.Println("Nama Anak : ", pasien.NamaAnak)

}

/*
================================================================
JSON Array complex encode decode
================================================================
*/
func TestJSONarraycomplex(t *testing.T) {
	pasien := Pasien{
		FirstName:  "Contoh First Name",
		MiddleName: "Contoh Middle Name",
		LastName:   "Contoh Last Name",
		Umur:       17,
		Married:    false,
		NamaAnak:   []string{"ayu", "sandi", "paraswati"}, // JSON array
		DAlamat: []Alamat{
			{
				Kota:      "contoh kota 1",
				Kecamatan: "contoh kecamatan 1",
				Desa:      "contoh desa 1",
			},
			{
				Kota:      "contoh kota 2",
				Kecamatan: "contoh kecamatan 2",
				Desa:      "contoh desa 2",
			},
		}, // JSON array complex
	}

	bytepasien, error := json.Marshal(pasien) // encoding JSON
	if error != nil {
		panic(error)
	}
	fmt.Println(string(bytepasien))
	fmt.Println("----------------------------------------------------")

	pointertostruct1 := Pasien{} // pointer to struct structure, so that decode can recognize the data

	error = json.Unmarshal(bytepasien, &pointertostruct1) // decoding JSON from encoding result, and the pointer to struct to recognize the data structure
	if error != nil {
		panic(error)
	}

	fmt.Println(pasien)
	fmt.Println("----------------------------------------------------")
	fmt.Println(pasien.FirstName)
	fmt.Println(pasien.NamaAnak)
	fmt.Println(pasien.DAlamat)

}

package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONTag(t *testing.T) {
	pasien := Pasien{
		FotoDiri:   "http://contoh.com/foto/1.jpeg",
		FirstName:  "contoh nama depan 1",
		MiddleName: "contoh nama tengah 1",
		LastName:   "contoh nama akhir 1",
		Umur:       18,
		NamaAnak:   []string{"1", "2", "3"},
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
		},
	}
	// encoding JSON
	bytepasien, error := json.Marshal(pasien)
	if error != nil {
		panic(error)
	}

	fmt.Println("================================================================")
	fmt.Println("hasil encode: ==================================================")
	fmt.Println("================================================================")
	fmt.Println(string(bytepasien))

	// decoding JSON
	pointerkestructjson := Pasien{}

	error = json.Unmarshal(bytepasien, &pointerkestructjson) // pointer
	if error != nil {
		panic(error)
	}

	fmt.Println("================================================================")
	fmt.Println("hasil decode: ==================================================")
	fmt.Println("================================================================")

	fmt.Println(pasien.FotoDiri)
	fmt.Println(pasien.FirstName, pasien.MiddleName, pasien.LastName)
	fmt.Println(pasien.NamaAnak)
	fmt.Println(pasien.DAlamat)
}

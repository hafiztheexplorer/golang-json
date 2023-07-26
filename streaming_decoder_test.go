package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// streaming decoder & encoder

func TestStreamingDecoder(t *testing.T) {
	// decode
	reader, error := os.Open("Contohfilejson.json")
	if error != nil {
		panic(error)
	}
	decoder := json.NewDecoder(reader)

	pasien := Pasien{} // awas ini linked ke struct Pasien{}, hover to know the position / di json_structs_place.go
	error = decoder.Decode(&pasien)
	if error != nil {
		panic(error)
	}

	fmt.Println(pasien)
	////////////////////////////////////////////////////////////////

	// encode
	writer, error := os.Create("Contohfilejson2.json")
	if error != nil {
		panic(error)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(pasien)

	// perhatian, nama_anak dan d_alamat akan null karena tidak diisi di atas

	fmt.Println(pasien)
}

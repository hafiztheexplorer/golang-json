package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*================================================================
JSON Encode from GO
================================================================*/

// all of this untuk konversi (ENCODE) tipe data yang ada di GO ke JSON
func logJSON(data interface{}) {
	byte, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(byte))
}

func TestMarshal(t *testing.T) {
	logJSON("hello world")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"ohayoo", "konbanwa", "konnichiwa"})
}

/*================================================================
JSON Object Encode from GO
================================================================*/

type Pasien struct {
	FirstName  string
	MiddleName string
	LastName   string
	Umur       int16
	Married    bool
}

func TestJSONObject(t *testing.T) {

	pasien := Pasien{
		FirstName:  "Contoh First Name",
		MiddleName: "Contoh Middle Name",
		LastName:   "Contoh Last Name",
		Umur:       17,
		Married:    false,
	}

	byte, error := json.Marshal(pasien)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(byte))
}

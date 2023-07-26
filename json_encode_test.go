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
func JSONencode(data interface{}) {
	byte, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(byte))
}

func TestMarshal(t *testing.T) {
	JSONencode("hello world")
	JSONencode(1)
	JSONencode(true)
	JSONencode([]string{"ohayoo", "konbanwa", "konnichiwa"})
}

/*================================================================
JSON Object Encode from GO
================================================================*/

func TestJSONObject(t *testing.T) {

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

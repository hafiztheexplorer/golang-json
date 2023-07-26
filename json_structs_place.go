package golangjson

type Pasien struct {
	FotoDiri   string   `json:"foto_diri"`
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Umur       int16    `json:"umur"`
	Married    bool     `json:"married"`
	NamaAnak   []string `json:"nama_anak"`
	DAlamat    []Alamat `json:"d_alamat"`
}

type Alamat struct {
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

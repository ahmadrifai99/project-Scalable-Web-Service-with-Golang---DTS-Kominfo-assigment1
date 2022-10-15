package main

import (
	"fmt"
	"os"
)

type Person struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	orang := []Person{
		{
			"1",
			"Fauzia",
			"Jakarta",
			"Dokter",
			"Pengen Aja",
		}, {
			"2",
			"Rifai",
			"Tangerang Selatan",
			"Astronot",
			"Pengen",
		},
		{
			"3",
			"Fadhl",
			"Jakarta Barat",
			"Pelukis",
			"Ngebet banget",
		}, {
			"4",
			"Asha",
			"Ciledug",
			"Presiden",
			"Pengen",
		}, {
			"5",
			"Evan",
			"Tangerang ",
			"Atlit",
			"Mau",
		},
	}
	args := os.Args

	fmt.Println(getDataById(args[1], orang))
}

func getDataById(id string, orang []Person) string {

	kondisi := false
	indexOrang := 0

	for i, value := range orang {
		if value.Id == id {
			kondisi = true
			indexOrang = i
			break
		}
	}

	if kondisi {
		return fmt.Sprintf("data ditemukan : %+v", orang[indexOrang])
	} else {
		return "Data gada"
	}
}

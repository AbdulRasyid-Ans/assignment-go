package main

import (
	"assignemnt-1/helpers"
	"fmt"
)

func main() {
	dataSiswa := []helpers.Siswa{
		{
			Id:        1,
			Nama:      "Jono Sujono",
			Alamat:    "Jl. Nusa Tenggara 1 No.6 RT.01/RW.01",
			Pekerjaan: "Software Engineer",
			Alasan:    "Karena golang eksekusinya cepat.",
		},
		{
			Id:        2,
			Nama:      "Entis Sutisna",
			Alamat:    "Jl. Rajawali No.2 RT.03/RW.01",
			Pekerjaan: "Backend Engineer",
			Alasan:    "Karena golang mempunyai banyak tools.",
		},
		{
			Id:        3,
			Nama:      "Michael Johnson",
			Alamat:    "Jl. Victoria 3 No.11 RT.02/RW.02",
			Pekerjaan: "Software Engineer",
			Alasan:    "Karena golang mudah dipelajari.",
		},
	}

	idxSiswa := helpers.CariIndexSiswa(dataSiswa, helpers.ArgsInput())

	if idxSiswa >= 0 {
		fmt.Printf("Data Siswa: %+v\n", dataSiswa[idxSiswa])
	} else {
		fmt.Println("Data Siswa tidak ditemukan")
	}

}

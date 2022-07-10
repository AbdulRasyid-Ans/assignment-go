package helpers

import (
	"os"
	"strconv"
)

func ArgsInput() (res int) {
	if len(os.Args) > 1 {
		res, _ = strconv.Atoi(os.Args[1])
	}
	return res
}

func CariIndexSiswa(daftarSiswa []Siswa, id int) (res int) {
	for idx, siswa := range daftarSiswa {
		if siswa.Id == id {
			res = idx
			break
		} else {
			res = -1
		}
	}
	return res
}

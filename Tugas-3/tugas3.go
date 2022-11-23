package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	newPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	newLebar, _ := strconv.Atoi(lebarPersegiPanjang)
	newAlas, _ := strconv.Atoi(alasSegitiga)
	newTinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = newPanjang * newLebar
	var kelilingPersegiPanjang int = 2 * (newPanjang + newLebar)
	var luasSegitiga int = (newAlas * newTinggi) / 2

	fmt.Println("Luas Persegi Panjang = ", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang = ", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga = ", luasSegitiga)
	fmt.Println()

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	// nilai >= 80 indeksnya A
	// nilai >= 70 dan nilai < 80 indeksnya B
	// nilai >= 60 dan nilai < 70 indeksnya c
	// nilai >= 50 dan nilai < 60 indeksnya D
	// nilai < 50 indeksnya E

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John = A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John = B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John = C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John = D")
	} else {
		fmt.Println("Nilai John = E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe = A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe = B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe = C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe = D")
	} else {
		fmt.Println("Nilai Doe = E")
	}

	fmt.Println()

	// soal 3
	var tanggal = 4
	var bulan = 10
	var tahun = 2002
	var newBulan string

	switch bulan {
	case 1:
		newBulan = "Januari"
		break
	case 2:
		newBulan = "Februari"
		break
	case 3:
		newBulan = "Maret"
		break
	case 4:
		newBulan = "April"
		break
	case 5:
		newBulan = "Mei"
		break
	case 6:
		newBulan = "Juni"
		break
	case 7:
		newBulan = "Juli"
		break
	case 8:
		newBulan = "Agustus"
		break
	case 9:
		newBulan = "September"
		break
	case 10:
		newBulan = "Oktober"
		break
	case 11:
		newBulan = "November"
		break
	case 12:
		newBulan = "Desember"
		break
	default:
		newBulan = "undefined"
		break
	}

	fmt.Println(tanggal, newBulan, tahun)
	fmt.Println()

	// soal 4
	// Baby boomer, kelahiran 1944 s.d 1964
	// Generasi X, kelahiran 1965 s.d 1979
	// Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// Generasi Z, kelahiran 1995 s.d 2015

	born := 2002
	if born >= 1944 && born <= 1964 {
		fmt.Println("Baby boomer")
	} else if born >= 1965 && born <= 1979 {
		fmt.Println("Generasi X")
	} else if born >= 1980 && born <= 1994 {
		fmt.Println("Generasi Y")
	} else if born >= 1995 && born <= 2015 {
		fmt.Println("Generasi Z")
	}

}

package main

import (
	"fmt"
)

// soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// soal 2
func introduce(nama string, gender string, pekerjaan string, umur string) (panggilan string) {
	if gender == "laki-laki" {
		panggilan = "Pak "
	} else if gender == "perempuan" {
		panggilan = "Bu "
	}

	panggilan = panggilan + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"

	return
}

func main() {

	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	fmt.Println()

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun
	fmt.Println()

	// soal 3

	// soal 4
	var dataFilm = []map[string]string{}
	// Buatlah closure function disini
	var tambahDataFilm = func(title string, jam string, genre string, tahun string) {
		var film = map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}

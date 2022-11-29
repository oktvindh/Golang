package main

import (
	"fmt"
	"math"
	"strings"
)

func ubahNilai(luasOriginal *float64, kelilingOriginal *float64, radius float64) {
	*luasOriginal = math.Pi * math.Pow(radius, 2)
	*kelilingOriginal = 2 * math.Pi * radius
}

func introduce(sentence *string, nama string, gender string, pekerjaan string, umur string) {
	if strings.ToLower(gender) == "laki-laki" {
		*sentence = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else if strings.ToLower(gender) == "perempuan" {
		*sentence = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
}

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	var radius float64 = 5
	var newRadius *float64 = &radius

	luasLingkaran = math.Pi * math.Pow(radius, 2)
	kelilingLingkaran = 2 * math.Pi * radius

	radius = 10

	fmt.Printf("Luas awal: %f\n", luasLingkaran)
	fmt.Printf("Keliling awal: %f\n", kelilingLingkaran)

	ubahNilai(&luasLingkaran, &kelilingLingkaran, *newRadius)
	fmt.Printf("\nLuas baru: %f\n", luasLingkaran)
	fmt.Printf("Keliling baru: %f\n", kelilingLingkaran)
	fmt.Println()

	// Soal 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Println()

}

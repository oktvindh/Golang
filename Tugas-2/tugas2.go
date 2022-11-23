package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"

	fmt.Println(kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5)
	fmt.Println()

	// soal 2
	halo := "Halo Dunia"
	find := "Dunia"
	ganti := "Golang"
	fmt.Println(strings.Replace(halo, find, ganti, 1))
	fmt.Println()

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var find1 = "s"
	var ganti1 = "S"
	var newKataKedua = strings.Replace(kataKedua, find1, ganti1, 1)

	var find2 = "r"
	var ganti2 = "R"
	var newKataKetiga = strings.Replace(kataKetiga, find2, ganti2, 1)

	var newKataKempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama + " " + newKataKedua + " " + newKataKetiga + " " + newKataKempat)
	fmt.Println()

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	num1, err := strconv.Atoi(angkaPertama)
	num2, err := strconv.Atoi(angkaKedua)
	num3, err := strconv.Atoi(angkaKetiga)
	num4, err := strconv.Atoi(angkaKeempat)

	if err == nil {
		fmt.Println(num1 + num2 + num3 + num4)
	}
	fmt.Println()

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var find3 = "halo"
	var ganti3 = "Hi"
	var newKalimat = strings.Replace(kalimat, find3, ganti3, 2)

	fmt.Println(("\"" + newKalimat + "\" " + "-"), angka)
}

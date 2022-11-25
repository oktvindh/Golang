package main

import "fmt"

func main() {
	// soal 1
	var k int = 20
	for i := 1; i <= k; i++ {
		if i%3 == 0 && i%2 == 1 {
			fmt.Println(i, "- I love Coding")
		} else {
			if i%2 != 0 {
				fmt.Println(i, "- Santai")
			} else {
				fmt.Println(i, "- Berkualitas")
			}
		}
	}
	fmt.Println()

	// soal 2
	for i := 1; i <= 7; i++ {
		var pagar = " "
		for j := 1; j <= i; j++ {
			pagar += "#"
		}
		fmt.Println(pagar)
	}
	fmt.Println()

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var hasil = kalimat[2:]
	fmt.Println(hasil)
	fmt.Println()

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", (i + 1), sayur)
	}
	fmt.Println()

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volume = 1
	for key, value := range satuan {
		volume = volume * value
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Printf("Volume Balok = %d\n", volume)
}

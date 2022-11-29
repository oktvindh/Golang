package main

import "fmt"

var suhu = 10

func convertFarenheit(celcius int) int {
	farenheit := celcius*(9/5) + 32
	return farenheit
}

func cetak(suhu int) {
	fmt.Println("Suhu: ", convertFarenheit(suhu))
}

func main() {
	cetak(suhu)
}

package main

import "fmt"

func reverseArray(arr [7]int) [7]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := [7]int{1, 12, 5, 4, 3, 6, 8}
	fmt.Println(arr)
	fmt.Println(reverseArray(arr))
}

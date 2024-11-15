package main

import "fmt"

func main () {
	var jumlah int
	fmt.Println("Jumlah bilek = ")
	fmt.Scanln(&jumlah)

	for i := 2; i <= jumlah; i++ {
		if jumlah%i == 0 {
			fmt.Print(i, " ")
			}
		}
	}


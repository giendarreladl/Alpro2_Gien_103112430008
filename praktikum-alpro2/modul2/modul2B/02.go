package main

import (
	"fmt"
)

func main() {
	var n int
	var bunga, pita string
	jumlah := 0

	fmt.Print("N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if jumlah > 0 {
			pita += " - "
		}
		pita += bunga
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}

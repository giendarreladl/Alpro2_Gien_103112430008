package main

import "fmt"

func main() {

	var beratKiri, beratKanan, bedaBerat float32
	var isOleng bool

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai")
			break
		}

		if beratKiri+beratKanan > 150 {
			fmt.Println("Proses selesai")
			break
		}

		if beratKiri > beratKanan {
			bedaBerat = beratKiri - beratKanan
		} else {
			bedaBerat = beratKanan - beratKiri
		}

		isOleng = bedaBerat >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", isOleng)
	}
}

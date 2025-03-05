package main

import "fmt"

func main() {

	var a1, a2, a3, a4, a5 int
	var y1, y2, y3 rune

	fmt.Scan(&a1, &a2, &a3, &a4, &a5)
	fmt.Scanln()
	fmt.Scanf("%c%c%c\n", &y1, &y2, &y3)

	y1 += 1
	y2 += 1
	y3 += 1

	fmt.Printf("%c%c%c%c%c\n", a1, a2, a3, a4, a5)
	fmt.Printf("%c%c%c\n", y1, y2, y3)
}

package main

import "fmt"

func main() {
	input := "Kenapa Internet Kok Lambat"

	// jadikan rune
	perHuruf := []rune(input)

	// harusnya gini
	penampung := []rune{} // semua kata

	// penampung per kata
	tampung := []rune{}

	// for dari belakang
	for i := len(perHuruf) - 1; i >= 0; i-- {
		// jika ketemu space baru balik tampung dan masukkan ke penampung
		if string(perHuruf[i]) == " " {
			// balik tampung
			for j := len(tampung) - 1; j >= 0; j-- {
				// masukkan ke penampung
				penampung = append(penampung, tampung[j])
			}
			// tambahkan spasi setelah kata
			penampung = append(penampung, ' ')
			// reset tampung
			tampung = []rune{}
		} else {
			// masukkan ke tampung
			tampung = append(tampung, perHuruf[i])
		}
	}

	// masukkan kata terakhir ke penampung
	for j := len(tampung) - 1; j >= 0; j-- {
		penampung = append(penampung, tampung[j])
	}

	fmt.Println(string(penampung))
}

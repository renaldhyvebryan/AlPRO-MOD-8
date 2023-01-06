package main

import "fmt"

func tebakan(player rune, nilai int) int {
	var jawab int
	i := 0
	fmt.Printf("%c - masukkan angka tebakan ke-%d: ", player, i+1)
	fmt.Scan(&jawab)
	for i < 2 && jawab != nilai {
		i++
		fmt.Printf("%c - masukkan angka tebakan ke-%d: ", player, i+1)
		fmt.Scan(&jawab)
	}
	return jawab
}

func tukerPemenang(benar bool, winner, player *rune) {
	if benar {
		tampung := *winner
		*winner = *player
		*player = tampung
	}
}

func mulaiRonde(ronde int, winner rune, nilai *int) {
	fmt.Printf("Ronde %v:\n", ronde)
	fmt.Printf("%c - masukkan angka rahasia: ", winner)
	fmt.Scan(nilai)
}

func main() {
	var winner, player rune
	var ronde, nilai, answer int
	winner = 'A'
	player = 'B'
	ronde = 1
	mulaiRonde(ronde, winner, &nilai)
	for nilai != -101 {
		answer = tebakan(player, nilai)
		tukerPemenang(nilai == answer, &winner, &player)
		fmt.Printf("%c adalah pemenangnya\n\n", winner)
		ronde++
		mulaiRonde(ronde, winner, &nilai)
	}
	fmt.Println("Permainan Selesai")
}

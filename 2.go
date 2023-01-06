package main

import "fmt"

func main() {
	var n, skor1, skor2, skor3, skor4, skor5, skor6, totalA, totalB int
	fmt.Scan(&n)
	i := 0
	for i < n {
		fmt.Scan(&skor1, &skor2, &skor3, &skor4, &skor5, &skor6)
		totalA += skor1 + skor2 + skor3
		totalB += skor4 + skor5 + skor6
		i++
	}
	if totalA > totalB {
		fmt.Printf("Petinju A: %d dan petinju B: %d\nPemenang adalah petinju A", totalA, totalB)
	} else if totalB > totalA {
		fmt.Printf("Petinju A: %d dan petinju B: %d\nPemenang adalah petinju B", totalA, totalB)
	} else {
		fmt.Printf("Draw")
	}

}

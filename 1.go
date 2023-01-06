package main

import "fmt"

func main() {
	var dadu1, dadu2, hasil int
	var cek bool
	fmt.Scan(&dadu1, &dadu2)
	cek = dadu1%2 == 0 && dadu2%2 == 0
	for !cek {
		if dadu1%2 == 1 && dadu2%2 == 1 {
			hasil++
		}
		fmt.Scan(&dadu1, &dadu2)
		cek = dadu1%2 == 0 && dadu2%2 == 0
	}
	fmt.Println(hasil)
}

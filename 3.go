package main

import "fmt"

func main() {
	var n1, n2, kpk int
	fmt.Scan(&n1, &n2)
	kpk = n1
	for kpk%n2 != 0 {
		kpk += n1
	}
	fmt.Println(kpk)
}

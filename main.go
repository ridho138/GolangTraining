package main

import "fmt"

func main() {

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d Genap\n", i)
	// 	} else {
	// 		fmt.Printf("%d Ganjil\n", i)
	// 	}
	// }

	var names = []string{"Kevin Hugo", "Kadek Bintang Anjasmara", "Guntur Satrya Saputro", "Achmad Fathoni", "Edwin Setya Noegroho", "Jaka Prima Maulana", "Stevanus Dewana", "Hans Parson", "Rizki Ramadhan", "Mochammad Zayyan Ramadhan"}

	for i, s := range names {
		fmt.Println(i+1, s)
	}
}

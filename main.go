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

	name()
}

func name() {
	var names = []string{"Kevin Hugo", "Kadek Bintang Anjasmara", "Guntur Satrya Saputro", "Achmad Fathoni", "Edwin Setya Noegroho", "Jaka Prima Maulana", "Stevanus Dewana", "Hans Parson", "Rizki Ramadhan", "Mochammad Zayyan Ramadhan"}

	var namesPointer []*string

	for i := 0; i < len(names); i++ {
		namesPointer = append(namesPointer, &names[i])
	}

	var cetakNama = func(l []*string) {
		for _, s := range l {
			fmt.Println(*s)
		}
	}
	cetakNama(namesPointer)
}

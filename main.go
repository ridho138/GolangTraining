package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d Genap\n", i)
		} else {
			fmt.Printf("%d Ganjil\n", i)
		}
	}
}

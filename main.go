package main

import (
	"fmt"
	"golangTraining/service"
)

func main() {

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d Genap\n", i)
	// 	} else {
	// 		fmt.Printf("%d Ganjil\n", i)
	// 	}
	// }

	//name()

	userSvc := service.NewUserService()
	res := userSvc.Register(&service.User{Nama: "budi"})
	fmt.Println(res)
	res2 := userSvc.Register(&service.User{Nama: "Anto"})
	fmt.Println(res2)

	resHasil := userSvc.GetUser()
	fmt.Println("----------------Hasil Get User--------------")
	for _, v := range resHasil {
		fmt.Println(v.Nama)
	}
}

func name() {

	type Persons struct {
		name string
	}

	var names = []string{"Kevin Hugo", "Kadek Bintang Anjasmara", "Guntur Satrya Saputro", "Achmad Fathoni", "Edwin Setya Noegroho", "Jaka Prima Maulana", "Stevanus Dewana", "Hans Parson", "Rizki Ramadhan", "Mochammad Zayyan Ramadhan"}

	var namesPointer []*Persons

	for i := 0; i < len(names); i++ {
		var orang Persons
		orang.name = names[i]
		namesPointer = append(namesPointer, &orang)
	}

	var cetakNama = func(l []*Persons) {
		for _, s := range l {
			fmt.Println(s.name)
		}
	}
	cetakNama(namesPointer)
}

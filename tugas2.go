package main

import "fmt"

func main(){
	angka := 12

	if angka % 2 != 0 {
		fmt.Println(angka, "adalah bilangan ganjil")
	} else {
		fmt.Println(angka, "adalah bilangan genap")
	}
}
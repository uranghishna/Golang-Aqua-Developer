package main

import "fmt"

// functions
func swap(x,y string)(string, string){
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a,b)

	fmt.Println("hello world")
	
	// Numeric Decimal
	var decimaln = 2.623463
	fmt.Printf("Bilangan decimal : %f\n", decimaln)
	fmt.Printf("Bilangan decimal : %.3f\n", decimaln)

	// Boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// String
	var message = 
	`Nama saya : Koje.
	salam kenal.
	saya belajar golang.`
	fmt.Println(message)

	nama:= "hishna"
	fmt.Println("nama saya",nama)

	// Variables Declaration
	var first string = "hishna"
	var last string
	last = "difa"
	fmt.Println("halo %s %s!\n", first, last)

	// Declaration Variables without Type Data
	var name1 string = "John"
	name2 := "Wick"
	fmt.Println("halo %s %s!\n", name1, name2)

	// Declaration Multi Variable
	// var firstt, second, third string
	// firstt, second, third = "satu", "dua", "tiga"
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	
	// // Declaration Underscore Variables
	// _="belajar golang"
	// _="golang itu mudah"
	// nama, _ := "ja","wi"

	// Constants
	const firstName string = "koje"
	fmt.Println("halo",firstName,"!\n")

	// if..else..then
	var point = 8
	if point == 10{
		fmt.Println("lulus sempurna")
	}else if point > 5{
		fmt.Println("lulus")
	}else if point == 4{
		fmt.Println("hampir lulus")
	}else{
		fmt.Println("tidak lulus, nilai anda %d\n", point)
	}

	// switch..case
	// var point = 6
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")	
	default:
		fmt.Println("not bad")
	}

	// for..range
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("element %d : %s", i, fruit)		
	}

	// for..loop
	for i:=0; i<5; i++ {
		fmt.Println("Angka", i)
	}

	// for..loop break continue
	for i:=1; i<=10; i++ {
		if i%2 == 1 {
			continue
		}
		if i>8{
			break
		}
		fmt.Println("Angka", i)
	}
	
	

}



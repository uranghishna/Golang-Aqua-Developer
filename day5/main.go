package main

import "fmt"


type Product struct{
	id int
	name string
	price int
}

var allProd = []Product{
	{id: 1, name : "Benih lele", price:50000},
	{id: 2, name : "Pakan lele cap menara", price:25000},
	{id: 3, name : "Probiotik A", price:75000},
	{id: 4, name : "Probiotik nila B", price:10000},
	{id: 5, name : "Pakan nila", price:20000},
	{id: 6, name : "Benih nila", price:20000},
	{id: 7, name : "Cupang", price:5000},
	{id: 8, name : "Benih nila", price:30000},
	{id: 9, name : "Benih cupang", price:10000},
	{id: 10, name : "Probiotik B", price:10000},
}

func main() {
	// NOMOR 1 A
	fmt.Println("Nomor 1 A")
	for _, prod := range allProd {
		fmt.Println("Barang :", prod.name)
		fmt.Println("Harga :", prod.price)
	}
	
	var point int
	fmt.Println("Masukan jumlah point anda : ")	
	fmt.Scan(&point)

	sortedProd := SortProductByPrice(allProd)
	
	var items []Product
	 
	if point > 0 {
		for _,prod := range sortedProd {
			if point >=  prod.price && !CheckArray(prod.id,items)  {
				items = append(items,prod)
				point = point - prod.price
			}
		}
	}
	fmt.Printf("Items yang anda pilih : ")
	for i, item := range items {
		if i != len(items)-1{
			fmt.Printf("%s, ", item.name)
		} else {
			fmt.Printf("%s.", item.name)
		}
	}
	fmt.Println("")
	fmt.Println("Sisa point anda :", point)


	// NOMOR 1 B
	fmt.Println("Nomor 1 B")
	fmt.Println(SortProductByPriceMin(allProd))
	fmt.Println(SortProductByPriceMax(allProd))


	// NOMOR 1 C
	fmt.Println("Nomor 1 C")
	fmt.Println(SortByPrice(allProd))
}

// NOMOR 1 A
func SortProductByPrice(Products []Product) []Product {
	for i := 0; i < len(Products); i++ {
		for j := 0; j < len(Products); j++ {
			if Products[i].price < Products[j].price {
				temp := Products[i]
				Products[i] = Products[j]
				Products[j] = temp
			}
		}
	}
	return Products
}

func CheckArray(Id int, Products []Product) bool{
	if len(Products) > 0{
		for _, prod := range Products {
			if prod.id == Id{
				return true
			}
		}
	}
	return false
}


// NOMOR 1 B
func SortProductByPriceMin(Products []Product) Product {
	for i := 0; i < len(Products); i++ {
		for j := 0; j < len(Products); j++ {
			if Products[i].price < Products[j].price {
				temp := Products[i]
				Products[i] = Products[j]
				Products[j] = temp
			}
		}
	}
	return Products[0]
}

func SortProductByPriceMax(Products []Product) Product {
	for i := 0; i < len(Products); i++ {
		for j := 0; j < len(Products); j++ {
			if Products[i].price > Products[j].price {
				temp := Products[i]
				Products[i] = Products[j]
				Products[j] = temp
			}
		}
	}
	return Products[0]
}


// NOMOR 1 C
func SortByPrice(Products []Product) []Product {
	var Ceban []Product
	for _,prod := range Products {	
		if prod.price == 10000 {
			Ceban=append(Ceban,prod)
		}
	}
	return Ceban
}
package main

import "fmt"

//:Bir sayının tek mi çift mi olduğunu ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız
func main() {

	fmt.Print("Enter number : ")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "is Even number")
	} else {
		fmt.Println(n, "is Odd number")
	}
}

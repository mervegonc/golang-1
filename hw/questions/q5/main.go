package main

import "fmt"

func main() {
	var smsize, i, position int

	fmt.Print("Enter the Array Size to find the Smallest = ")
	fmt.Scan(&smsize)

	smArr := make([]int, smsize)

	fmt.Print("Enter the Smallest Array Items  = ")
	for i = 0; i < smsize; i++ {
		fmt.Scan(&smArr[i])
	}
	smallest := smArr[0]

	for i = 0; i < smsize; i++ {
		if smallest > smArr[i] {
			smallest = smArr[i]
			position = i
		}
	}
	fmt.Println("\nThe Smallest Number in this smArr    = ", smallest)
	fmt.Println("The Index Position of Smallest Number = ", position)
}

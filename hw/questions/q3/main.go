package main

import "fmt"

func sumOfNaturalNumbers(n int) int {
	var sum int
	sum = (n * (n + 1) / 2)
	return sum

}
func main() {
	var n, answer int
	n = 100
	fmt.Println("Program to find the sum of the Natural number using the formula.")
	answer = sumOfNaturalNumbers(n)
	fmt.Println("The sum of", n, "natural numbers is", answer)
}

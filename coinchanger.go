package main

import "fmt"

func main() {
	n := 4
	fmt.Println("Change for", n, "[1 1 1 1]:", coinChanger(n))

	n = 99
	fmt.Println("Change for", n, "[25 25 25 10 10 1 1 1 1]:", coinChanger(n))

	n = 17
	fmt.Println("Change for", n, "[10 5 1 1]:", coinChanger(n))

	n = 6
	fmt.Println("Change for", n, "[5 1]:", coinChanger(n))

	n = 33
	fmt.Println("Change for", n, "[25 5 1 1 1]:", coinChanger(n))
}

func coinChanger(amount int) []int{
	values := []int{25, 10, 5 , 1}
	coins := make([]int, 0)

	for _,element := range values {
		remaining, sumValue := change(amount, element)
		amount = remaining

		for i := 1; i <= sumValue; i++ {
			coins = append(coins, element)
		}
	}

	return coins
}

func change(amount, value int) (remaining, sumValue int) {
	sumValue = amount / value // gets total coins of a value used
	remaining = amount % value //gets new total after total coins used
	return remaining, sumValue
}
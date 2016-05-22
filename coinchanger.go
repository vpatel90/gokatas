package main

import "fmt"

func main() {
	n := 4
	fmt.Println("Coins [1 1 1 1]:", coinChanger(n))
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
	return remaining, amount / value
}
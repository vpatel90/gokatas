package main

import "fmt"

func main() {
	n := 4
	fmt.Println("Coins :", coinChanger(n))
}

func coinChanger(amount int) []int{
	coins := make([]int, 0)
	fmt.Println(amount)
	for i := 1; i <= amount; i++ {
		coins = append(coins, 1)

	}
	return coins
}
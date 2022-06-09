package main

import (
	"fmt"
	"math/rand"
)

var PlayerDice [][]int

func GetRandomNumber() int {
	number := rand.Intn(7)
	if number == 0 {
		number = 3
	}
	return number
}

func ThrowDice(Player, M int) []int {
	var Dice []int
	for i := 0; i < M; i++ {
		Dice = append(Dice, GetRandomNumber())
	}
	return Dice
}

func PlayerThrowDice(N, M int) [][]int {
	for i := 0; i < N; i++ {
		var playerDice [][]int
		for i := 0; i < N; i++ {
			playerDice = append(playerDice, ThrowDice(i, le))
		}
	}
	return playerDice
}

func main() {
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
	fmt.Println(PlayerThrowDice(2, 4))
}

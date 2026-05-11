package main

import "fmt"

type functionType func(int) int

func main() {
	array := []int{10, 20, 30, 40, 50}
	doubbleArray := transFormNumber(array, doubble)
	trippleArray := transFormNumber(array, triple)
	fmt.Println(doubbleArray)
	fmt.Println(trippleArray)
}

func transFormNumber(array []int, transform functionType) []int {
	dArray := []int{}
	for _, val := range array {
		dArray = append(dArray, transform(val))
	}
	return dArray
}

func doubble(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
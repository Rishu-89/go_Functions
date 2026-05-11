// package anonymousfunction
package main

import "fmt"

type functionType func(int) int

func main() {
	array := []int{10, 20, 30, 40, 50}
	doubbleArray := transFormNumber(&array, func(number int)int{
		return number*2
	} )
	fmt.Println(doubbleArray)
}

func transFormNumber(array *[]int, transform functionType) []int {
	dArray := []int{}
	for _, val := range *array {
		dArray = append(dArray, transform(val))
	}
	return dArray
}


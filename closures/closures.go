// package anonymousfunction
package main

import "fmt"

type functionType func(int) int

func main() {
	array := []int{10, 20, 30, 40, 50}
	doubbleFunction:=factoryFunction(2)
	trippleFunction:=factoryFunction(3)

	doubbleArray:=transFormNumber(&array,doubbleFunction)
	trippleArray:=transFormNumber(&array,trippleFunction)

	fmt.Println(doubbleArray)
	fmt.Println(trippleArray)
}

func transFormNumber(array *[]int, transform functionType) []int {
	dArray := []int{}
	for _, val := range *array {
		dArray = append(dArray, transform(val))
	}
	return dArray
}

func factoryFunction(factor int)func(int)int{
	return func(number int)int{
		return number*factor
	}
}
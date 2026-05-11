package main

import "fmt"

func main() {

	sum:=sumUp(10,20,30,40)
	fmt.Println(sum)
	array:=[]int{10,20,30,40}
	anotherSum:=sumUp(array...)
	fmt.Println(anotherSum)
}

func sumUp(numbers ...int) int {
	temp := 0
	for _, val := range numbers {
		temp += val
	}
	return temp
}
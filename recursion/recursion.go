package main

import "fmt"

func main() {
	factorial :=factorialFunction(5)
		fmt.Println(factorial)
}

func factorialFunction(val int)int{
	if val==0{
		return 1
	}
	return val*factorialFunction(val-1)
}
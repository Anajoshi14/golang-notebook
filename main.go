package main

import (
	"fmt"

	"github.com/Anajoshi14/golang_repo/calculator"
	"github.com/Anajoshi14/golang_repo/phone"
)

func main() {
	array := []int{5, 15, 20}
	c := calculator.CreateCalculator()
	// r := c.MultiplyNumbers(array)
	//fmt.Println(r)
	//PrintPrathmesh()

	ph := phone.CreatePhone(c)
	r := ph.Calci.MultiplyNumbers(array)
	fmt.Println(r)

}

func PrintPrathmesh() {
	i := 0
	for i < 5 {
		fmt.Println("Prathmesh")
		i++
	}

}

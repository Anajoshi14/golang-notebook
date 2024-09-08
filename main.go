package main

import (
	"fmt"

	pntr "github.com/Anajoshi14/golang_repo/Pointers"
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
	a := 10
	pntr.ChangeInt(&a)
	fmt.Println(a)
	// whenever calling a function from diff package, write package name.function name and(&variable name)
	//&variable name : whenever we want to pass pointer to a variable and not a copy of variable then we should do this

}

func PrintPrathmesh() {
	i := 0
	for i < 5 {
		fmt.Println("Prathmesh")
		i++
	}

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set the handler function for the root URL
	http.HandleFunc("/", handler)

	// Define the port on which the server will listen
	port := "8080"

	fmt.Printf("Starting server on port %s...\n", port)

	// Start the server and listen on the specified port
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// func main() {
// 	array := []int{5, 15, 20}
// 	c := calculator.CreateCalculator()
// 	// r := c.MultiplyNumbers(array)
// 	//fmt.Println(r)
// 	//PrintPrathmesh()

// 	ph := phone.CreatePhone(c)
// 	r := ph.Calci.MultiplyNumbers(array)
// 	fmt.Println(r)
// 	a := 10
// 	pntr.ChangeInt(&a)
// 	fmt.Println(a)
// 	// whenever calling a function from diff package, write package name.function name and(&variable name)
// 	//&variable name : whenever we want to pass pointer to a variable and not a copy of variable then we should do this

// }

func PrintPrathmesh() {
	i := 0
	for i < 5 {
		fmt.Println("Prathmesh")
		i++
	}

}

// handler function for the root URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love you Prathmesh, Muhhh")
}

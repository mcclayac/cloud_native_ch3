package main

import "fmt"

func main() {

	fmt.Println("Cloud Native Chapter 3")
	{
		fmt.Println("\nQuick booleans")
		and := true && false
		fmt.Println(and)

		or := true || false
		fmt.Println(or)

		not := !true
		fmt.Println(not)
	}
}

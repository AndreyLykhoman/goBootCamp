package main

import "fmt"

func second() {
	fmt.Println("second.begin")



	panic("Something went wrong((")

	fmt.Println("second.end")
}

func first() {
	fmt.Println("first.begin")

	second()

	fmt.Println("first.end")
}

func main() {
	fmt.Println("main.begin")

	defer func() {
		fmt.Println("main.defer")

		if error := recover(); error != nil {
			fmt.Println("main.recover  :", error)
		}
	}()

	first()

	fmt.Println("main.end")
}

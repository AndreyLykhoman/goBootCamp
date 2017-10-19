	package main

	import "fmt"

	func main() {
		id := make(chan string)

		go func() {
			var counter int64 = 1
			for {
				id <- fmt.Sprintf("%x", counter)
				counter += 1
			}
		}()

		fmt.Printf("%s\n", <-id) // will be 1
		fmt.Printf("%s\n", <-id) // will be 2
	}

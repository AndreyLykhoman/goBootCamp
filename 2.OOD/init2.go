package main 

import "fmt"

var stuff = "not ready"

func init() {
  stuff = "ready"
  fmt.Println("init1")
}

func init() {
  stuff = "ready2"
  fmt.Println("init2")
}

func main() {
  fmt.Println("The stuff is", stuff)
}

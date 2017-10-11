package main

import "fmt"

func sum(xs []int64) int64 {    
    //total := 0
    //total := int64(0)
    var total int64
    for _, v := range xs {
        total += v
    }
    return total
}
 
func main() {
    xs := []int64{32,12,1,232,37}
    fmt.Println(sum(xs))
}

func sum(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}


func main() {
    fmt.Println(sum(1,2,3))
}

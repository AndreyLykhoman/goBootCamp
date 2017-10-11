func newInt1() *int { return new(int) }

func newInt2() *int {
    var i int
    return &i
}

package main
import (
	"fmt"
)
type MyInt = int
type MyInt2 = MyInt
func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}
func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}
func main() {
	i := MyInt(0)
	i.add(3)
	i.showYourSelf()
}

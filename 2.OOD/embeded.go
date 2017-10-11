package main
import "fmt"
type User struct {
	Name string
	Age  int
}
func (u *User) Greeting() {
	fmt.Println("Hello I am user")
}
type Admin struct {
	User
	Level int
}
func (a *Admin) Greeting() {
	fmt.Println("Hello I am superuser")
}
func main() {
	u := &User{}
	a := Admin{}
	u.Greeting()
	a.Greeting()
	a.User.Greeting()
}

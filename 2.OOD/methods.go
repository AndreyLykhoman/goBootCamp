func (c Circle) Render() {
  fmt.Println("Circle with X:", c.x, ", and Y is ", c.y, ", and R is ", c.r)
}
 
func (c *Circle) Rotate() {
  c.x, c.y = c.y, c.x
}

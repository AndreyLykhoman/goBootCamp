func имяФункции([список параметров])  [Возвращаемые значения] {
    тело функции
} 

func sum(xs []int64) int64 {    
    total := 0
    for _, v := range xs {
        total += v
    }
    return total
}
 
func main() {
    xs := []int64{32,12,1,232,37}
    fmt.Println(sum(xs))
}

// total := нулевое значение типа int64
func sumNew(xs []int64) (total int64) {    
    for _, v := range xs {
        total += v
    }
    return total
}

type Circle struct {
    x float64
    y float64
    r float64
}

func (s Shape) Render() {
  fmt.Println(“Shape width is ”, s.width,”, and height is ”, s.height)
}
 
func (s *Shape) Rotate() {
  s.width, s.height = s.height, s.width
}

func multireturn(s string) (string, int) {
	return s, len(s)
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}

type Speaker interface {
    SayHello()
}

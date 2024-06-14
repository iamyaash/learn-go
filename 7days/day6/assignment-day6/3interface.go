package main
import "fmt"

type Shape interface {
	Perimeter() float64
}

type Square struct {
	side float64
}

func main() {
	sqr := Square{side:4}
	perSqr := sqr.Perimeter()
	fmt.Println(perSqr)
}

func (s Square) Perimeter() float64{
	return s.side * 4
}

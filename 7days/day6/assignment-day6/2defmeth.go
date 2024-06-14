package main 
import "fmt"

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func main() {
	rect := Rectangle{height : 10, width: 5}
	area := rect.Area()
	fmt.Println(area)
}

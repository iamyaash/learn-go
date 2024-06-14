package main
import "fmt"

func main() {
	var a int = 10
	var p *int = &a

	fmt.Println(p)

	*p = 20
	fmt.Println(a)
}

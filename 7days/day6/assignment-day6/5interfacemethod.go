package main

import "fmt"

type Drawable interface {
	Draw() string
}

type Triangle struct {
	base float64
	height float64
}

func main() {
	tri := Triangle{base:8, height:17}
	sol := tri.Draw()
	fmt.Println(sol)
}

func (t *Triangle) Draw() string{
    return fmt.Sprintf("Triangle with base %.2f and height %.2f", t.base, t.height)
}

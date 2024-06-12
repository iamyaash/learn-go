package main
import "fmt"

func main() {
    multier := func() func(int) int {
        res := 1
        return func(x int) int {
            res *= x
            return res
        }
    }
    val := multier()

    for i := 1; i <= 10; i++ {
        fmt.Println(val(i))
    }
}
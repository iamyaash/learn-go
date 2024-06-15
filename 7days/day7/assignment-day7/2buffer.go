package main
import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Buffered messages 1"
	messages <- "Buffered messages 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

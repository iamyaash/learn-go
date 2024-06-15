package main
import "fmt"

func main() {
	messages := make(chan string)
	
	go func() {
		messages <- "Hello from 1"
	}()

	go func() {
		messages <- "Hello from 2"
	}()
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

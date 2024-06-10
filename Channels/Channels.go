package Channels

import "fmt"

func main() {

	messages := make(chan string) // pipes for goroutines

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

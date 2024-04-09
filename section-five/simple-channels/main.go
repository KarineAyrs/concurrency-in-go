package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		//get user input

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		// wait for a response

		response := <-pong

		fmt.Println("response:", response)
	}

	fmt.Println("all done. closing channels")

	close(ping)
	close(pong)
}

package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "server 2"
	}
}

func main() {
	fmt.Println("select with channels")
	fmt.Println("--------------------")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	for {
		// chooses identical cases randomly
		select {
		case s1 := <-ch1:
			fmt.Println("Case 1: ", s1)
		case s2 := <-ch1:
			fmt.Println("Case 2: ", s2)
		case s3 := <-ch2:
			fmt.Println("Case 3: ", s3)
		case s4 := <-ch2:
			fmt.Println("Case 4: ", s4)
			//default:
			// avoiding deadlock
		}
	}
}

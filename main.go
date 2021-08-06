package main

import (
	"fmt"
	"time"
)

/*
//Basic Go Routine Example
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println(s)
	}
}
*/

//Channel Example
func say(s chan<- string) {
	for i := 0; i < 3; i++ {
		s <- "Fresco"
	}
}

func printer(s <-chan string) {
	for {
		msg := <-s
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func play(c chan string) {
	for i := 0; i < 3; i++ {
		c <- "Play"
	}
}

func main() {
	//Channel Example
	var s chan string = make(chan string)
	go say(s)
	go play(s)
	go printer(s)
	var input string
	fmt.Scanln(&input)
	/*
		//Basic Go Routine Example
		go say("Fresco")
		say("Play")
		//var input string
		//fmt.Scanln(&input)
	*/
}

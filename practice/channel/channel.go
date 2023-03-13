package main

import "fmt"

func doSomething(ch chan int) {
	x := 42
	ch <- x // send x on the channel
}

func main() {
	ch := make(chan int) // create an unbuffered channel

	go doSomething(ch) // launch a goroutine to do something with x

	x := <-ch // receive x from the channel
	fmt.Println(x)
}

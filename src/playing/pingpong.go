package main

import (
	"fmt"
	"time"
)

//the pingers print a ping and wait for a pong
func pinger(pinger chan int, ponger chan int) {

	for {

		<-pinger
		fmt.Println("Ping")
		time.Sleep(time.Second * 1)
		ponger <- 1

	}
}

//the ponger print a pong and wait for a ping
func ponger(pinger chan int, ponger chan int) {

	for {
		<-ponger
		fmt.Println("Pong")
		time.Sleep(time.Second * 1)
		pinger <- 1

	}
}

func main() {

	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	fmt.Println("hola")

	//the main goroutine starts the ping/pong by sending int the ping channel.
	ping <- 1

	for {
		//block the main thread until an interrupt
		time.Sleep(time.Second)
	}

}

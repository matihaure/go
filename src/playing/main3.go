package main

import (
	"fmt"
	"time"
)

func func1() {

	for {
		time.Sleep(time.Second * 5)
		println("FUNC1 EXEC")
	}
}
func func2() {

	time.Sleep(time.Second * 2)
	println("FUNC2 EXEC")
}

func main() {

	fmt.Println("HOLA MUNDO")
	//time.Sleep(time.Second * 5)
	//go func1()
	//go func2()

	//Channel de enteros SIN BUffer

	//	c := make(chan int)

	//Chanel de enteros con Buffer

	c2 := make(chan int, 5)

	//Funcion lambda - solo se ejecuta aca
	//Le agrego valores al channel
	/* 	go func() {

		c <- 1

	}() */

	//Funcion lambda - solo se ejecuta aca
	//Le agrego valores al channel
	go func() {

		for i := 1; i < 21; i++ {
			c2 <- i
		}
		close(c2)

	}()

	//Snifeo el channel en runtime.

	for i := range c2 {
		fmt.Println(i)
	}

}

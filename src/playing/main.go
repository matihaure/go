package main

import (
	"fmt"
	"strings"
)

type Phone struct {
	Model string
	OS    string
	RAM   int
}

func main() {

	//var m1 = "hola"
	//m2 := "ho"

	//fmt.Println(strings.Contains(m1, m2))

	//stringcont(m1, m2)
	createarray(1, 2)

	//p := Phone{}

	//fmt.Println(p.createPhone())

}

func stringcont(a string, b string) {

	fmt.Println(strings.Contains(a, b))

}

func createarray(a int, b int) {

	var arr1 = []int{a, b}
	var arr2 = []string{"hola", "perro", "manuel"}
	arr2 = append(arr2, "perrito", "caballito")
	arr1 = append(arr1, 2, 2, 3)

	fmt.Println(arr1)
	fmt.Println(arr2)

}

package main

import (
	"fmt"
)

func main() {

	//fmt.Println("hello world")
	/*
		a := true

		for i := 0; a == true && i < 10; i++ {
			conditional(a, i)
			fmt.Print("Iteracion N° ")
			fmt.Println(i)
			if i == 5 {
				a = false
			}
		}
	*/

	//array := []string{"hola", "matias", "caballo", "perro"}

	mymap := make(map[string]interface{})

	mymap["Nombre"] = "Matias"
	mymap["Apellido"] = "Haure"
	mymap["Edad"] = 30

	//fmt.Println(array)
	//printarray(array)

	printmap(mymap)

}

func conditional(a bool, i int) {

	mem := &a

	if mem == nil {
		fmt.Println("valor nulo")
	} else if *mem {
		fmt.Println("TRUE")

	} else {
		fmt.Println("FALSE")
	}

	fmt.Println("el valor en memoria de a es: ")
	fmt.Println(mem)
	fmt.Println("el valor de a es: ")
	fmt.Println(a)
}

func printarray(array []string) {

	for i, value := range array {
		fmt.Println(i)
		fmt.Println(value)
	}

}

func printmap(m map[string]interface{}) {

	for k, v := range m {

		fmt.Print(k + ": ")
		fmt.Println(v)

	}

}

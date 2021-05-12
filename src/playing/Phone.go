package main

import "fmt"

type Phone interface {
	TurnOn()
	Charge()
}

type Samnsung struct {
	Model string
	OS    string
}

type IPhone struct {
	Model string
	OS    string
}

func (s *Samnsung) Charge() {
	fmt.Println(s.Model + " is charging")
}

func (i *IPhone) TurnOn() {
	fmt.Println(i.Model + " is turning on")
}

func main() {

	s := Samnsung{Model: "S7",
		OS: "Android",
	}
	i := IPhone{Model: "12",
		OS: "ios",
	}

	s.Charge()
	i.TurnOn()

}

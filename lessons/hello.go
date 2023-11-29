package main

import "fmt"

func main() {
	fmt.Println("This line is printed")
	variables()
	Arrays()
}

func variables() {
	name := "John Doe"
	intgr := 24
	floate := 154.61
	boole := true

	fmt.Println(name)
	fmt.Println(floate)
	fmt.Println(intgr)
	fmt.Println(boole)
}

func Arrays() {
	arrstring := [3]string{"hello", " ", "world"}
	arrint := [3]int{1, 2, 3}

	fmt.Println(arrstring[2])
	fmt.Println(arrint[2])
}

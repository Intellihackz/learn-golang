package main

import (
	"fmt"
	"lessons/setup/leetcodes"
)

func main() {
	fmt.Println("This line is printed")
	variables()
	Arrays()
	slices()
	ifstate()
	loops()
	leetcodes.Leetcode()
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

func slices() {
	slicestring := []string{"hello", "world", "!"}
	fmt.Println(slicestring)
}

func ifstate() {
	age := 18
	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")

	}
}

func loops() {
	sum := 0
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		sum = sum + i
	}
	fmt.Println(sum)
}

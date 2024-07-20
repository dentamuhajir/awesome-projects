package main

import "fmt"

func main() {
	fmt.Print("Hello, world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 1
	fmt.Println(i)

	whatWasSaid := saysSomething()
	fmt.Print(whatWasSaid)
	whatWasSaid, theOtherWasSaid := saysSomethingMore()
	fmt.Print(whatWasSaid, theOtherWasSaid)
}

func saysSomething() string {
	return "something"
}

func saysSomethingMore() (string, string) {
	return "say something", " more"
}

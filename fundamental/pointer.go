package fundamental

import "log"

func Pointer() {
	var myString string
	myString = "green"
	log.Printf(myString)
	changeUsingPointer(&myString)
	log.Printf("after change using pointer", myString)
}

func changeUsingPointer(s *string) {
	//
	log.Printf("s is set to ", s)
	newValue := "red"
	*s = newValue
}

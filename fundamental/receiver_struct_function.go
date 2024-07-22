package fundamental

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName + " from receiver struct"
}

func ReceiverStructFunction() {
	var myVar = myStruct{}
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Jane",
	}

	log.Println("myVar1:", myVar.printFirstName())
	log.Println("myVar2:", myVar2)
}

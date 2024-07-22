package fundamental

import "log"

type Animal struct {
	Name           string
	Classification string
}

func MapAndSlices() {
	myMap := make(map[string]string)
	myMap["dog"] = "husky"
	myMap["cat"] = "meow"
	myMap["dog"] = "samson"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	myMapInt := make(map[string]int)
	myMapInt["dog"] = 1
	myMapInt["cat"] = 2

	log.Println(myMapInt["cat"])
	log.Println(myMapInt["dog"])

	myMapStruct := make(map[string]Animal)
	myMapStruct["dog"] = Animal{
		Name:           "Husky Struct",
		Classification: "dog",
	}
	log.Println(myMapStruct["dog"].Name)
}

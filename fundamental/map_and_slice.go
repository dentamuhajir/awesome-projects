package fundamental

import (
	"log"
	"sort"
)

type Animal struct {
	Name           string
	Classification string
}

func Maps() {
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

func Slicing() {
	var mySlice []string

	mySlice = append(mySlice, "dog")
	mySlice = append(mySlice, "cat")
	mySlice = append(mySlice, "bird")

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(numbers)

	log.Println(numbers[0:3])

}

package fundamental

import (
	"log"
)

func LoopAndRangingOverData() {
	//loopData()
	//rangingArray()
	rangingStruct()
}

func loopData() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func rangingArray() {
	animals := []string{"Dog", "Cat", "Bird", "horse", "fish"}
	food := make(map[string]string)
	food["asian"] = "soup"
	food["western"] = "spaghetti"
	for i, animal := range animals {
		log.Println(i, animal)
	}
	for key, value := range food {
		log.Println(key, value)
	}
}

func rangingStruct() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Denta", "Muhajir", "dentamuhajir@gmail.com", 24})
	users = append(users, User{"John", "Mayer", "johnmayer@gmail.com", 37})
	users = append(users, User{"Tom", "Anderson", "tomanderson@gmail.com", 44})

	for _, i := range users {
		log.Println(i.FirstName, i.LastName, i.Age, i.Email)
	}

}

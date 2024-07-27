package main

import (
	"fmt"
	"net/http"
)

func main() {

	//fundamental.VarAndFunction()
	//fundamental.Pointer()
	//fundamental.ReceiverStructFunction()
	//fundamental.Maps()
	//fundamental.Slicing()
	//fundamental.DecisionStructure()
	//fundamental.LoopAndRangingOverData()
	//fundamental.LoopAndRangingOverData()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)

}

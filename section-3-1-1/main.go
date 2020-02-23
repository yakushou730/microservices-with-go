package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	ex1, err := json.Marshal([]int{1, 2, 3})
	if err != nil {
		log.Println("Error Marshalling: " + err.Error())
	}
	fmt.Println(ex1)
	fmt.Println(string(ex1))

	var ex1Dat []int
	err = json.Unmarshal(ex1, &ex1Dat)
	if err != nil {
		log.Println("Error Unmarshalling: " + err.Error())
	}
	fmt.Println(ex1Dat[1])
}

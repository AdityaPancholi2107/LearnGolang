package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:  "title"`
	Author string `json:  "author"`
}

func main() {

	book := Book{Title: "Learning Go", Author: "Aditya"}
	fmt.Printf("%+v\n", book)

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}

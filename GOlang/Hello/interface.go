package main

import (
	"fmt"

	"golang.org/x/text/message"
)

// defining an interface
type Writer interface{
	Write(message string)
}
// implementing an interface
type ConsolWriter struct{}

// implement the write method 
func (cw ConsolWriter) Write(message string){
	fmt.Println("Writing:", message)

}

func main(){
// creating a variable of type Writer interface
	var writer Writer

	writer = ConsoleWriter{}

	writer.Write("Hello, Go!!!")



}


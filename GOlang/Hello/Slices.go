package main

import "fmt"

type Part string

func showline (line []Part){
	for i:= 0; i < len(line); i++{
		part := line[i]
		fmt.Println(part)
	}


}
func main(){

	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"

	assemblyLine = append(assemblyLine, "washer", "Wheel")
	fmt.Println("\n Added two parts")
	showline(assemblyLine)

	assemblyLine= assemblyLine[3:]
	fmt.Println("\nSliced:")
}
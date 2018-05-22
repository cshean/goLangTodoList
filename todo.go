package main

import (
	"fmt"
	"bufio"
	"os"
)

func addNewTasks(reader *bufio.Reader){
	var anotherTask bool = true
	var text string = ""
	var response string = ""
	for anotherTask == true {

		fmt.Println("Enter a task:")
		text, _ = reader.ReadString('\n')
		//Add the task to something
		fmt.Println(text)
		fmt.Println("Do you have another task to enter? (Y/N)")
		response, _  = reader.ReadString('\n')
		if response != "Y\r\n" {
			anotherTask = false
		//probably should add in another statement here to account for incorrect Y/N
		} 
	}
}

func openSticky(){
	
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var isValidInput bool = false
	var text string = ""
	fmt.Println("Are there any new tasks to be done today? (Y/N):")
	text, _ = reader.ReadString('\n')
	for  isValidInput == false {
		if text == "Y\r\n" {
			isValidInput = true
			addNewTasks(reader)
		} else if text == "N\r\n" {
			isValidInput = true
		} else {
			fmt.Println("Please enter either Y or N:")
			text, _ = reader.ReadString('\n')
		}
	}
	openSticky()

}

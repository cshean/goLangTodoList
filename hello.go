package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("TODO list:")
	fmt.Println("1: View list")
	fmt.Println("2: Enter New Task")
	fmt.Println("3: Edit Task")
	fmt.Println("4: Delete Task")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

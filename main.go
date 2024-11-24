package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(
`Select one of the following options:

* 1 to Add Task.
* 2 to View Tasks.
* 3 to Delete Task.
* 4 to Update Task.
* 5 to Exit the program.`,
	)
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdout)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
	}
}

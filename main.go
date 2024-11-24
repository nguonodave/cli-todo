package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Task struct {
	ID          int
	Description string
	Priority    int
}

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

	tasks := []Task{}
	id := 1

	scanner := bufio.NewScanner(os.Stdout)
	for scanner.Scan() {
		input := scanner.Text()

		if input == "1" {
			tasks = add(tasks, id)
			id++
			fmt.Println(tasks)
		} else if input == "2" {
			view(tasks)
		} else if input == "3" {
			tasks = delete(tasks)
		} else if input == "4" {
			update(tasks)
		} else if input == "5" {
			os.Exit(0)
		}
	}
}

func add(t []Task, id int) []Task {
	var desc string
	fmt.Print("desc: ")
	_, descErr := fmt.Scan(&desc)
	if descErr != nil {
		fmt.Println("an error occured while scanning the description", descErr)
		return t
	}

	var prio int
	fmt.Print("prio: ")
	_, prioErr := fmt.Scan(&prio)
	if prioErr != nil {
		fmt.Println("an error occured while scanning the priority", prioErr)
		return t
	}

	task := Task{
		ID:          id,
		Description: desc,
		Priority:    prio,
	}

	t = append(t, task)
	return t
}

func view(t []Task) {
	sort.Slice(t, func(i, j int) bool {
		return t[i].Priority < t[j].Priority
	})

	for _, v := range t {
		fmt.Println(v)
	}
}

func update(t []Task) {
	var id int
	fmt.Print("enter task id: ")
	_, idErr := fmt.Scan(&id)
	if idErr != nil {
		fmt.Println("an error occured while scanning the priority", idErr)
		return
	}

	if id > len(t) {
		fmt.Println("id not found")
		return
	}

	taskToUpdate := t[id-1]

	var desc string
	fmt.Print("update desc: ")
	_, descErr := fmt.Scan(&desc)
	if descErr != nil {
		fmt.Println("an error occured while scanning the description", descErr)
		return
	}

	var prio int
	fmt.Print("update prio: ")
	_, prioErr := fmt.Scan(&prio)
	if prioErr != nil {
		fmt.Println("an error occured while scanning the priority", prioErr)
		return
	}

	taskToUpdate.Description = desc
	taskToUpdate.Priority = prio
	t[id-1] = taskToUpdate
}

func delete(t []Task) []Task {
	var id int
	fmt.Print("enter task id: ")
	_, idErr := fmt.Scan(&id)
	if idErr != nil {
		fmt.Println("an error occured while scanning the priority", idErr)
		return t
	}

	if id > len(t) {
		fmt.Println("id not found")
		return t
	}

	taskToDelete := t[id-1]
	newtasks := []Task{}

	for _, v := range t {
		if v != taskToDelete {
			newtasks = append(newtasks, v)
		}
	}

	return newtasks
}

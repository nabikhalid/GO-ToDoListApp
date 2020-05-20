package main

import "fmt"

type item struct {
	todo     string
	priority int
}

// no need for getters, just use dot notation

// setters

func (i *item) setToDo(newToDo string) {
	i.todo = newToDo
}

func (i *item) setPriority(newPriority int) {
	i.priority = newPriority
}

type todoList struct {
	list  []item
	title string
}

func (l *todoList) addItem(i item) {
	l.list = append(l.list, i)
}

// func removeItem

// func clearList

// func printList

func main() {
	fmt.Println("To Do List.")

	newList := todoList{title: "New List"}

	newItem := item{todo: "learn go", priority: 10}

	newList.addItem(newItem)

	fmt.Println(newList.list)

}

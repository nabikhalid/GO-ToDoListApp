package main

import (
	"errors"
	"fmt"
)

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
	// size  int, len(list)
}

func (l *todoList) addItem(i item) {
	l.list = append(l.list, i)
}

// returns struct, fix dat
func (l *todoList) removeItem(i item) (*item, error) {
	for index, value := range l.list {
		if value == i {
			// remove
			// l.list[index] = nil

			l.list = append(l.list[:index], l.list[index+1:]...) // according to Stack Overflow, this is how you remove an item from a list while maintaining order

			return &value, nil
		}
	}
	return nil, errors.New("list does not contain given item")
}

func (l *todoList) clearList() {
	l.list = nil
}

func (l *todoList) printList() {

	fmt.Println(l.title)

	for i := 0; i < len(l.list); i++ {
		fmt.Println(l.list[i].todo, "- Priority Level:", l.list[i].priority)
	}

	fmt.Println("End of List.")
}

// func (l *todoList) prioritySort() {

// }

func main() {

	fmt.Println("To Do List Demo")

	// newList := todoList{title: "New List"}

	// newItem := item{todo: "learn go", priority: 10}

	// newList.addItem(newItem)

	// fmt.Println(newList.list)

	newlist := todoList{title: "Nabi's ToDo List"}

	item1 := item{todo: "learn go", priority: 2}
	item2 := item{todo: "finish this app", priority: 3}
	item3 := item{todo: "read book", priority: 1}

	newlist.addItem(item1)
	newlist.addItem(item2)
	newlist.addItem(item3)

	newlist.printList()

	fmt.Println(newlist.removeItem(item1))
	fmt.Println("FIRST ITEM REMOVED")

	newlist.printList()

	newlist.clearList()
	fmt.Println("LIST CLEARED")

	newlist.printList()

	fmt.Println(newlist.removeItem(item1))

}

// func removeItem(slice []int, s int) []int {
//     return append(slice[:s], slice[s+1:]...)
// }

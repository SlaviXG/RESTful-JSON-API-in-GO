package main

import "fmt"

var currentId int

var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Make the presentation"})
	RepoCreateTodo(Todo{Name: "Organize the meetup"})
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoDeleteTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo with the id %d was not found", id)
}

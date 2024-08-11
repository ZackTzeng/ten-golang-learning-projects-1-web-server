package todo

import (
	"fmt"
)

type Todo struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type TodoStore struct {
	todoList []Todo
}

func NewTodo(subject string, content string) *Todo {
	t := Todo{
		Subject: subject,
		Content: content,
	}
	return &t
}

func NewTodoStore() TodoStore {
	ts := TodoStore{
		todoList: []Todo{},
	}
	return ts
}

func (tl *TodoStore) AddTodo(t Todo) {
	tl.todoList = append(tl.todoList, t)
	tl.PrintList()
}

func (t *Todo) PrintTodo() {
	fmt.Println("{" + t.Subject + ": " + t.Content + "}")
}

func (tl *TodoStore) PrintList() {
	fmt.Println(tl.todoList)
}

func (tl *TodoStore) GetTodoList() []Todo {
	return tl.todoList
}

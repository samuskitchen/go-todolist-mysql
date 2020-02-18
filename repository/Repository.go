package repository

import (
	"github.com/samuskitchen/go-todolist-mysql/domain"
)

type TodoRepo interface {
	CreateItem(todo domain.TodoItemModel) (interface{}, error)
	UpdateItem(todo domain.TodoItemModel) (bool, error)
	DeleteItem(id int) (bool, error)
	GetItemByID(id int) (bool, error)
	GetTodoItems( completed bool) interface{}
}

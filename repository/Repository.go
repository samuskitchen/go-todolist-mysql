package repository

import(
	"context"
	"github.com/samuskitchen/go-todolist-mysql/domain"
)

type RepoTodo interface {
	CreateItem(ctx context.Context, todo domain.TodoItemModel) error
	UpdateItem(ctx context.Context, todo domain.TodoItemModel) error
	DeleteItem(ctx context.Context, id int) error
	GetItemByID(ctx context.Context, id int) (bool, error)
	GetCompletedItems(ctx context.Context,  completed bool) ([]domain.TodoItemModel, error)
	GetIncompleteItems(ctx context.Context,  inCompleted bool) ([]domain.TodoItemModel, error)
}

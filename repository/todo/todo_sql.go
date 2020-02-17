package todo

import (
	"github.com/jinzhu/gorm"
	repo "github.com/samuskitchen/go-todolist-mysql/repository"
)

// NewSQLDomainRepo retunrs implement of post repository interface
func NewSQLTodoRepo(Conn *gorm.DB) repo.RepoTodo {
	return &sqlTodoRepo{
		Conn: Conn,
	}
}

type sqlTodoRepo struct {
	Conn *gorm.DB
}

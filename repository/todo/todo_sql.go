package todo

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/samuskitchen/go-todolist-mysql/domain"
	repo "github.com/samuskitchen/go-todolist-mysql/repository"
	log "github.com/sirupsen/logrus"
)

// NewSQLTodoRepo returns implement of post repository interface
func NewSQLTodoRepo(Conn *gorm.DB) repo.TodoRepo {
	return &sqlTodoRepo{
		Conn: Conn,
	}
}

type sqlTodoRepo struct {
	Conn *gorm.DB
}

func (sql *sqlTodoRepo) CreateItem(todo domain.TodoItemModel) (interface{}, error) {

	sql.Conn.Create(&todo)
	result := sql.Conn.Last(&todo)

	if result.Error != nil {
		log.Warn("TodoItem problem to create in database")
		return nil, errors.New("TodoItem problem to create in database")
	}

	return result.Value, nil
}

func (sql *sqlTodoRepo) UpdateItem(todo domain.TodoItemModel) (bool, error) {
	todoUpdate := &domain.TodoItemModel{}

	sql.Conn.First(todoUpdate, todo.Id)
	todoUpdate.Completed = todo.Completed

	result := sql.Conn.Save(todoUpdate)

	if result.Error != nil {
		log.Warn(result.Error.Error())
		return false, errors.New("TodoItem problem to update in database")
	}

	return true, nil
}

func (sql *sqlTodoRepo) DeleteItem(id int) (bool, error) {
	todo := &domain.TodoItemModel{}
	sql.Conn.First(&todo, id)

	result := sql.Conn.Delete(&todo)

	if result.Error != nil {
		log.Warn(result.Error.Error())
		return false, errors.New("TodoItem problem to delete in database")
	}

	return true, nil
}

func (sql *sqlTodoRepo) GetItemByID(id int) (bool, error) {
	todo := &domain.TodoItemModel{}
	result := sql.Conn.First(&todo, id)

	if result.Error != nil {
		log.Warn("TodoItem not found in database")
		return false, errors.New("TodoItem not found in database")
	}

	return true, nil
}

func (sql *sqlTodoRepo) GetTodoItems(completed bool) interface{} {
	var todo []domain.TodoItemModel
	todoItems := sql.Conn.Where("completed = ?", completed).Find(&todo).Value
	return todoItems
}

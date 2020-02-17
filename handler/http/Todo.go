package http

import (
	"encoding/json"
	"github.com/samuskitchen/go-todolist-mysql/driver"
	"github.com/samuskitchen/go-todolist-mysql/repository"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewTodoHandler(db *driver.DB) *Todo {
	return &Todo{
		repo: repository.NewSQLTodoRepo(db.SQL),
	}
}

type Todo struct {
	repo repository.RepoTodo
}

func (rp *Todo) CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")
	todo := &TodoItemModel{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)

}

package http

import (
	"github.com/gorilla/mux"
	"github.com/samuskitchen/go-todolist-mysql/domain"
	"github.com/samuskitchen/go-todolist-mysql/domain/request"
	"github.com/samuskitchen/go-todolist-mysql/driver"
	"github.com/samuskitchen/go-todolist-mysql/handler/command"
	"github.com/samuskitchen/go-todolist-mysql/repository"
	"github.com/samuskitchen/go-todolist-mysql/repository/todo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func NewTodoHandler(db *driver.DB) *Todo {
	return &Todo{
		repo: todo.NewSQLTodoRepo(db.SQL),
	}
}

type Todo struct {
	repo repository.TodoRepo
}

func (rp *Todo) CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")

	todo := domain.TodoItemModel{Description: description, Completed: false}

	result, err := rp.repo.CreateItem(todo)

	if err != nil {
		command.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	command.RespondWithJSON(w, http.StatusOK, result)
}

func (rp *Todo) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := rp.repo.GetItemByID(id)
	if err != nil {
		command.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		completed, _ := strconv.ParseBool(r.FormValue("completed"))
		log.WithFields(log.Fields{"Id": id, "Completed": completed}).Info("Updating TodoItem")

		todo := domain.TodoItemModel{Id: id, Description: "", Completed: completed}
		result, err := rp.repo.UpdateItem(todo)

		if err != nil {
			command.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		update := request.Update{Update: result}
		command.RespondWithJSON(w, http.StatusOK, update)
	}
}

func (rp *Todo) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := rp.repo.GetItemByID(id)
	if err != nil {
		command.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting TodoItem")

		result, err := rp.repo.DeleteItem(id)

		if err != nil {
			command.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		delete := request.Delete{Delete: result}
		command.RespondWithJSON(w, http.StatusOK, delete)
	}
}

func (rp *Todo) GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get completed TodoItems")
	completedTodoItems := rp.repo.GetTodoItems(true)

	command.RespondWithJSON(w, http.StatusOK, completedTodoItems)
}

func (rp *Todo) GetIncompleteItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Incomplete TodoItems")
	IncompleteTodoItems := rp.repo.GetTodoItems(false)

	command.RespondWithJSON(w, http.StatusOK, IncompleteTodoItems)
}

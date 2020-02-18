package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samuskitchen/go-todolist-mysql/domain"
	dh "github.com/samuskitchen/go-todolist-mysql/handler/http"
	//"github.com/samuskitchen/go-todolist-mysql/domain"
	"github.com/samuskitchen/go-todolist-mysql/driver"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func main() {

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer connection.SQL.Close()

	connection.SQL.Debug().DropTableIfExists(&domain.TodoItemModel{})
	connection.SQL.Debug().AutoMigrate(&domain.TodoItemModel{})

	log.Info("Starting TodoList API server")
	router := mux.NewRouter()

	tHandler := dh.NewTodoHandler(connection)
	router.HandleFunc("/health", HealThz).Methods("GET")
	router.HandleFunc("/todo", tHandler.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", tHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/todo/{id}", tHandler.DeleteItem).Methods("DELETE")
	router.HandleFunc("/todo-completed", tHandler.GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", tHandler.GetIncompleteItems).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)


	http.ListenAndServe(":8085", handler)
}

func HealThz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}


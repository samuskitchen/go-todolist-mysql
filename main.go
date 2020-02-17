package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

	//connection.SQL.Debug().DropTableIfExists(&domain.TodoItemModel{})
	//connection.SQL.Debug().AutoMigrate(&domain.TodoItemModel{})


	log.Info("Starting TodoList API server")
	router := mux.NewRouter()

	dHandler := dh.NewTodoHandler(connection)
	router.Handle("/", domainRouter(dHandler))

	http.ListenAndServe(":8080", router)
}

func domainRouter(dHandler *dh.Todo) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", HealThz).Methods("GET")
	r.HandleFunc("/todo", dHandler.CreateItem).Methods("POST")

	return r
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


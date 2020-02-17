package main

import (
	
	"fmt"
	"github.com/gorilla/mux"
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

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUser, "", dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer connection.SQL.Close()


	log.Info("Starting TodoList API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealThz).Methods("GET")
	http.ListenAndServe(":8080", router)
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


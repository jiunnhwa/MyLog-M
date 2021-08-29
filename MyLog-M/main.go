package main

import (
	"MyLog-M/driver/sqlite"
	delivery "MyLog-M/internal/delivery/http"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := sqlite.Open("log.db")
	defer sqlite.Close(db)

	handler := delivery.New(db)

	//VIEWS
	http.HandleFunc("/", home)

	http.HandleFunc("/api/log", handler.MyLog)
	http.HandleFunc("/api/tail", handler.MyTail)
	http.HandleFunc("/api/view", handler.MyView)

	log.Println(http.ListenAndServe(":9000", nil))
}

//handles home page, updates the view data and serve
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to MyLog-As-A-Service.")
}

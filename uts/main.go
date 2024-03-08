package main

import (
	"fmt"
	"log"
	"net/http"
	"uts/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/det/rooms", controllers.GetAllDetailRooms).Methods("GET")

	router.HandleFunc("/rooms", controllers.InsertToRoom).Methods("POST")

	http.Handle("/", router)
	fmt.Println("connected to port 8080")
	log.Println("connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

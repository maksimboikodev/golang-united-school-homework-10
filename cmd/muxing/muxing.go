package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GolangUnited/helloweb/cmd/handlers"
	"github.com/gorilla/mux"
)

func AddHandlers(router *mux.Router) {
	router.HandleFunc("/name/{PARAM}", handlers.NameParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handlers.BadParam).Methods(http.MethodGet)
	router.HandleFunc("/data", handlers.BodyParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", handlers.HeadersParam).Methods(http.MethodPost)
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.
main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	AddHandlers(router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 8081
	}

	Start(host, port)
}

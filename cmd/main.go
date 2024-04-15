package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bconskri/go_final_project/database"
	"github.com/bconskri/go_final_project/handlers"
	"github.com/go-chi/chi/v5"
)

func getPort() string {
	const default_port = "7540"
	if val, exists := os.LookupEnv("TODO_PORT"); exists {
		return ":" + val
	}
	return ":" + default_port
}

func main() {
	const webDir = "../web"

	database.ConnectDB()

	myHandler := chi.NewRouter()

	myHandler.Mount("/", http.FileServer(http.Dir(webDir)))
	myHandler.Get("/api/nextdate", handlers.NextDateGET)
	myHandler.Post("/api/task", handlers.TaskPost)
	myHandler.Get("/api/tasks", handlers.TasksRead)
	myHandler.Get("/api/task", handlers.TaskReadByID)
	myHandler.Put("/api/task", handlers.TaskUpdate)
	myHandler.Post("/api/task/done", handlers.TaskDone)
	myHandler.Delete("/api/task", handlers.TaskDelete)

	s := &http.Server{
		Addr:           getPort(),
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

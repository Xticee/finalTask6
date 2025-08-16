package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	file, err := os.OpenFile("../logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return
	}
	defer file.Close()

	logger := log.New(file, "app ", log.Lshortfile|log.LstdFlags)

	app := server.CreateApp(logger)

	if err = app.Server.ListenAndServe(); err != nil {
		app.Logger.Fatal("Unable to start server")
	}
}

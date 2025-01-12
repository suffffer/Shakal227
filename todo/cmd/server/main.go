package main

import (
	"github.com/suffffer/Shakal227/internal/app"
	"log"
)

var err error

func main() {
	err = app.RunApp()
	log.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
}

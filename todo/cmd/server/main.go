package main

import (
	"github.com/suffffer/Shakal227/internal/app"
	"log"
)

var err error

func main() {
	err = app.RunApp()
	if err != nil {
		log.Fatal(err)
	}
}

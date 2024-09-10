package main

import (
	"gotask/app"
	"log"
	"os"
)

func main() {
	println("Hello, World!")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
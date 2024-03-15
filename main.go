package main

import (
	"fizzbuzzlbc/application"
	"log"
)

func main() {

	app, err := application.NewApplication()
	if nil != err {
		log.Fatal(err)
	}

	if err := app.Run(); nil != err {
		log.Fatal(err)
	}

}

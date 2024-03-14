package main

import (
	"fizzbuzzlbc/application"
	"log"
)

func main() {

	if err := application.NewApplication().Run(); nil != err {
		log.Fatal()
	}

}

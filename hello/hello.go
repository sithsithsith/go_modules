package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zsithh/go_tutorials/go_module/greetings"
)

func main() {
	//setting up logs
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Get arguments from user input
	args := os.Args[1:]
	name := ""
	if len(args) > 0 {
		name = args[0]
	}

	//Get message from greetings package
	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	//Err handling
	if name == "" {
		return "", errors.New("no name provided")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

package main

import (
	"errors"
	"fmt"
)

func (tl *TodosList) Write() error {
	// Add todo
	return nil
}

func (tl TodosList) Read() error {
	if len(tl) < 1 {
		return errors.New("No todos to read")
	}
	fmt.Println("Reading todos:")
	for _, td := range tl {
		fmt.Println(td)
	}
	return nil
}

func (tl TodosList) Delete() (todo, error) {
	// Delete todo
	return todo{}, nil
}

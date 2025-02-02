package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// type todo struct {
// 	Id        int       `json:"id"`
// 	Title     string    `json:"title"`
// 	Body      string    `json:"body"`
// 	CreatedAt time.Time `json:"created_at"`
// }

func (tl TodosList) Write() error {
	td := todo{
		Id:        3,
		Title:     "third todo",
		Body:      "This is my third todo",
		CreatedAt: time.Now(),
	}
	tl = append(tl, td)
	bs, err := json.Marshal(&tl)
	if err != nil {
		return err
	}
	err = os.WriteFile("testdata.json", bs, 0666)
	if err != nil {
		return err
	}
	fmt.Println(tl)
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

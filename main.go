package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type TodosList []todo

func main() {
	tl := &TodosList{}

	// Read json file into data
	data, err := os.ReadFile("testdata.json")
	if err != nil {
		log.Fatal(err)
	}

	// fill todoslist with json file
	err = json.Unmarshal(data, tl)
	if err != nil {
		log.Fatal(err)
	}

	// switch for cl args

	// loop over todos
	for _, t := range *tl {
		fmt.Println(t)
	}
}

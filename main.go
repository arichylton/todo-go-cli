package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func (ct *CustomTime) UnmarshalJSON(bs []byte) error {
	str := strings.Trim(string(bs), `"`)
	if str == "" || str == "null" {
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02", str)
	// Parse the string into a time object with the same format used in MarshalJSON
	if err != nil {
		return err
	}
	*ct = CustomTime(parsedTime)
	return nil
}

type todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type TodosList []todo

func main() {
	tl := &TodosList{}

	// Check if the file exists, and if not, create it with an empty list
	_, err := os.Stat("testdata.json")
	if os.IsNotExist(err) {
		// File doesn't exist, create a new one with an empty TodosList
		fmt.Println("Warning: 'testdata.json' not found. Creating a new file with an empty list.")
		file, err := os.Create("testdata.json")
		if err != nil {
			log.Fatal("Unable to create file: ", err)
		}
		defer file.Close()

		// Initialize the file with an empty JSON array
		emptyJSON := []byte("[]")
		_, err = file.Write(emptyJSON)
		if err != nil {
			log.Fatal("Unable to write to file: ", err)
		}
	} else if err != nil {
		// Other errors (e.g., permission issues)
		log.Fatal("Unable to check file existence: ", err)
	}

	// Read the file into data
	data, err := os.ReadFile("testdata.json")
	if err != nil {
		log.Fatal("Unable to read file: ", err)
	}

	// Fill todoslist with the contents of the file
	err = json.Unmarshal(data, tl)
	if err != nil {
		log.Fatal("Unable to unmarshal JSON: ", err)
	}

	// Proceed with the cmdSwitcher function if no errors
	err = cmdSwitcher(tl, os.Args)
	if err != nil {
		log.Fatal("cmd switcher: ", err)
	}
}

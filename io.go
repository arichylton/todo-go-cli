package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Custom type for formatting time
type CustomTime time.Time

// Implement the JSON Marshaler interface for CustomTime
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	// Format your time as desired (example: "Jan 02, 2006 03:04 PM")
	return []byte(`"` + time.Time(ct).Format("2006-01-02") + `"`), nil
}

// type todo struct {
// 	Id        int       `json:"id"`
// 	Title     string    `json:"title"`
// 	Body      string    `json:"body"`
// 	CreatedAt time.Time `json:"created_at"`
// }

func (tl TodosList) Write(title string, body string) error {
	id := 0
	if len(tl) > 0 {
		id = tl[len(tl)-1].Id + 1
	}

	td := todo{
		Id:        id,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}
	tl = append(tl, td)
	bs, err := json.Marshal(&tl)
	if err != nil {
		return err
	}
	err = os.WriteFile("testdata.json", bs, 0666)
	if err != nil {
		fmt.Println("something wrong")
		return err

	}
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

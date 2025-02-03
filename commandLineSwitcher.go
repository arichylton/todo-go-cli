package main

import (
	"log"
)

func cmdSwitcher(tl *TodosList, args []string) error {
	args = args[1:]
	action := ""
	if len(args) > 1 {
		action = args[0]
		if action == "w" {
			actionSwitcher(tl, action, args[1:]...)
		}
	} else {
		actionSwitcher(tl, action)
	}

	return nil
}

func actionSwitcher(tl *TodosList, action string, args ...string) {
	switch action {
	case "r":
		tl.Read()
	case "w":
		err := tl.Write(args[0], args[1])
		if err != nil {
			log.Panic("Something bad happened ", err)
		}
	case "d":
		tl.Delete()
	default:
		tl.Read()
	}
}

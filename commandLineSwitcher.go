package main

import "fmt"

func cmdSwitcher(tl *TodosList, args []string) error {

	args = args[1:]
	action := ""
	if len(args) > 0 {
		action = args[0]
	}
	actionSwitcher(action)

	return nil
}

func actionSwitcher(action string) {
	switch action {
	case "r":
		fmt.Println("reading")
	case "w":
		fmt.Println("writing")
	case "d":
		fmt.Println("deleting")
	default:
		fmt.Println("reading default")
	}
}

package main

func cmdSwitcher(tl *TodosList, args []string) error {
	args = args[1:]
	action := ""
	if len(args) > 0 {
		action = args[0]
	}
	actionSwitcher(tl, action)

	return nil
}

func actionSwitcher(tl *TodosList, action string) {
	switch action {
	case "r":
		tl.Read()
	case "w":
		tl.Write()
	case "d":
		tl.Delete()
	default:
		tl.Read()
	}
}

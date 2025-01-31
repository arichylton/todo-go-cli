package main

func (tl *TodosList) Write() error {
	// Add todo
	return nil
}

func (tl TodosList) Read() error {
	// Read todos list
	return nil
}

func (tl TodosList) Delete() (todo, error) {
	// Delete todo
	return todo{}, nil
}

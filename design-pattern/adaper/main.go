package main

import "fmt"

type computer interface {
	insertSquareComputer()
}

type client struct{}

func (c *client) insertSquareComputer(computer computer) {
	computer.insertSquareComputer()
}

type mac struct{}

func (m *mac) insertSquareComputer() {
	fmt.Println("Insert square port into mac machine")
}

type windows struct{}

func (m *windows) insertCircleComputer() {
	fmt.Println("Insert circle port into windows machine")
}

type windowsAdapter struct {
	windows windows
}

func (w *windowsAdapter) insertSquareComputer() {
	w.windows.insertCircleComputer()
}

func main() {
	c := client{}
	m := mac{}

	c.insertSquareComputer(&m)

	w := windows{}
	wa := windowsAdapter{windows: w}

	c.insertSquareComputer(&wa)
}

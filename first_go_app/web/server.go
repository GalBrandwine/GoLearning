package web

import (
	"fmt"
	"io/ioutil"
)

/*
Using this tutorial: https://golang.org/doc/articles/wiki/
*/

// Page a struct of two fields representing the title and body
type Page struct {
	Title string
	Body  []byte
}

// Init initiate server
func Init() {
	fmt.Println("Server initiation function")

}

// Save is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

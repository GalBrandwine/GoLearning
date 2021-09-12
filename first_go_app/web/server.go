package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/internal/logger"
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
	http.HandleFunc("/", handler)
	http.HandleFunc("/backdoor", handlerBackDoor)
	logger.LogFatalError(http.ListenAndServe(":8080", nil))
}

// Save is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage constructs the file name from the title parameter,
// reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values.
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Handler HTTP request callback
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	homePage, _ := ioutil.ReadFile("web/html/page1.html")
	fmt.Fprintf(w, string(homePage))
}

func handlerBackDoor(w http.ResponseWriter, r *http.Request) {
	logger.LogInfo("Exiting via backdoor")
	os.Exit(0)
}

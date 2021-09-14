package web_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/web"
)

func TestSave(t *testing.T) {

	// Prepare
	p := web.Page{}
	p.Title = "Gal is the king"
	p.Body = []byte("Have a king chair")

	// Run
	p.Save()

	// Test
	data, err := ioutil.ReadFile(p.Title + ".txt")
	if err != nil {
		fmt.Errorf(err.Error())
		t.Fail()
	}

	test := true
	for i, ch := range data {
		if ch != p.Body[i] {
			test = false
		}
	}
	if !test {
		t.Fail()
	}
}

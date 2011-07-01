package main

import (
       "fmt"
       "io/ioutil"
       "os"
)

type Page struct {		  
     Title string		  
     Body []byte // []byte and not string because io libs work on bytes
}
// Page struct describes what's in memory
// How do we store it on disk?

func (p *Page) save() os.Error {
	// We're passing the return value of WriteFile which has type os.Error
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, os.Error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) // err stores os.Error from ReadFile
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body:body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("I'm a test page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
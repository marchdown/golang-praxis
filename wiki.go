package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"http"
	"template"
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

const lenPath = len("/view/")
var port = flag.String("port", "2048", "TCP port for inbound connections")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
//	p1 := &Page{Title: "TestPage", Body: []byte("I'm a test page.")}
//	p1.save()
//	p2, _ := loadPage("TestPage")
	flag.Parse()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":" + *port, nil)
}


//now write the last two handlers.

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFile("edit.html", nil)
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
//code code
}
package main

import (
	"regexp"
	"flag"
	"io/ioutil"
	"os"
	"http"
	"template"
)

type Page struct {		  
     Title string		  
     Body []byte // []byte and not string because io libs work on bytes
}

var templates = make(map[string]*template.Template)
func init () {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
	}
}

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
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[lenPath:]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w,p)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
func getTitle(w http.ResponseWriter, r *http.Request) (title string, err os.Error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		http.NotFound(w, r)
		err = os.NewError("Invalid Page Title")
	}
	return
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
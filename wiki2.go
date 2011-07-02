package main

import ( 
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"http" )

type Page struct {
	Title string
	Body []byte
}

//save page
func (p *Page) save() os.Error { // what are parameters and what are arguments?
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//open page
func openPage(title string) (*Page, os.Error) { //from disk to memory
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	//check for errors
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

const lenPath = len("/view/")
var port = flag.String("port", "2048", "Listening port")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:] //tail of request path after lenPath chars
	p, _ := openPage(title)
	fmt.Fprintf(w, "<em>%s</em><br><p>%s</p>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":" + *port, nil)
}

	

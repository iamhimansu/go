package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)
type User struct {
	Name string
	XSS string
	Age int
	Meta UserMeta
}

type UserMeta struct{
	Visits int
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl := filepath.Join("exp", "index.gohtml")
	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing template.", http.StatusInternalServerError)
		return
	}

	user := User{
		Name: "susan",
		XSS: `<script>alert('hello')</script>`,
		Age: 20,
		Meta: UserMeta{
			Visits: 90,
		},
	}

	err = t.Execute(w, user)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><div><a href=\"mailto:iamhimanshu7102@gmail.com\">iamhimanshu7102@gmail.com</a></div>")
}
func x(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("welcome to x"))
	w.Write([]byte(r.URL.RawQuery))
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "Page not found")

// 	}
// }
// type RequestHandler struct{}
// func (request *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/x":
// 		x(w, r)
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "Page not found")

// 	}
// }
func main() {
	request := chi.NewRouter()
	request.Get("/", homeHandler)
	request.Get("/contact", contactHandler)
	request.Get("/x", x)
	request.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", request)
	fmt.Println("Server started")
}

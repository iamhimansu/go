package main

import (
	"html/template"
	"os"
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
func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "susan",
		XSS: `<script>alert('hello')</script>`,
		Age: 20,
		Meta: UserMeta{
			Visits: 90,
		},
	}

	t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}

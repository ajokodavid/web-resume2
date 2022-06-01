package main

import (
	"fmt"
	
	"os"

	"net/http"

	"html/template"
)

var tmpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	type userDetails struct {
		Email string
		Password string
		label string
		label2 string
	}

	var user_details = userDetails{
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
		label: "Email",
		label2: "Password",
	}

	tmpl = template.Must(template.ParseFiles("template/index.html"))

	tmpl.Execute(w,user_details)
}

func resume(w http.ResponseWriter, r *http.Request) {
	type myDetails struct{
		Work string
		Place string
		Email string
		PhoneNum string
	}

	my_details := myDetails{
		Work: "Web Developer",
		Place: "Lagos, Nigeria",
		Email: "bdave3674@gmail.com",
		PhoneNum: "07013419386, 07069185764",
	}

	tmpl = template.Must(template.ParseFiles("template/resume.html"))

	tmpl.Execute(w, my_details)
}

func main() {
	var mux = http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/resume", resume)

	fmt.Println("Server started....")

	port := os.Getenv("PORT")

	// fmt.Print("Listening on :" + port)

	http.ListenAndServe(":" + port, mux)
}
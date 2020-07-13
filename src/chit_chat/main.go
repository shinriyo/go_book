package main

import (
	"net/http"
	"html/template"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	// リスト2.3
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.ReadRequest) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}

	templates: = template.Must(template.ParseFiles(files...))
		threads, err := data.Threads(); if err == nil {
			templates.ExecuteTemplate(w, "lauyout", threads)
		}
	}
	templates := templates.Must

}


// P38
func index2(w http.ResponseWriter, r *http.ReadRequest) {
	threads, err := data.Threads(); if err == nil {
	_, err := session(w, r)
	piblic_tmpl_files := []string{"templates/layout.html",
										"templates/navbar.html",
										"templates/index.html"}
	private_tmpl_files := []string{"templates/layout.html",
										"templates/navbar.html",
										"templates/index.html"}

	var templates *template.Template
	if err != nil {
		templates =	templates.Must(template.ParseFiles(private_tmpl_files...))
	} else {
		templates = templates.Must(template.ParseFiles(public_tmpl_files...))
	}
}

// P44
func index3(writer http.ResponseWriter, request *http.ReadRequest) {
	threads, err := data.Threads(); if err == nil {
	_, err := session(writer, request)
	var templates *template.Template
	if err != nil {
		generateHTML(writer, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(writer, threads, "layout", "private.navbar", "index")
	}
}

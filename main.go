package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

type User struct {
	Id       int
	Username string
	Email    string
}

func HelloWeb(writer http.ResponseWriter, request *http.Request) {
	jhon := User{
		Id:       1,
		Username: "Jhon Doe Again Again",
		Email:    "jhondoe@mail.com",
	}

	tmpl.ExecuteTemplate(writer, "Index", jhon)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWeb)

	staticDirectory := http.Dir("assets")
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(staticDirectory)))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

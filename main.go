package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	route.HandleFunc("/hello", hello).Methods("GET")
	route.HandleFunc("/", homePage).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/add-project", blogPage).Methods("GET")
	route.HandleFunc("/project-detail/{index}", blogDetail).Methods("GET")
	route.HandleFunc("/send-data-add", sendDataAdd).Methods("POST")

	fmt.Println("Server running on port 8000")
	http.ListenAndServe("localhost:8000", route)

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/detail-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func sendDataAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	projectName := r.PostForm.Get("project-name")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	description := r.PostForm.Get("desc-project")
	var techno []string
	techno = r.Form["techno"]
	uploadImg := r.PostForm.Get("Imageee")

	fmt.Println("Project Name : " + projectName)
	fmt.Println("Start Date : " + startDate)
	fmt.Println("End Date : " + endDate)
	fmt.Println("Description : " + description)
	fmt.Println(techno)
	fmt.Println("Upload Image : " + uploadImg)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

package main

import (
	"connection"
	"customhttp"

	"encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"

	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	customhttp.HandleFunc(router, "/connect/{user}/{psw}", Connect).Methods("GET")
	customhttp.HandleFunc(router, "/user-agent", UserAgent).Methods("GET")
	customhttp.HandleFunc(router, "/comment", GetComments).Methods("GET")
	customhttp.HandleFunc(router, "/comment", AddComments).Methods("POST")
	customhttp.HandleFunc(router, "/connect", connection.Connect).Methods("POST")

	log.Println("Server start !")
	log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(router)))
}

func Connect(w http.ResponseWriter, r *http.Request) {
	USER := "CalvinKim"
	PSW := "Jc8b&RM52AL"
	params := mux.Vars(r)
	log.Println("Connecting...")
	user, _ := params["user"]
	psw, _ := params["psw"]
	w.Header().Add("x-user", USER)
	w.Header().Add("x-psw", PSW)
	if user == USER && psw == PSW {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`Utiliser le mot de passe trouver`))
	} else {
		w.WriteHeader(403)
		_, _ = w.Write([]byte(`Forbidden`))
	}
}

func UserAgent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("user-agent", "toto")
	user := request.Header.Get("user-agent")
	log.Println(user)
	if user == "toto" {
		writer.WriteHeader(200)
		_, _ = writer.Write([]byte(`felicitation ! vous êtes connecté !`))
	} else {
		writer.WriteHeader(403)
		_, _ = writer.Write([]byte(`Forbidden`))
	}
}

type Comment struct {
	Comment string `json:"comment"`
}

var comments = make([]Comment, 0)

func AddComments(w http.ResponseWriter, r *http.Request) {
	log.Println("add a comment")
	var comment Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)
	comments = append(comments, comment)
	w.WriteHeader(200)
}

func GetComments(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comments)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Gere toute les routes du serveur HTTP
func handleRequest(router *mux.Router) {
	router.HandleFunc("/", nihao).Methods("GET")
	router.HandleFunc("/get/test", test).Methods("GET")
	router.HandleFunc("/getiaresponse/{phrase}", getIAResponse).Methods("GET")

	/*	router.HandleFunc("/people/{id}", test2).Methods("POST")
		router.HandleFunc("/people/{id}", test3).Methods("DELETE")*/
}

//Page d'acceuil de l'api
func nihao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, welcome to the poliright golang api !</h1>")
}

//Recupere une phrase et renvoie un json
func getIAResponse(w http.ResponseWriter, r *http.Request) {
	log := r.RemoteAddr + r.URL.String() + " " + r.Method
	writeFile(log)
	urlPart := strings.Split(r.URL.Path, "/getiaresponse/")

	// Envoie la phrase au client
	_, err := con.Write([]byte(urlPart[1] + "\n"))
	fmt.Print(err)
	d := json.NewDecoder(con)
	IA := ia{}
	d.Decode(&IA)
	json.NewEncoder(w).Encode(IA)
}

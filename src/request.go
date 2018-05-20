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
	/* 	var one = []byte{}
	   	_, errr := con.Read(one)
	   	fmt.Println(errr)
	   	if errr != nil {
	   		con.Close()
	   		fmt.Println(errr)
	   		writeFile(errr.Error())
	   		env := env{}
	   		readEnv(&env)
	   		iaConnect(&env)
	   		m := er{er: "Erreur de connexion à l'ia, reconnexion en cours."}
	   		json.NewEncoder(w).Encode(m)
		   } */
	log := r.RemoteAddr + r.URL.String() + " " + r.Method
	writeFile(log)
	urlPart := strings.Split(r.URL.Path, "/getiaresponse/")

	fmt.Printf("%s", urlPart[1])
	IA := ia{}
	// Envoie la phrase au client
	if con != nil {
		n, err := con.Write([]byte(urlPart[1] + "\n"))

		fmt.Println(n)
		if err != nil {
			IA.Error = "Connexion impossible à l'IA."
			con.Close()
		} else {
			IA.Error = "none"
		}
	} else {
		IA.Error = "Connexion impossible à l'IA."
	}

	if con != nil {
		d := json.NewDecoder(con)
		d.Decode(&IA)
	}
	json.NewEncoder(w).Encode(IA)
}

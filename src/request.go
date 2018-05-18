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
	fmt.Printf(r.URL.Path)
	urlPart := strings.Split(r.URL.Path, "/getiaresponse/")

	fmt.Printf(urlPart[1])
	// Envoie la phrase au client
	_, err := con.Write([]byte(urlPart[1] + "\n"))
	d := json.NewDecoder(con)
	IA := ia{}
	d.Decode(&IA)
	if err != nil {
		IA.Error = err.Error()
		//con.Close()
		//		writeFile(err.Error())
		/* 		env := env{}
		   		readEnv(&env)
		   		iaConnect(&env) */
	} else {
		IA.Error = "none"
	}
	json.NewEncoder(w).Encode(IA)
}

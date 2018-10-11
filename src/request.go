package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

//Gere toute les routes du serveur HTTP
func handleRequest(router *mux.Router) {
	router.HandleFunc("/", nihao).Methods("GET")
	//router.HandleFunc("/get/test", test).Methods("GET")
	router.HandleFunc("/getiaresponse/{phrase}", getIAResponse).Methods("GET")

}

//Page d'acceuil de l'api
func nihao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, welcome to the poliright golang api !</h1>")
}

//Recupere une phrase et renvoie un json
func getIAResponse(w http.ResponseWriter, r *http.Request) {
	newurl, _ := url.QueryUnescape(r.URL.String())
	log := r.RemoteAddr + newurl + " " + r.Method
	writeFile(log)

	urlPart := strings.Split(newurl, "/getiaresponse/")

	fmt.Printf("Du site web : \n%s\n", urlPart[1])
	IA := ia{}
	txt := map[string]string{"text": urlPart[1]}
	jsonval, _ := json.Marshal(txt)
	var jsonstr = []byte(jsonval)

	fmt.Printf("Vers l'ia : %s\n", jsonstr)
	response, err := http.NewRequest("POST", "http://163.5.220.30:5000/ia", bytes.NewBuffer(jsonstr))
	checkError(err)

	if err != nil {
		IA.Error = "Connexion impossible à l'IA."
	} else {
		IA.Error = "none"
	}

	client := &http.Client{}
	resp, _ := client.Do(response)
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &IA)

	fmt.Printf("L'IA renvoie : \n%s\n", string(body))
	//	d := json.NewDecoder()
	//d.Decode(&IA)

	json.NewEncoder(w).Encode(IA)

	// Envoie la phrase au client
	// if con != nil {
	// 	n, err := con.Write([]byte(urlPart[1] + "\n"))

	// 	fmt.Println(n)
	// 	if err != nil {
	// 		IA.Error = "Connexion impossible à l'IA."
	// 		con.Close()
	// 	} else {
	// 		IA.Error = "none"
	// 	}
	// } else {
	// 	IA.Error = "Connexion impossible à l'IA."
	// 	con.Close()
	// }

	// if con != nil {
	// 	d := json.NewDecoder(con)
	// 	d.Decode(&IA)
	// }
	// json.NewEncoder(w).Encode(IA)
	// con.Close()
}

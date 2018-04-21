package main

//Ma librairie
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

//Vérifie les erreurs et quitte le programme si il en trouve une
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Lis le fichier env
func readEnv(env *env) {
	fileConfig, erreur := ioutil.ReadFile("env.yaml")
	checkError(erreur)
	err := yaml.Unmarshal(fileConfig, &env)
	checkError(err)
	fmt.Printf(env.PortWebRequest)
}

//Page d'acceuil de l'api
func nihao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, welcome to the poliright golang api !</h1>")
}

//Set le port
func balanceTonPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Page test get.</h1>")
}

//Gere toute les routes du serveur HTTP
func handleRequest(router *mux.Router) {
	router.HandleFunc("/", nihao).Methods("GET")
	router.HandleFunc("/get/test", test).Methods("GET")
	/*	router.HandleFunc("/people/{id}", test1).Methods("GET")
		router.HandleFunc("/people/{id}", test2).Methods("POST")
		router.HandleFunc("/people/{id}", test3).Methods("DELETE")*/
}

func main() {
	fmt.Printf("%s\n%s\n", build, version)
	fmt.Println("Launching server...")

	env := env{}
	readEnv(&env)

	/*port, err := balanceTonPort()
	checkError(err)*/

	go iaConnect(&env)

	r := mux.NewRouter()
	handleRequest(r)
	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
}

//Gere l'echange d'info avec l'IA
func iaConnect(env *env) {
	// listen on all interfaces
	ln, errp := net.Listen("tcp", ":"+env.PortSocket)
	checkError(errp)
	// accept connection on port
	conn, erra := ln.Accept()
	checkError(erra)

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			// output message received
			fmt.Print("Message Received:", string(message))

			//			tab, _ := bufio.NewReader(conn).ReadString('\n')
			//fmt.Printf(tab)
		}
	}
}

package main

//Ma librairie
import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

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

func main() {
	writeFile("Le serveur est up.")

	env := env{}
	readEnv(&env)

	//	port, err := balanceTonPort()
	//	checkError(err)

	iaConnect(&env)

	//	r := mux.NewRouter()
	//	handleRequest(r)
	//	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
	//	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
}

//Gere l'echange d'info avec l'IA
func iaConnect(env *env) {
	// listen on all interfaces
	ln, errp := net.Listen("tcp", ":"+env.PortSocket)
	checkError(errp)

	phrase := "phrase de test"

	// accept connection on port
	conn, erra := ln.Accept()
	checkError(erra)
	con = conn

	// Envoie la phrase au client
	con.Write([]byte(phrase + "\n"))
}

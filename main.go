package main

//Ma librairie
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type IA struct {
	phrase string
}

var (
	Version = "1.0.0"
	Build   = time.Now()
)

func nihao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, welcome to the poliright golang api !</h1>")
}

func balanceTonPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

/*func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}*/

func handleRequest(router *mux.Router) {
	router.HandleFunc("/", nihao).Methods("GET")
	//router.HandleFunc("/get/iaresponse", getIAResponse).Methods("GET")
	/*	router.HandleFunc("/people/{id}", test1).Methods("GET")
		router.HandleFunc("/people/{id}", test2).Methods("POST")
		router.HandleFunc("/people/{id}", test3).Methods("DELETE")*/
}

func main() {
	fmt.Printf("%s\n%s\n", Build, Version)

	fmt.Println("Launching server...")

	port, err := balanceTonPort()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(port)

	//	port := os.Getenv("PORT")
	/*	port := "8080"
			fmt.Printf(port)
		router := mux.NewRouter()
		handleRequest(router)
		//	log.Fatal(http.ListenAndServe(":"+port, router))*/
	//	mux := http.NewServeMux()
	//	http.ListenAndServe(":8000", mux)

	http.HandleFunc("/", nihao)

	// listen on all interfaces
	ln, _ := net.Listen("tcp", port)

	// accept connection on port
	conn, _ := ln.Accept()

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			// output message received
			fmt.Print("Message Received:", string(message))
			newmessage := "Je suis une phrase politique"
			// send new string back to client
			conn.Write([]byte(newmessage))

			tab, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Printf(tab)
		}
	}
}

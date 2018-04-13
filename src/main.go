package main

//Ma librairie
import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
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

func getIAResponse(w http.ResponseWriter, r *http.Request) {
	resp := IA{}
	resp.phrase = "Je suis une phrase politique"

	respJson, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
	//	json.NewEncoder(w).Encode(phrase)
}

/*func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}*/

func handleRequest(router *mux.Router) {
	router.HandleFunc("/get/iaresponse", getIAResponse).Methods("GET")
	/*	router.HandleFunc("/people/{id}", test1).Methods("GET")
		router.HandleFunc("/people/{id}", test2).Methods("POST")
		router.HandleFunc("/people/{id}", test3).Methods("DELETE")*/
}

func main() {
	fmt.Printf("%s\n%s\n", Build, Version)

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

	/*router := mux.NewRouter()
	handleRequest(router)
	log.Fatal(http.ListenAndServe(":8000", router))
	*/
	//	client()
	/*mux := http.NewServeMux()
	mux.HandleFunc("/", nihao)
	http.ListenAndServe(":8000", mux)*/
}

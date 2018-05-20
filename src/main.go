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
	"os/signal"
	"syscall"

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

	//Création d'une channel continue pour catch un signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//Catch un signal
	go goCatchSignal(c)
	go iaConnect(&env)

	r := mux.NewRouter()
	handleRequest(r)
	//	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
}

//Gere l'echange d'info avec l'IA
func iaConnect(env *env) {
	// listen on all interfaces
	ln, errp := net.Listen("tcp", ":"+env.PortSocket)
	defer ln.Close()
	if errp != nil {
		log.Fatalf("Socket listen port %s à renvoyé une erreur,%s", env.PortSocket, errp)
	}
	log.Printf("Listen port en cours: %s", env.PortSocket)

	for {
		// accept connection on port
		conn, erra := ln.Accept()
		if erra != nil {
			log.Fatalln(erra)
			continue
		} else {
			message, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Printf("%s", message)
			checkError(err)

			if message == "ia\n" {
				con = conn
				log.Printf("Connexion accepté : %s.", conn.RemoteAddr().String())
			} else {

			}
		}
	}
	// accept connection on port
	/*conn, erra := ln.Accept()
	fmt.Printf("New socket connection")
	checkError(erra)
	con = conn*/
	/* 	if phrase != "" {
	// Envoie la phrase au client
	conn.Write([]byte(phrase + "\n"))

	d := json.NewDecoder(conn)
	IA := ia{}
	d.Decode(&IA)
	phrase = "" */
	//		fmt.Println(IA)
	//		fmt.Printf(IA.RelationBetween[0][0])
	//	}

}

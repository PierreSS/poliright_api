package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//Vérifie les erreurs des fichiers et quitte le programme si il en trouve une
func checkFileError(err error, str string) {
	if err != nil {
		fmt.Printf("%s\n", str)
	}
}

//Récupère une string et l'écrit dans les logs
func writeFile(body string) {
	//Check body of request
	body = time.Now().Format("2006-01-02 15:04:05 : ") + body
	fmt.Printf("%s\n", body)

	//Create & write EOF
	file, err := os.OpenFile("poli.log", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		file, err = os.Create("poli.log")
		checkFileError(err, "erreur de création/écriture dans poli.log")
	}

	//Count Lines of log & erase if x lines
	f, _ := os.Open("poli.log")
	fileScanner := bufio.NewScanner(f)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	//Recréer le fichier à plus de 20000 lignes
	if lineCount >= 20000 {
		err := os.Remove("poli.log")
		checkFileError(err, "erreur d'effacement de poli.log")
		file, err = os.Create("log")
		checkFileError(err, "erreur de recréation de poli.log")
		defer f.Close()
		fmt.Printf("%d\n", lineCount)
	}

	//Write body in file
	_, err = file.WriteString(body)
	checkFileError(err, "erreur d'écriture dans poli.log")
	_, err = file.WriteString("\n")
	checkFileError(err, "erreur d'écriture dans poli.log")
	//Synchronisation des fichiers et fermement de la requête
	file.Sync()
	defer file.Close()
}

//Catch un signal
func goCatchSignal(c chan os.Signal) {
	sig := <-c
	fmt.Printf("\n%sSortie de programme suite à %s\n", time.Now().Format("2006-01-02 15:04:05 : [Program] : "), sig)
	os.Exit(1)
}

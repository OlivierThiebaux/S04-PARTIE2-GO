package main

import ("log"
		"fmt"
	)

func main() {
	//var task string
	var reponse string
	log.Printf("Ajouter une tâche: ")
	fmt.Scanln(reponse)
	log.Printf("reponse: %v", reponse)
}

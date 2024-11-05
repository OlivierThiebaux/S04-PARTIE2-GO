package main

import ("log"
		"fmt"
	)

func main() {
	//var task string
	var reponse string
	log.Printf("Voulez vous ajouter une tâche (O/N): ")
	fmt.Scanln(reponse)
	log.Printf("reponse: %v", reponse)
	switch reponse {
	case "y", "o","yes","oui":
		scan := bufio.NewScanner(os.Stdin)
        scan.Split(bufio.ScanLines)
        log.Println("Reponse oui")
        log.Printf("Saisir la description de la tâche")
        scan.Scan()
        input := scan.Text()
        task = input
        log.Printf("Valeur saisie : %v", input)
		
	case "n","non","no":
		log.Println ("Réponse non")
	default 
	}
	}

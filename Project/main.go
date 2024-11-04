package main

import (
	"fmt"
)

func main() {

	var task string
	// var name string
	var test string
	//var index int
	//var tasks [50]string

	for {

		//index++
		//fmt.Println(index)
		fmt.Printf("Voulez-vous ajouter une tâche ? (y/n): ")
		fmt.Scan(&test)

		if test != "y" && test != "n" {
			fmt.Println("La valeur n'est ni 'y' ni 'n'")

		}
		if test == "y" {
			fmt.Printf("Saisir le nom de la tache: ")
			fmt.Scan(&task)
			break
		} else if test == "n" {
			break
		}

	}
	fmt.Println("la derniere tâche saisie est : ", task)
}

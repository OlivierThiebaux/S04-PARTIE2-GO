package main

import "log"

func ajouter(a int, b int) int {
	return a + b
}

var somme int

func main() {
	log.Println(`Hello world !!!`)
	somme := ajouter(3, 4)
	log.Printf("la somme de 3 et 4 est : %v\n", somme)
	log.Printf("la multiplication de 3 et 4: %v\n", multiplication(3, 4))
	resultat, err := division(3, 0)
	if err != nil {
		log.Fatalf("alerte, division par zéro : %v", err)
	}
	log.Printf("la division de 3 et 4: %v", resultat)

}

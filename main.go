package main

import "fmt"

func ajouter(a int, b int) int {
	return a + b
}
var somme int

func main() {
	fmt.Println(`Hello world !!!`)
	somme := ajouter(3 , 4)
	fmt.Printf("la somme de 3 et 4 est : %v\n" , somme, "coucou", "arg")
}

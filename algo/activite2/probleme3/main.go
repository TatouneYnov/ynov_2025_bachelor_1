//var personne = []string("bib", "bab", "bob")
// pour i a 0 jusqua i la taille de personne avec un pas de 1
//		afficher message envoye a personne a l'indice
// complexité linéaire

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var personne = []string{"bib", "bab", "bob"}
	for i := 0; i < len(personne); i++ {
		fmt.Println("message envoyé a", personne[i])
	}
	duration := time.Since(start)
	fmt.Printf("\nTemps d'exécution: %v\n", duration)
}

// probleme1:
// var collegue = []string(bib, bab, bob)
// pour i a 0 jusqu'à la taille de collegue avec un pas de 1
// 	afficher collegue a l'indice i
// il est de complexité linéaire

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var collegue = []string{"bib", "bab", "bob"}
	for i := 0; i < len(collegue); i++ {
		fmt.Println(collegue[i])
	}

	duration := time.Since(start)
	fmt.Printf("\nTemps d'exécution: %v\n", duration)
}

//var personne = []string("bib", "bab", "bob")
// pour p a 0 jusqua la taille de personne avec un pas de 1
// 		pour i a 0 jusqua i la taille de personne avec un pas de 1
//			afficher p a envoyé un message a i
// complexité N²

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var personne = []string{"bib", "bab", "bob"}

	for p := 0; p < len(personne); p++ {
		for i := 0; i < len(personne); i++ {
			if personne[p] != personne[i] {
				fmt.Println(personne[p], "a envoyé un message a", personne[i])
			}
		}
	}
	duration := time.Since(start)
	fmt.Printf("\nTemps d'exécution: %v\n", duration)
}

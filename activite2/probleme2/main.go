// var montant []int(33, 22, 11, 44, 55)
// 	pour i de 0 jusqua la taille de montant avec un pas de 1
// 		montant a l'indice i + montant
// 	afficher le montant de depense de la journee
// il est de complexité linéaire

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var montant = []int{11, 22, 33, 44}
	total := 0
	for i := 0; i < len(montant); i++ {
		total += montant[i]
	}
	fmt.Println("la depense de la journée est de:", total)
	duration := time.Since(start)
	fmt.Printf("\nTemps d'exécution: %v\n", duration)
}

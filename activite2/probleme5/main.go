// ENTRÉE: liste d'objets

// POUR chaque élément i de 0 à (taille_liste - 2)
//     POUR chaque élément j de (i + 1) à (taille_liste - 1)
//         SI liste[i] == liste[j] ALORS
//             AFFICHER "Doublon trouvé: " + liste[i]
//             RETOURNER VRAI
//         FIN SI
//     FIN POUR
// FIN POUR

// AFFICHER "Aucun doublon trouvé"
// RETOURNER FAUX

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var liste = []string{"table", "lit", "voiture"}
	found := false
	for i := 0; i < len(liste)-1; i++ {
		for j := i + 1; j < len(liste); j++ {
			if liste[i] == liste[j] {
				fmt.Println("Doublon trouvé:", liste[i])
				found = true
			}
		}
	}
	if !found {
		fmt.Println("Aucun doublon trouvé")
	}
	duration := time.Since(start)
	fmt.Printf("\nTemps d'exécution: %v\n", duration)
}

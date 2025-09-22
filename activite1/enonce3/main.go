package main

//saisir le prix du produit que l'on achete et la somme que l'on donne
// Le programme doit déterminer si le client n'a pas donné assez, s'il a donné exactement la
// bonne somme, ou s'il doit recevoir de la monnaie.
//afficher la couleur du ticket

import "fmt"

func main() {
	var prixproduit, sommedonne float64
	fmt.Print("Quel est le prix du produit ? ")
	fmt.Scanln(&prixproduit)

	fmt.Print("Quelle somme donnez-vous ? ")
	fmt.Scanln(&sommedonne)

	if sommedonne == prixproduit {
		fmt.Println("Montant exact ! Pas de monnaie à rendre.")
		fmt.Println("Ticket : VERT")
	} else if sommedonne < prixproduit {
		manque := prixproduit - sommedonne
		fmt.Printf("Montant insuffisant ! Il manque %.2f€\n", manque)
		fmt.Println("Ticket : ROUGE")
	} else {
		monnaie := sommedonne - prixproduit
		fmt.Printf("Voici votre monnaie : %.2f€\n", monnaie)
		fmt.Println("Ticket : BLEU")
	}
}

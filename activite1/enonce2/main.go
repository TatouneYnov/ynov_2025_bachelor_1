package main

// Énoncé 2 : une station météo enregistre la température
// le matin et l'après-midi. On veut afficher la
// température la plus basse et la plus haute de la
// journée. La station note aussi le nom du technicien
// qui a saisi les mesures

import "fmt"

func main() {
	var tempMatin, tempApresMidi float64
	var technicien string

	fmt.Print("Entrez le nom du technicien : ")
	fmt.Scanln(&technicien)

	fmt.Print("Entrez la température du matin (°C) : ")
	fmt.Scanln(&tempMatin)

	fmt.Print("Entrez la température de l'après-midi (°C) : ")
	fmt.Scanln(&tempApresMidi)

	var tempMin, tempMax float64

	if tempMatin <= tempApresMidi {
		tempMin = tempMatin
		tempMax = tempApresMidi
	} else {
		tempMin = tempApresMidi
		tempMax = tempMatin
	}

	fmt.Println("\n--- Rapport météorologique ---")
	fmt.Println("Technicien :", technicien)
	fmt.Printf("Température du matin : %.1f°C\n", tempMatin)
	fmt.Printf("Température de l'après-midi : %.1f°C\n", tempApresMidi)
	fmt.Printf("Température minimale de la journée : %.1f°C\n", tempMin)
	fmt.Printf("Température maximale de la journée : %.1f°C\n", tempMax)
}

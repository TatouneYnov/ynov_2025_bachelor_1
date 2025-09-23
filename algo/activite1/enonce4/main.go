package main

// On indique la vitesse mesurée de la
// voiture et la vitesse maximale autorisée. Le programme
//doit dire si le conducteur est en excès de vitesse. Le
// modèle de la voiture est aussi enregistré

import "fmt"

func main() {
	var vitesse, vitessemax float64
	var modele string
	fmt.Println("Vitesse mesuré par le radar: ")
	fmt.Scanln(&vitesse)

	fmt.Println("Vitesse max autorisé: ")
	fmt.Scanln(&vitessemax)

	fmt.Println("Quel est le modèle de votre voiture ?")
	fmt.Scanln(&modele)

	if vitesse > vitessemax {
		fmt.Println("Vous etes en excès de vitesse avec votre ", modele)
	}
}

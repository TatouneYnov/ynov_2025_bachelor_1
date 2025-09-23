package main

//un professeur attribue une note à un
//étudiant sur 20. Selon la valeur, le programme doit
//indiquer si la performance correspond à un échec, à
//passable, à bien ou à excellent. La note finale doit
//toujours être affichée entre 0 et 20, si ce n’est pas
//le cas, le programme doit alors afficher une erreur.
//Le programme enregistre aussi la salle de classe où
//l’examen a eu lieu

import "fmt"

func main() {
	var note float64
	var salle string

	fmt.Print("Entrez la salle de classe : ")
	fmt.Scanln(&salle)

	fmt.Print("Entrez la note de l'étudiant (sur 20) : ")
	fmt.Scanln(&note)

	if note < 0 || note > 20 {
		fmt.Println("Erreur : Entrez une note entre 0 et 20")
	} else {
		fmt.Printf("Note finale : %.1f/20\n", note)
		fmt.Println("Salle :", salle)

		if note < 10 {
			fmt.Println("Performance : Échec")
		} else if note >= 10 && note < 12 {
			fmt.Println("Performance : Passable")
		} else if note >= 12 && note < 16 {
			fmt.Println("Performance : Bien")
		} else {
			fmt.Println("Performance : Excellent")
		}
	}
}

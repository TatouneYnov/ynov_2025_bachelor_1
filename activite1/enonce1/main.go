package main

//si la personne est majeur lui donner les document officiel.
//On note l'age de la personne en années.
//afficher si la personne est majeur ou mineur et son prenom

import "fmt"

func main() {
	var age int
	var prenom string

	fmt.Print("Entrez votre prénom : ")
	fmt.Scanln(&prenom)

	fmt.Print("Entrez votre âge : ")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println(prenom, "est majeur, vous recevez vos documents officiel.")
	} else {
		fmt.Println(prenom, "est mineur.")
	}
}

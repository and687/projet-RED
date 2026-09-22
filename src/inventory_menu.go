package src

import (
	"fmt"
)

var items []string

func AfficherInventaire() {
	var choix int

	if len(items) > 0 {
		fmt.Println("Vous avez dans votre inventaire :")
		for i, item := range items {
			fmt.Printf("%d : %s\n", i+1, item)
		}

		fmt.Println("Souhaitez-vous utiliser un objet ? (Tapez un numéro, ou 0 pour retourner au menu principal)")
		fmt.Scan(&choix)

		if choix == 0 {
			AfficherMenu()
		}

		if choix < 1 || choix > len(items) {
			fmt.Println("Choix invalide.")
			AfficherInventaire()
			return
		}

		itemChoisi := items[choix-1]
		UtilisationItem(itemChoisi)

	} else {
		fmt.Println("Votre inventaire est vide")
		AfficherMenu()
	}
}

func UtilisationItem(item string) {
	switch item {
	case "Trousse de soin":
		if Personnage.Health < Personnage.HealthMax {
			fmt.Println("Vous utilisez une trousse de soin. +20 PV")
			Personnage.Health += 20
			if Personnage.Health > Personnage.HealthMax {
				Personnage.Health = Personnage.HealthMax
			}
			SupprimerItems(item)
			AfficherInventaire()
		} else {
			fmt.Print("Vous etes déja au maximim de vos ponts de vie")
			AfficherInventaire()
		}
	// case "Fléchette empoisonnée":
	// 	fmt.Println("Vous utilisez une fléchette empoisonnée.")
	// 	for i := 0; i < 3; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		data.Monster.Health -= 10

	// 		fmt.Println(data.Monster.Health, "/", data.Monster.HealthMax)
	// }

	// SupprimerItems(item)
	// AfficherInventaire()

	default:
		fmt.Println("Cet objet n'a pas d'effet particulier.")
		AfficherInventaire()
	}
}

func SupprimerItems(item string) {
	for i, p := range items {
		if p == item {
			items = append(items[:i], items[i+1:]...)
			fmt.Println("Vous avez utilisé :", item)
			return
		}
	}
	fmt.Println("Objet introuvable dans l'inventaire.")
}

func AfficherEquipement() {

	fmt.Println(Personnage.Stuff.Casque, "PV")
	fmt.Println(Personnage.Stuff.Torse, "PV")
	fmt.Println(Personnage.Stuff.Jambiere, "PV")

}

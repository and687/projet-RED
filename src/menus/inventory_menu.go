package menus

import (
	"fmt"
	"projet-RED/src/models"
)

var items []string

func AfficherInventaire() {
	var choix int

	if len(items) > 0 {
		fmt.Println("Vous avez dans votre inventaire :")
		for i, item := range items {
			fmt.Printf("%d : %s\n", i+1, item)
		}

		fmt.Println("Souhaitez-vous utiliser un objet ? (Tapez un numéro, ou 0 pour annuler)")
		fmt.Scan(&choix)

		if choix == 0 {
			return
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
		fmt.Println("Vous utilisez une trousse de soin. +20 PV")
		models.Personnage.Health += 20
		if models.Personnage.Health > models.Personnage.HealthMax {
			models.Personnage.Health = models.Personnage.HealthMax
		}
		SupprimerItems(item)
		AfficherInventaire()

	case "Fléchette empoisonnée":
		fmt.Println("Vous utilisez une fléchette empoisonnée.")
		SupprimerItems(item)
		AfficherInventaire()
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

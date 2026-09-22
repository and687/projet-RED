package src

import (
	"fmt"
	"time"
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

	case "Fléchette empoisonnée":
		fmt.Println("Vous utilisez une fléchette empoisonnée.")
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ennemi.Health -= 10

			fmt.Println(ennemi.Health, "/", ennemi.HealthMax)
		}

		SupprimerItems(item)
		AfficherInventaire()

	case "Double pistolets":
		ennemi.Health -= 20
		fmt.Println("Arme utilisée : Double pistolets")
		fmt.Println("Dégâts infligés : 20")
		fmt.Println("PV restants de l'adversaire :", ennemi.Health)
		AfficherInventaire()
	case " Fusil d'assault":
		ennemi.Health -= 30
		fmt.Println("Arme utilisée : Fusil d'assault")
		fmt.Println("Dégâts infligés : 30")
		fmt.Println("PV restants de l'adversaire :", ennemi.Health)
		AfficherInventaire()
	case "Dragon Slayer":
		ennemi.Health -= 100
		fmt.Println("Arme utilisée : Dragon Slayer")
		fmt.Println("Dégâts infligés : 100")
		fmt.Println("PV restants de l'adversaire :", ennemi.Health)
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

func AfficherEquipement() {

	fmt.Println(Personnage.Stuff.Casque, "PV")
	fmt.Println(Personnage.Stuff.Torse, "PV")
	fmt.Println(Personnage.Stuff.Jambiere, "PV")

}

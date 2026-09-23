package src

import (
	"fmt"
	"time"
)

var items []string

func AfficherInventaire() {
	var choix int

	if len(items) == 0 {
		fmt.Println("Votre inventaire est vide")
		AfficherMenu()
		return
	}

	fmt.Println("Vous avez dans votre inventaire :")

	for i, item := range items {
		fmt.Printf("%d : %s\n", i+1, item)
	}

	fmt.Println("Souhaitez-vous utiliser un objet ?")
	fmt.Println("Tapez un numéro, ou 0 pour retourner au menu principal")

	fmt.Scan(&choix)

	if choix == 0 {
		AfficherMenu()
		return
	}

	if choix < 1 || choix > len(items) {
		fmt.Println("Choix invalide.")
		AfficherInventaire()
		return
	}

	itemChoisi := items[choix-1]

	UtilisationItem(&Personnage, &ennemi, itemChoisi)
}

func UtilisationItem(personnage *Character, ennemi *Monster, item string) {
	switch item {

	case "Trousse de soin":
		if personnage.Health >= personnage.HealthMax {
			fmt.Println("Vous avez déjà le maximum de points de vie.")
			return
		}

		personnage.Health += 20

		if personnage.Health > personnage.HealthMax {
			personnage.Health = personnage.HealthMax
		}

		fmt.Println("Vous utilisez une trousse de soin. +20 PV")
		SupprimerItems(item)

	case "Fléchette empoisonée":
		fmt.Println("Vous utilisez une fléchette empoisonnée.")

		for i := 0; i < 3 && ennemi.Health > 0; i++ {
			time.Sleep(time.Second)

			ennemi.Health -= 10

			if ennemi.Health < 0 {
				ennemi.Health = 0
			}

			fmt.Println("PV de l'ennemi :", ennemi.Health, "/", ennemi.HealthMax)
		}

		SupprimerItems(item)

	case "Double pistolets":
		ennemi.Health -= 20

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}

		fmt.Println("Dégâts infligés : 20")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	case "Fusil d'assault":
		ennemi.Health -= 30

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}

		fmt.Println("Dégâts infligés : 30")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	case "Dragon Slayer":
		ennemi.Health -= 100

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}

		fmt.Println("Dégâts infligés : 100")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	default:
		fmt.Println("Cet objet n'a pas d'effet particulier.")
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

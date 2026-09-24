package src

import (
	"fmt"
	"time"
)

var items []string

func AfficherInventaire() {
	var choix int

	if len(items) == 0 {
		PrintCouleur(Rouge, "Votre inventaire est vide.")
		AfficherMenu()
		return
	}

	PrintCouleur(Bleu+Gras, "===== VOTRE INVENTAIRE =====")

	for i, item := range items {
		fmt.Printf("%s%d : %s%s\n", Jaune, i+1, item, Reset)
	}

	fmt.Println("Souhaitez-vous utiliser un objet ?")
	fmt.Println("Tapez un numéro, ou 0 pour retourner au menu principal")

	fmt.Scan(&choix)
	clearTerminal()

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

	case "Trousse de soins":
		if personnage.Health >= personnage.HealthMax {
			fmt.Println("Vous avez déjà le maximum de points de vie.")
			return
		}

		personnage.Health += 50

		if personnage.Health > personnage.HealthMax {
			personnage.Health = personnage.HealthMax
		}

		fmt.Println("Vous utilisez une trousse de soins. +50 PV")
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

	fmt.Println("Casque :", Personnage.Stuff.Casque.Nom, "- Défense :", Personnage.Stuff.Casque.Defense)

	fmt.Println("Torse :", Personnage.Stuff.Torse.Nom, "- Défense :", Personnage.Stuff.Torse.Defense)

	fmt.Println("Jambières :", Personnage.Stuff.Jambiere.Nom, "- Défense :", Personnage.Stuff.Jambiere.Defense)

}

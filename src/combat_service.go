package src

import (
	"fmt"
)

func LancerCombat(personnage *Character, ennemi *Monster) bool {
	return CharacterTurn(personnage, ennemi)
}

func CharacterTurn(personnage *Character, ennemi *Monster) bool {
	tour := 1

	for {
		fmt.Println("=== A votre tour ===")
		fmt.Println("1. Retour menu principal")
		fmt.Println("2. Attaquer")
		fmt.Println("3. Inventaire")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Retour au menu principal")
			return true
		case 2:
			degats := 5
			ennemi.Health -= degats
			fmt.Println("Attaque utilisée : Attaque basique")
			fmt.Println("Dégâts infligés :", degats)
			fmt.Println("PV restants de l'adversaire :", ennemi.Health)

		case 3:

			if len(items) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}

			fmt.Println("=== Inventaire ===")

			for i, item := range items {
				fmt.Printf("%d - %s\n", i+1, item)
			}

			fmt.Println("Tapez 0 pour retourner au combat.")

			var choixItem int
			fmt.Scan(&choixItem)

			if choixItem == 0 {
				continue
			}

			if choixItem < 1 || choixItem > len(items) {
				fmt.Println("Choix invalide.")
				continue
			}

			item := items[choixItem-1]

			UtilisationItem(personnage, ennemi, item)

			if ennemi.Health <= 0 {
				ennemi.Health = 0
			}
		}
		if ennemi.Health <= 0 {
			ennemi.Health = 0
			personnage.Money += ennemi.Recompense
			fmt.Println("Le monstre est vaincu !")
			fmt.Println("Tu gagnes", ennemi.Recompense, "d'or !")
			fmt.Println("Tu as maintenant", personnage.Money, "d'or.")
			AddExperience(personnage, ennemi.Experience)

			return false
		}

		degats := GoblinPattern(*ennemi, tour)

		defense := personnage.Stuff.Casque.Defense
		defense += personnage.Stuff.Torse.Defense
		defense += personnage.Stuff.Jambiere.Defense

		degatSubis := degats - defense

		if degatSubis < 0 {
			degatSubis = 0
		}
		personnage.Health -= degatSubis
		if personnage.Health < 0 {
			personnage.Health = 0
		}

		fmt.Println(ennemi.Name, "te frappe pour", degats, "dégâts.")
		fmt.Println("PV du joueur :", personnage.Health)

		tour++

		if IsDead(*personnage) {
			fmt.Println("Tu as perdu !")
			return false
		}
	}
}

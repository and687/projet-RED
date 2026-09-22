package services

import (
	"fmt"
	"projet-RED/src/data"
	"projet-RED/src/models"
)

func LancerCombat(personnage *models.Character, ennemi *data.Monster) {
	CharacterTurn(personnage, ennemi)
}

func CharacterTurn(personnage *models.Character, ennemi *data.Monster) {
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
			return

		case 2:
			degats := 5
			ennemi.Health -= degats
			fmt.Println("Attaque utilisée : Attaque basique")
			fmt.Println("Dégâts infligés :", degats)
			fmt.Println("PV restants de l'adversaire :", ennemi.Health)

		case 3:
			fmt.Println("=== Inventaire ===")
			fmt.Println("1. Double pistolets")
			fmt.Println("2. Fusil d'assault")
			fmt.Println("3. Dragon Slayer")
			fmt.Println("4. Trousse de soin")
			fmt.Println("5. Fléchette empoisonnée")

			var choixObjet int
			fmt.Scan(&choixObjet)

			switch choixObjet {
			case 1:
				ennemi.Health -= 20
				fmt.Println("Arme utilisée : Double pistolets")
				fmt.Println("Dégâts infligés : 20")
				fmt.Println("PV restants de l'adversaire :", ennemi.Health)
			case 2:
				ennemi.Health -= 30
				fmt.Println("Arme utilisée : Fusil d'assault")
				fmt.Println("Dégâts infligés : 30")
				fmt.Println("PV restants de l'adversaire :", ennemi.Health)
			case 3:
				ennemi.Health -= 100
				fmt.Println("Arme utilisée : Dragon Slayer")
				fmt.Println("Dégâts infligés : 100")
				fmt.Println("PV restants de l'adversaire :", ennemi.Health)
			case 4:
				personnage.Health += 20
				if personnage.Health > personnage.HealthMax {
					personnage.Health = personnage.HealthMax
				}
				fmt.Println("Tu utilises une trousse de soin")
				fmt.Println("PV du joueur :", personnage.Health)
			case 5:
				ennemi.Health -= 5
				fmt.Println("Tu utilises une fléchette empoisonnée")
				fmt.Println("PV restants de l'adversaire :", ennemi.Health)
			default:
				fmt.Println("Choix invalide.")
			}
		default:
			fmt.Println("Choix invalide.")
		}

		if ennemi.Health <= 0 {
			ennemi.Health = 0
			personnage.Money += ennemi.Recompense
			fmt.Println("Le monstre est vaincu !")
			fmt.Println("Tu gagnes", ennemi.Recompense, "d'or !")
			fmt.Println("Tu as maintenant", personnage.Money, "d'or.")
			return
		}

		degats := data.GoblinPattern(*ennemi, tour)
		personnage.Health -= degats
		if personnage.Health < 0 {
			personnage.Health = 0
		}

		fmt.Println(ennemi.Name, "te frappe pour", degats, "dégâts.")
		fmt.Println("PV du joueur :", personnage.Health)

		tour++

		if models.IsDead(*personnage) {
			fmt.Println("Tu as perdu !")
			return
		}
	}
}
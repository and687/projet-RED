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
		fmt.Println("Un ", ennemi.Name, " apparait")
		fmt.Println("1. Coup de poings")
		fmt.Println("2. Armes")
		fmt.Println("3. Inventaire")
		fmt.Println("0. Retour menu principal")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			fmt.Println("Retour au menu principal")
			return true
		case 1:
			degats := 15
			ennemi.Health -= degats
			fmt.Println("Attaque utilisée : Attaque coup de poing")
			fmt.Println("Dégâts infligés :", degats)
			fmt.Println("PV restants de l'adversaire :", ennemi.Health)

		case 2:

			if len(weapon) == 0 {
				fmt.Println("Vous n'avez pas d'armes.")
				continue
			}

			fmt.Println("=== Armement ===")

			for i, w := range weapon {
				fmt.Printf("%d - %s\n", i+1, w)
			}

			fmt.Println("Tapez 0 pour retourner au combat.")

			var choixweapon int
			fmt.Scan(&choixweapon)

			if choixweapon == 0 {
				continue
			}

			if choixweapon < 1 || choixweapon > len(weapon) {
				fmt.Println("Choix invalide.")
				continue
			}

			weapon := weapon[choixweapon-1]

			UtilisationWeapon(personnage, ennemi, weapon)

			if ennemi.Health <= 0 {
				ennemi.Health = 0
			}

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

		fmt.Println(ennemi.Name, "te frappe pour", degats, "dégâts.", "armure est de", personnage.Stuff.Casque.Defense+
			personnage.Stuff.Torse.Defense+personnage.Stuff.Jambiere.Defense)
		fmt.Println("le monstre vous a infligé :", degatSubis)
		fmt.Println("PV du joueur :", personnage.Health)

		tour++

		if IsDead(*personnage) {

			if ennemi.Name == "Manequin d'entraînement" {
				fmt.Println("Tu as perdu l'entraînement.")
				fmt.Println("Tes statistiques et ton inventaire sont conservés.")
				return false
			}

			fmt.Println("Tu as perdu !")
			ReinitialiséPersonnage(personnage)
			return false

			fmt.Println("Tu as perdu le combat")
			return false
		}
	}
}

func TrainingFight(personnage *Character) bool {
	healthAvantTraining := personnage.Health

	monstre := InitGoblin(
		"Manequin d'entraînement",
		5,
		50,
		0,
		50,
		0,
	)

	resultat := CharacterTurn(personnage, &monstre)

	personnage.Health = healthAvantTraining

	return resultat
}

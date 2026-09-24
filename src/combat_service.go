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
		PrintCouleur(Rouge+Gras, "=== À VOTRE TOUR ===")
		fmt.Println(Jaune + "Un " + ennemi.Name + " apparaît !" + Reset)

		fmt.Println(Cyan + "1. Coup de poing" + Reset)
		fmt.Println(Bleu + "2. Armes" + Reset)
		fmt.Println(Vert + "3. Inventaire" + Reset)
		fmt.Println(Gris + "0. Retour au menu principal" + Reset)

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			fmt.Println("Retour au menu principal")
			return true
		case 1:
			degats := 15
			ennemi.Health -= degats
			PrintCouleur(Rouge, "Attaque utilisée : coup de poing")
			fmt.Println(Rouge+"Dégâts infligés :"+Reset, degats)
			fmt.Println(Vert+"PV restants :"+Reset, ennemi.Health)
			
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

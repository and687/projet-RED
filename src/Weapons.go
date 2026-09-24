package src

import "fmt"

var weapon []string

func weapons() {
	var choix int

	if len(weapon) == 0 {
		fmt.Println("Vous n'avez pas d'armes")
		AfficherMenu()
		return

		fmt.Println("Vous avez dans votre armement :")

		for i, w := range weapon {
			fmt.Printf("%d : %s\n", i+1, w)
		}

		fmt.Println("Souhaitez-vous utiliser une armes ?")
		fmt.Println("Tapez un numéro, ou 0 pour retourner au menu principal")

		fmt.Scan(&choix)
		clearTerminal()

		if choix == 0 {
			AfficherMenu()
			return
		}

		if choix < 1 || choix > len(weapon) {
			fmt.Println("Choix invalide.")
			AfficherInventaire()
			return
		}

		weaponChoisi := weapon[choix-1]

		UtilisationWeapon(&Personnage, &ennemi, weaponChoisi)
	}
}

func UtilisationWeapon(personnage *Character, ennemi *Monster, weapon string) {

	switch weapon {

	case "Double pistolets":
		ennemi.Health -= 20

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}
		fmt.Println("Vous utilisez les double pistolets")
		fmt.Println("Dégâts infligés : 20")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	case "Fusil d'assault":
		ennemi.Health -= 30

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}
		fmt.Println("Vous utilisez le fusil d'assault")
		fmt.Println("Dégâts infligés : 30")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	case "Dragon Slayer":
		ennemi.Health -= 100

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}
		fmt.Println("Vous utilisez la Dragon Slayer")
		fmt.Println("Dégâts infligés : 100")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)

	case "Épée à energie":
		ennemi.Health -= 200

		if ennemi.Health < 0 {
			ennemi.Health = 0
		}
		fmt.Println("Vous utilisez l'épée à énergie")
		fmt.Println("Dégâts infligés : 200")
		fmt.Println(ennemi.Health, "/", ennemi.HealthMax)
	}
}

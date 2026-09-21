package services

import ( "fmt"
	"projet-RED/src/data"
    "projet-RED/src/models")

	func characterTurn(personnage *models.Character, ennemi *data.Monster) {
		fmt.Println("=== a votre tour===")
		fmt.Println("1. Menu")
		fmt.Println("2. Attaquer")
		fmt.Println("3. Inventaire")

		var choix int 
		fmt.Scan(&choix)

		switch choix {
		case 1:
			AfficherMenu()

		case 2:
			attaque := "attaque basique"
			degats := 5

			ennemi.Health -=degats

			fmt.Println("Attaque utilisée :", attaque)
			fmt.Println("Degats infligés :", degats)
			fmt.Println("PV restants de l'adversaire :", ennemi.Health)

		case 3:
			fmt.Println("===Inventaire===")
			for i, item :=range items{
				fmt.Println(i+1, "-", item)
			}

			var choixArme int 
			fmt.Scan(&choixArme)

			arme = items[choixArme-1]

			switch arme {
			case "Doublbe pistolets":
				ennemi.Health -=10
				fmt.Println("Arme utilisée : Double pistolets")
				fmt.Println("degats infligés : 10")

		case "Fusil d'assault":
    ennemi.Health -= 20
    fmt.Println("Arme utilisée : Fusil d'assault")
    fmt.Println("Dégâts infligés : 20")

case "Dragon Slayer":
    ennemi.Health -= 50
    fmt.Println("Arme utilisée : Dragon Slayer")
    fmt.Println("Dégâts infligés : 50")
	
	case "Trousse de soin"
personnage.Health +=20
fmt.Println("Tu utilises une trousse de soin")

case "Fléchette empoisonnée"
ennemi.Health -= 10
fmt.Println("Tu utilises une fléchette empoisonnée")

}
			}
		}

		if ennemi.Health > 0 {
			personnage.Health -= ennemi.Degats
			fmt.Println(ennemi.Name, "te frappe pour", ennemi.Degats, "degats.")
			mt.Println("PV du joueur :", personnage.Health)
} else {
    fmt.Println("Le monstre est mort.")
}

if models.IsDead(*personnage) {
	fmt.Println("you loose !")
	fmt.Println("1. quitter le jeu")
	fmt.Println("2. R")

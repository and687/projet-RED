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
			case "DoublePistolets":
				ennemi.Health -=10
				fmt.Println("Arme utilisée : DoublePistolets")
				fmt.Println("degats infligés : 10")

package src

import (
	"fmt"
	"os"
)

func Quitter() {
	fmt.Println("merci d'avoir joué au jeu")
	os.Exit(0)
}

func Start() {
	fmt.Println("bienvenue dans le jeu !")
	fmt.Println("Choisissez la classe :")
	fmt.Println("1. Chercheur - PV : 100/200, -Money : 150")
	fmt.Println("2. Pilote -PV : 80/160, -Money : 250")
	fmt.Println("3. Mercenaire -PV : 150/300, -Money : 50")
	étage2 = false
	étage3 = false
	étage4 = false

	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		nom := CharacterCreation()
		Personnage = *InitCharacter(nom, "chercheur", 100, 1, 200, 150)
		items = append(items, "Trousse de soins", "Trousse de soins", "Trousse de soins")
		DisplayInfo(&Personnage)
	case 2:
		nom := CharacterCreation()
		Personnage = *InitCharacter(nom, "pilote", 80, 1, 160, 250)
		items = append(items, "Trousse de soins", "Trousse de soins", "Trousse de soins")
		DisplayInfo(&Personnage)
	case 3:
		nom := CharacterCreation()
		Personnage = *InitCharacter(nom, "mercenaire", 150, 1, 300, 50)
		items = append(items, "Trousse de soins", "Trousse de soins", "Trousse de soins")
		DisplayInfo(&Personnage)
	default:
		fmt.Println("Choix invalide")
		Start()
	}
}

func AfficherMenu() {
	var choix int

	fmt.Println(" Menu principal\n", "1 : Combat\n", "2 : Marchand\n", "3 : Tailleur\n", "4 : Inventaire\n", "5 : Equipement\n", "9 : Quitter le jeu")
	fmt.Printf("Vous avez %d d'or !\n", Personnage.Money)
	fmt.Printf("Vous avez %d / %d Pv\n", Personnage.Health, Personnage.HealthMax)
	fmt.Scan(&choix)

	switch choix {

	case 1:
		MenuCombat()

	case 2:
		Acheter()

	case 3:
		Tailleur()

	case 4:
		AfficherInventaire()

	case 5:
		AfficherEquipement()

	case 9:
		Quitter()

	default:
		println("Choix invalide")
		AfficherMenu()
	}
}

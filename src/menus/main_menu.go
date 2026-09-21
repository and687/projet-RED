package menus

import (
	"fmt"
	"projet-RED/src/models"
)

func Start() {
	fmt.Println("bienvenue dans le jeu !")
	fmt.Println("Choisissez la classe :")
	fmt.Println("1. Chercheur - PV : 100/200, -Money : 150")
	fmt.Println("2. Pilote -PV : 80/160, -Money : 250")
	fmt.Println("3. Mercenaire -PV : 150/300, -Money : 50")

	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		nom := models.CharacterCreation()
		models.Personnage = *models.InitCharacter(nom, "chercheur", 100, 1, 200, 150)
		models.DisplayInfo(&models.Personnage)
	case 2:
		nom := models.CharacterCreation()
		models.Personnage = *models.InitCharacter(nom, "pilote", 80, 1, 160, 250)
		models.DisplayInfo(&models.Personnage)
	case 3:
		nom := models.CharacterCreation()
		models.Personnage = *models.InitCharacter(nom, "mercenaire", 150, 1, 300, 50)
		models.DisplayInfo(&models.Personnage)
	default:
		fmt.Println("Choix invalide")
		Start()
	}
}

func AfficherMenu() {
	var choix int

	fmt.Println(" Menu principal\n", "1 : Combat\n", "2 : Marchand\n", "3 : Forgeron\n", "4 : Inventaire\n", "9 : Quitter le jeu")
	fmt.Printf("Vous avez %d d'or !\n", models.Personnage.Money)
	fmt.Printf("Vous avez %d / %d Pv\n", models.Personnage.Health, models.Personnage.HealthMax)
	fmt.Scan(&choix)

	switch choix {

	case 1:
	// MenuCombat()

	case 2:
		Acheter()

	// case 3:
	// 	Forgeron()

	case 4:
		AfficherInventaire()

	// case 9:
	// 	Quitter()

	default:
		println("Choix invalide")
		AfficherMenu()
	}

}

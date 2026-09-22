package menus

import (
	"fmt"
	"projet-RED/src/data"
	"projet-RED/src/models"
	"projet-RED/src/services"
)

func MenuCombat() {
	fmt.Println("1. Squelette d'entraînement")
	fmt.Println("2. Squelette")
	fmt.Println("3. Zombie")
	fmt.Println("4. Vampire")
	fmt.Println("5. Dragon légendaire")
	fmt.Println("6. Quitter les combats")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		monster := data.InitGoblin("Squelette d'entrainement", 5, 50, 10, 50)
		services.LancerCombat(&models.Personnage, &monster)
	case 2:
		monster := data.InitGoblin("Squelette", 5, 50, 10, 50)
		services.LancerCombat(&models.Personnage, &monster)
	case 3:
		monster := data.InitGoblin("Zombie", 10, 100, 20, 100)
		services.LancerCombat(&models.Personnage, &monster)
	case 4:
		monster := data.InitGoblin("Vampire", 30, 200, 50, 200)
		services.LancerCombat(&models.Personnage, &monster)
	case 5:
		monster := data.InitGoblin("Dragon Legendaire", 100, 500, 2000, 500)
		services.LancerCombat(&models.Personnage, &monster)
	case 6:
		fmt.Println("Tu quittes les combats.")
		return
	default:
		fmt.Println("Choix invalide")
	}
}
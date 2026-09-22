package src

import (
	"fmt"
    "math/rand"
)

func MenuCombat() {
	fmt.Println("1. Squelette d'entraînement")
	fmt.Println("2. Squelette")
	fmt.Println("3. Zombie")
	fmt.Println("4. Vampire")
	fmt.Println("5. Dragon légendaire")
	fmt.Println("6. Quitter les combats")
    fmt.Println("7. la salade de monstre")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		monster := InitGoblin("Squelette d'entrainement", 5, 50, 10, 50, 5)
		LancerCombat(&Personnage, &monster)
	case 2:
		monster := data.InitGoblin("Squelette", 5, 50, 10, 50, 100)
		LancerCombat(&Personnage, &monster)
	case 3:
		monster := InitGoblin("Zombie", 10, 100, 20, 100, 20)
		LancerCombat(&Personnage, &monster)
	case 4:
		monster := InitGoblin("Vampire", 30, 200, 50, 200, 30)
		LancerCombat(&Personnage, &monster)
	case 5:
		monster := InitGoblin("Dragon Legendaire", 100, 500, 2000, 500, 50)
		LancerCombat(&Personnage, &monster)
	case 6:
		fmt.Println("Tu quittes les combats.")
		return
	default:
		fmt.Println("Choix invalide")

    case 7: 
    fmt.Println("=== bienvenue dans la salade de monstres ===")
    SaladeDeMonstres()
	}
}

func SaladeDeMonstres() {
	monstres := []Monster{
		InitGoblin("Squelette", 5, 50, 10, 50, 10),
		InitGoblin("Zombie", 10, 100, 20, 100, 20),
		InitGoblin("Vampire", 30, 200, 50, 200, 30),
		InitGoblin("Dragon légendaire", 100, 500, 2000, 500, 50),
    }

	for !IsDead(Personnage) {
	indice := rand.Intn(len(monstres))
	monstre := monstres[indice]

		fmt.Println("Nouveau monstre :", monstre.Name)
		quitter := LancerCombat(&Personnage, &monstre)

		if quitter {
			fmt.Println("Tu quittes la salade de monstres.")
			return
		}

		if IsDead(Personnage) {
			fmt.Println("La salade de monstres est terminée.")
			return
		}

		fmt.Println("Le prochain monstre arrive!!!")
	}
}
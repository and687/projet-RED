package src

import (
	"fmt"
	"math/rand"
)

var étage2 = false
var étage3 = false
var étage4 = false
var étage5 = false

func MenuCombat() {
	fmt.Println("1. Manequin d'entraînement")
	fmt.Println("2. Premier étage")

	if étage2 {
		fmt.Println("3. Second étage")
	}
	if étage3 {
		fmt.Println("4. Troisieme étage")
	}
	if étage4 {
		fmt.Println("5. Quatrieme étage")
	}
	if étage5 {
		fmt.Println("7. la salade de monstre")
	}
	fmt.Println("6. Quitter les combats")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		TrainingFight(&Personnage)
	case 2:
		monster := InitGoblin("Squelette", 8, 55, 35, 55, 60)
		LancerCombat(&Personnage, &monster)
		if monster.Health <= 0 {
			étage2 = true
			fmt.Println("Le deuxième étage est débloqué !")
		}
	case 3:
		if !étage2 {
			fmt.Println("Vous n'avez pas accès à cet étage")
			return
		}
		monster := InitGoblin("Zombie", 12, 95, 60, 95, 90)
		LancerCombat(&Personnage, &monster)

		if monster.Health <= 0 {
			étage3 = true
			fmt.Println("Le troisième étage est débloqué !")
		}
	case 4:
		if !étage3 {
			fmt.Println("Vous n'avez pas accès à cet étage")
			return
		}
		monster := InitGoblin("Vampire", 18, 150, 110, 150, 120)
		LancerCombat(&Personnage, &monster)

		if monster.Health <= 0 {
			étage4 = true
			fmt.Println("Le quatrième étage est débloqué !")
		}
	case 5:
		if !étage4 {
			fmt.Println("Vous n'avez pas accès à cet étage")
			return
		}
		monster := InitGoblin("Dragon Legendaire", 28, 350, 300, 350, 200)
		LancerCombat(&Personnage, &monster)
		if monster.Health <= 0 {
			étage5 = true
			fmt.Println("La salade de monstre est débloqué !")
		}

	case 6:
		fmt.Println("Tu quittes les combats.")
		return

	default:
		fmt.Println("Choix invalide")

	case 7:
		if !étage5 {
			fmt.Println("Vous n'avez pas accès à cet étage")
			return
		}
		fmt.Println("=== bienvenue dans la salade de monstres ===")
		SaladeDeMonstres()
	}
}

func SaladeDeMonstres() {
	monstres := []Monster{
		InitGoblin("Squelette", 8, 55, 35, 55, 60),
		InitGoblin("Zombie", 12, 95, 60, 95, 90),
		InitGoblin("Vampire", 18, 150, 110, 150, 120),
		InitGoblin("Dragon légendaire", 8, 350, 300, 350, 200),
	}

	for !IsDead(Personnage) {
		indice := rand.Intn(len(monstres))
		monstre := monstres[indice]

		fmt.Println("Nouveau monstre :")
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

package menus

import (
	"fmt"
	"projet-RED/src/models"
)

var Potion bool = true
var Flechettes bool = true
var DoublePistolets bool = true
var FusilAssault bool = true
var DragonSlayer bool = true
var Gratuit bool = true

func Acheter() {
	for {
		fmt.Println("Bienvenu dans mon magasin. Que voulez-vous acheter ?")

		if Potion {
			fmt.Println("1 : Trousse de soin (50g)")
		}
		if Flechettes {
			fmt.Println("2 : Flechette Empoisonée (50g)")
		}
		if DoublePistolets {
			fmt.Println("3 : Double pistolets (200g)")
		}
		if FusilAssault {
			fmt.Println("4 : Fusil d'assault (400g)")
		}
		if DragonSlayer {
			fmt.Println("5 : Épée: The dragon slayer (1000g)")
		}
		if Gratuit {
			fmt.Println("6 : Trousse de soin (Gratuit)")
		}
		fmt.Println("0 : Quitter le magasin")

		var choix int
		fmt.Scan(&choix)

		switch {

		case choix == 1 && Potion:
			if models.Personnage.Money >= 50 {
				models.Personnage.Money -= 50
				items = append(items, "Trousse de soin")
				fmt.Println("Trousse de soin ajoutée à votre inventaire !")
			} else {
				fmt.Println("Vous n'avez pas assez d'or !")
			}

		case choix == 2 && Flechettes:
			if models.Personnage.Money >= 50 {
				models.Personnage.Money -= 50
				items = append(items, "Flechette empoisonée")
				fmt.Println("Flechette empoisonée ajoutée à votre inventaire !")
			} else {
				fmt.Println("Vous n'avez pas assez d'or !")
			}

		case choix == 3 && DoublePistolets:
			if models.Personnage.Money >= 200 {
				models.Personnage.Money -= 200
				items = append(items, "Double pistolets")
				fmt.Println("Épée en fer ajoutée à votre inventaire !")
				DoublePistolets = false
			} else {
				fmt.Println("Vous n'avez pas assez d'or !")
			}

		case choix == 4 && FusilAssault:
			if models.Personnage.Money >= 400 {
				models.Personnage.Money -= 400
				items = append(items, "Fusil d'assault")
				fmt.Println("Fusil d'assault ajoutée à votre inventaire !")
				FusilAssault = false
			} else {
				fmt.Println("Vous n'avez pas assez d'or !")
			}

		case choix == 5 && DragonSlayer:
			if models.Personnage.Money >= 1000 {
				models.Personnage.Money -= 1000
				items = append(items, "Dragon Slayer")
				fmt.Println("Épée: The Dragon Slayer !!!!! Ajoutée à votre inventaire !")
				DragonSlayer = false
			} else {
				fmt.Println("Vous n'avez pas assez d'or !")
			}

		case choix == 6 && Potion:
			items = append(items, "Trousse de soin")
			fmt.Println("Trousse de soin ajoutée à votre inventaire !")
			Gratuit = false

		case choix == 0:
			AfficherMenu()
				return

		default:
			fmt.Println("Choix invalide ou article indisponible.")
		}

	}
}

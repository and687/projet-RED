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
var stockage bool = true
var capacité = 10

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
		if stockage {
			fmt.Println("6 : Agrandir sont inventaire de +5 (300)")
		}
		if Gratuit {
			fmt.Println("7 : Trousse de soin (Gratuit)")
		}
		fmt.Println("0 : Quitter le magasin")

		var choix int
		fmt.Scan(&choix)

		switch {

		case choix == 1 && Potion:
			if len(items) < capacité {
				if models.Personnage.Money >= 0 {
					models.Personnage.Money -= 0
					items = append(items, "Trousse de soin")
					fmt.Println("Trousse de soin ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !", models.Personnage.Money)
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 2 && Flechettes:
			if len(items) < capacité {
				if models.Personnage.Money >= 50 {
					models.Personnage.Money -= 50
					items = append(items, "Flechette empoisonée")
					fmt.Println("Flechette empoisonée ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 3 && DoublePistolets:
			if len(items) < capacité {
				if models.Personnage.Money >= 200 {
					models.Personnage.Money -= 200
					items = append(items, "Double pistolets")
					fmt.Println("Épée en fer ajoutée à votre inventaire !")
					DoublePistolets = false
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}
		case choix == 4 && FusilAssault:
			if len(items) < capacité {
				if models.Personnage.Money >= 400 {
					models.Personnage.Money -= 400
					items = append(items, "Fusil d'assault")
					fmt.Println("Fusil d'assault ajoutée à votre inventaire !")
					FusilAssault = false
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 5 && DragonSlayer:
			if len(items) < capacité {
				if models.Personnage.Money >= 1000 {
					models.Personnage.Money -= 1000
					items = append(items, "Dragon Slayer")
					fmt.Println("Épée: The Dragon Slayer !!!!! Ajoutée à votre inventaire !")
					DragonSlayer = false
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 6 && stockage:
			if models.Personnage.Money >= 0 {
				models.Personnage.Money -= 0
				capacité += 5
				fmt.Println("Vous avez augmenté votre ivnentaire de 5 emplacements")
				stockage = false
			}

		case choix == 7 && Potion:
			if len(items) < capacité {
				items = append(items, "Trousse de soin")
				fmt.Println("Trousse de soin ajoutée à votre inventaire !")
				Gratuit = false
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 0:
			AfficherMenu()
			return

		default:
			fmt.Println("Choix invalide ou article indisponible.")
		}

	}
}

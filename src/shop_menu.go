package src

import (
	"fmt"
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
		fmt.Println("1 : Trousse de soins (50g)")
		fmt.Println("2 : Fléchette Empoisonée (60g)")
		if DoublePistolets {
			fmt.Println("3 : Double pistolets (180g)")
		}
		if FusilAssault {
			fmt.Println("4 : Fusil d'assault (350g)")
		}
		if DragonSlayer {
			fmt.Println("5 : Épée: The dragon slayer (800g)")
		}
		if stockage {
			fmt.Println("6 : Agrandir sont inventaire de +5 (250)")
		}
		if Gratuit {
			fmt.Println("7 : Trousse de soin (Gratuit)")
		}
		fmt.Println("8 : Peau de serpent (50g)")
		fmt.Println("9 : Idole d'or (70g)")
		fmt.Println("10 : Moustache (60g)")
		fmt.Println("11 : Casquette baseball (60g)")
		fmt.Println("12 : Dent de T-Rex (100g)")
		fmt.Println("13 : Tesseract (120g)")
		fmt.Println("0 : Quitter le magasin")

		var choix int
		fmt.Scan(&choix)

		switch {

		case choix == 1:
			if len(items) < capacité {
				if Personnage.Money >= 50 {
					Personnage.Money -= 50
					items = append(items, "Trousse de soins")
					fmt.Println("Trousse de soins ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !", Personnage.Money)
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 2 && Flechettes:
			if len(items) < capacité {
				if Personnage.Money >= 60 {
					Personnage.Money -= 60
					items = append(items, "Fléchette empoisonée")
					fmt.Println("Fléchette empoisonée ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 3 && DoublePistolets:
			if len(items) < capacité {
				if Personnage.Money >= 180 {
					Personnage.Money -= 180
					items = append(items, "Double pistolets")
					fmt.Println("Double pistolets ajoutée à votre inventaire !")
					DoublePistolets = false
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}
		case choix == 4 && FusilAssault:
			if len(items) < capacité {
				if Personnage.Money >= 350 {
					Personnage.Money -= 350
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
				if Personnage.Money >= 800 {
					Personnage.Money -= 800
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
			if Personnage.Money >= 250 {
				Personnage.Money -= 250
				if capacité < 30 {
					capacité += 10
				} else {
					stockage = false
				}
				fmt.Println("Vous avez augmenté votre ivnentaire de 10 emplacements")
			}

		case choix == 7:
			if len(items) < capacité {
				items = append(items, "Trousse de soin")
				fmt.Println("Trousse de soin ajoutée à votre inventaire !")
				Gratuit = false
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}
		case choix == 8:
			if len(items) < capacité {
				if Personnage.Money >= 50 {
					Personnage.Money -= 50
					items = append(items, "Peau de serpent")
					fmt.Println("Peau de serpent ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !", Personnage.Money)
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 9:
			if len(items) < capacité {
				if Personnage.Money >= 70 {
					Personnage.Money -= 70
					items = append(items, "Idole d'or")
					fmt.Println("Idole d'or ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 10:
			if len(items) < capacité {
				if Personnage.Money >= 60 {
					Personnage.Money -= 60
					items = append(items, "Moustache")
					fmt.Println("Moustache ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}
		case choix == 11:
			if len(items) < capacité {
				if Personnage.Money >= 60 {
					Personnage.Money -= 60
					items = append(items, "Casquette baseball")
					fmt.Println("Casquette baseball ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 12:
			if len(items) < capacité {
				if Personnage.Money >= 100 {
					Personnage.Money -= 100
					items = append(items, "dent de T-Rex")
					fmt.Println("dent de T-Rex Ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 13:
			if len(items) < capacité {
				if Personnage.Money >= 120 {
					Personnage.Money -= 120
					items = append(items, "Tesseract")
					fmt.Println("Tesseract Ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !")
				}
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

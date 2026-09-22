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
		fmt.Println("1 : Trousse de soin (50g)")
		fmt.Println("2 : Flechette Empoisonée (50g)")
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
		fmt.Println("8 : Peau de serpent (70g)")
		fmt.Println("9 : Idole d'or (70g)")
		fmt.Println("10 : Moustache (70g)")
		fmt.Println("11 : Casquette baseball (70g)")
		fmt.Println("12 : Dent de T-Rex (70g)")
		fmt.Println("13 : Tesseract (70g)")
		fmt.Println("0 : Quitter le magasin")

		var choix int
		fmt.Scan(&choix)

		switch {

		case choix == 1:
			if len(items) < capacité {
				if Personnage.Money >= 0 {
					Personnage.Money -= 0
					items = append(items, "Trousse de soin")
					fmt.Println("Trousse de soin ajoutée à votre inventaire !")
				} else {
					fmt.Println("Vous n'avez pas assez d'or !", Personnage.Money)
				}
			} else {
				fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			}

		case choix == 2 && Flechettes:
			if len(items) < capacité {
				if Personnage.Money >= 50 {
					Personnage.Money -= 50
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
				if Personnage.Money >= 200 {
					Personnage.Money -= 200
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
				if Personnage.Money >= 400 {
					Personnage.Money -= 400
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
				if Personnage.Money >= 1000 {
					Personnage.Money -= 1000
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
			if Personnage.Money >= 0 {
				Personnage.Money -= 0
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
				if Personnage.Money >= 0 {
					Personnage.Money -= 0
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
				if Personnage.Money >= 0 {
					Personnage.Money -= 0
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
				if Personnage.Money >= 200 {
					Personnage.Money -= 200
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
				if Personnage.Money >= 400 {
					Personnage.Money -= 400
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
				if Personnage.Money >= 1000 {
					Personnage.Money -= 1000
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
				if Personnage.Money >= 1000 {
					Personnage.Money -= 1000
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

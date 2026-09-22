package menus

import (
	"fmt"
	"projet-RED/src/models"
)

func Tailleur() {

	var approuved string
	var choix int

	fmt.Println("bonjour, que souhaitez que je vous prépare :")
	fmt.Println("1. Chapeau d'aventurier (protege des coups de fouet)")
	fmt.Println("2. Chemise à fleur ( parfait pour conduire des ferrari)")
	fmt.Println("3. Micro short (comme celui d'une grande exploratrice)")
	fmt.Scan(&choix)

	switch {

	case choix == 1:
		fmt.Println("Voulez-vous le chapeau d'aventurier ? (oui/non)")
		fmt.Println("- 1 peau de serpent")
		fmt.Println("- 1 idole d'or")
		fmt.Println("- 250 gold")
		fmt.Scan(&approuved)

		if approuved == "non" {
			fmt.Println("Très bien, la chemise à fleur ne sera pas fabriqué.")
			return
		}
		if approuved == "oui" {

			Peau := false
			Idole := false

			for _, item := range items {
				if item == "Peau de serpent" {
					Peau = true
				}
				if item == "Idole d'or" {
					Idole = true
				}
			}

			if len(items) < capacité && Peau && Idole && models.Personnage.Money >= 250 {
				models.Personnage.Money -= 250

				models.Personnage.Stuff.Casque = models.Objet{
					Nom:     "Chapeau d'aventurier",
					Defense: 10,
				}
				SupprimerItems("Peau de serpent")
				SupprimerItems("Idole d'or")

				fmt.Println("Bravo ! Vous avez fabriqué le chapeau d'aventurier !")
			} else {
				fmt.Println("Impossible de fabriquer l'objet (manque d'objets, d'or, ou inventaire plein).")
			}

		}
	}
	if choix == 2 {
		fmt.Println("Voulez-vous la chemise à fleur ? (oui/non)")
		fmt.Println("- 1 une moustache")
		fmt.Println("- 1 casquette de baseball")
		fmt.Println("- 250 gold")
		fmt.Scan(&approuved)

		if approuved == "non" {
			fmt.Println("Très bien, la chemise à fleur ne sera pas fabriqué.")
			return
		}

		if approuved == "oui" {

			Moustache := false
			Casquette := false

			for _, item := range items {
				if item == "Moustache" {
					Moustache = true
				}
				if item == "Masquette baseball" {
					Casquette = true
				}
			}

			if len(items) < capacité && Moustache && Casquette && models.Personnage.Money >= 250 {
				models.Personnage.Money -= 250

				models.Personnage.Stuff.Torse = models.Objet{
					Nom:     "Chemise à fleur",
					Defense: 50,
				}
				SupprimerItems("Moustache")
				SupprimerItems("Casquette")

				fmt.Println("Bravo ! Vous avez fabriqué la chemise à fleur !")
			} else {
				fmt.Println("Impossible de fabriquer l'objet (manque d'objets, d'or, ou inventaire plein).")
			}
		}
	}
	if choix == 3 {
		fmt.Println("Voulez-vous le micro short ? (oui/non)")
		fmt.Println("- 1 Dent de t-Rex")
		fmt.Println("- 1 Tesseract ")
		fmt.Println("- 250 gold")
		fmt.Scan(&approuved)

		if approuved == "non" {
			fmt.Println("Très bien, le micro short ne sera pas fabriqué.")
			return
		}

		if approuved == "oui" {

			Dent := false
			Tesseract := false

			for _, item := range items {
				if item == "dent de T-Rex" {
					Dent = true
				}
				if item == "Tesseract" {
					Tesseract = true
				}
			}

			if len(items) < capacité && Dent && Tesseract && models.Personnage.Money >= 250 {
				models.Personnage.Money -= 250

				models.Personnage.Stuff.Jambiere = models.Objet{
					Nom:     "Micro short",
					Defense: 30,
				}
				SupprimerItems("dent de T-Rex")
				SupprimerItems("Tesseract")

				fmt.Println("Bravo ! Vous avez fabriqué la chemise à fleur !")
			} else {
				fmt.Println("Impossible de fabriquer l'objet (manque d'objets, d'or, ou inventaire plein).")
			}
		}
	}
}

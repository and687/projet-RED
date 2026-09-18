package main_menu

import (
    "fmt"
    character "projet-RED/src/models"
)

func main() {
	fmt.Println("bienvenue dans le jeu !")
	fmt.Println("Choisissez la classe :")
    fmt.Println("1. Chercheur - PV : 100/200")
    fmt.Println("2. Pilote -PV : 80/160")
    fmt.Println("3. Mercenaire -PV : 150/300")

	var choix int
	fmt.Scan(&choix)

	switch choix {
case 1:
    nom := character.CharacterCreation()
    personnage := character.IntCharacter(nom, "chercheur", 100, 1, 200)
    character.DisplayInfo(personnage)
case 2:
    nom := character.CharacterCreation()
    personnage := character.IntCharacter(nom, "pilote", 80, 1, 160)
    character.DisplayInfo(personnage)
case 3:
    nom := character.CharacterCreation()
    personnage := character.IntCharacter(nom, "mercenaire", 150, 1, 300)
    character.DisplayInfo(personnage)
default:
    fmt.Println("Choix invalide")
}
}


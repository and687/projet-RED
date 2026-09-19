package Inventory

import (
	"fmt"
	main_menu "projet-RED/src/menus"
)

var items []string

func AfficherInventaire() {

	if len(items) > 0 {
		fmt.Println("vous avez dans votre inventaire :\n", items)
	} else {
		fmt.Println("Votre inventaire est vide")
		main_menu.AfficherMenu()
	}

}

func AjoutInventaire() {

	var choix int

	switch choix {

	// case 1:
	// 	Character.Money - 30
	// 	items = append(items, "potions de soin")
	// 	fmt.Println("vous avez dépensez 30g, ajout d'une potion de soins")

	// case 2:
	// 	Character.Money - 30
	// 	items = append(items, "potions de poison")
	// 	fmt.Println("vous avez dépensez 30g, ajout d'une potion de poison")

	// case 3:
	// 	Character.Money - 30
	// 	items = append(items, "potions de soin")
	// 	fmt.Println("vous avez dépensez 30g, ajout d'une potion de soins")

	// case 4:
	// 	Character.Money - 30
	// 	items = append(items, "potions de soin")
	// 	fmt.Println("vous avez dépensez 30g, ajout d'une potion de soins")
	}

}

func SupprimerItems() {

	var suppr int

	switch suppr {

	case 1:
		for i, p := range items {
			if p == "potion de soin" {
				items = append(items[:i], items[i+1:]...)
				fmt.Println("vous avez consomé une potion de soins")
				break
			}
		}

	case 2:
		for i, p := range items {
			if p == "potion de poison" {
				items = append(items[:i], items[i+1:]...)
				fmt.Println("vous avez consomé une potion de soins")
				break
			}
		}
	}

}

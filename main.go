package main

import (
	"fmt"
	"projet-RED/src"
)

func main() {
	for {
		src.Start()

		for {
			src.AfficherMenu()

			if src.IsDead(src.Personnage) {
				fmt.Println("Tu es mort !")
				fmt.Println("1. Recommencer")
				fmt.Println("2. Quitter")

				var choixMort int
				fmt.Scan(&choixMort)

				switch choixMort {
				case 1:
					break
				case 2:
					fmt.Println("Merci d'avoir joué au jeu")
					return
				default:
					fmt.Println("Choix invalide")
				}

				break
			}
		}
	}
}

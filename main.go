package main

import ("projet-RED/src/menus"
"projet-RED/src/models"
"fmt")

func main() {
    for {
        menus.Start()

        for {
            menus.AfficherMenu()

            if models.IsDead(models.Personnage) {
                fmt.Println("Tu es mort !")
                fmt.Println("1. Recommencer")
                fmt.Println("2. Quitter")

                var choixMort int
                fmt.Scan(&choixMort)

                switch choixMort {
                case 1:
                    break
                case 2:
                    return
                default:
                    fmt.Println("Choix invalide")
                }

                break
            }
        }
    }
}

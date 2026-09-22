package src

import (
    "fmt"
    "time"
)
 
func LancerCombat(personnage *models.Character, ennemi *data.Monster) bool {
    return CharacterTurn(personnage, ennemi)
}
 
func CharacterTurn(personnage *models.Character, ennemi *data.Monster) bool {
    tour := 1
 
    for {
        fmt.Println("=== A votre tour ===")
        fmt.Println("1. Retour menu principal")
        fmt.Println("2. Attaquer")
        fmt.Println("3. Inventaire")
 
        var choix int
        fmt.Scan(&choix)
 
        switch choix {
        case 1:
            fmt.Println("Retour au menu principal")
            return true
        case 2:
            degats := 5
            ennemi.Health -= degats
            fmt.Println("Attaque utilisée : Attaque basique")
            fmt.Println("Dégâts infligés :", degats)
            fmt.Println("PV restants de l'adversaire :", ennemi.Health)
 
        case 3:
 
            if len(inventaire.Items) == 0 {
                fmt.Println("Votre inventaire est vide.")
                continue
            }
 
            fmt.Println("=== Inventaire ===")
 
            for i, item := range inventaire.Items {
                fmt.Println(i+1, "-", item)
            }
 
            fmt.Print("Quel objet voulez-vous utiliser ? ")
 
            var choixItem int
            fmt.Scan(&choixItem)
 
            if choixItem < 1 || choixItem > len(inventaire.Items) {
                fmt.Println("Choix invalide.")
                continue
            }
 
            item := inventaire.Items[choixItem-1]
 
            switch item {
            case "Trousse de soin":
                if Personnage.Health < Personnage.HealthMax {
                    fmt.Println("Vous utilisez une trousse de soin. +20 PV")
                    Personnage.Health += 20
                    if Personnage.Health > Personnage.HealthMax {
                        Personnage.Health = Personnage.HealthMax
                    }
                    SupprimerItems(item)
                    AfficherInventaire()
                } else {
                    fmt.Print("Vous etes déja au maximim de vos ponts de vie")
                    AfficherInventaire()
                }
 
            case "Fléchette empoisonnée":
                fmt.Println("Vous utilisez une fléchette empoisonnée.")
                for i := 0; i < 3; i++ {
                    time.Sleep(1 * time.Second)
                    ennemi.Health -= 10
 
                    fmt.Println(ennemi.Health, "/", ennemi.HealthMax)
                }
 
                SupprimerItems(item)
                AfficherInventaire()
 
            case "Double pistolets":
                ennemi.Health -= 20
                fmt.Println("Arme utilisée : Double pistolets")
                fmt.Println("Dégâts infligés : 20")
                fmt.Println("PV restants de l'adversaire :", ennemi.Health)
                AfficherInventaire()
            case " Fusil d'assault":
                ennemi.Health -= 30
                fmt.Println("Arme utilisée : Fusil d'assault")
                fmt.Println("Dégâts infligés : 30")
                fmt.Println("PV restants de l'adversaire :", ennemi.Health)
                AfficherInventaire()
            case "Dragon Slayer":
                ennemi.Health -= 100
                fmt.Println("Arme utilisée : Dragon Slayer")
                fmt.Println("Dégâts infligés : 100")
                fmt.Println("PV restants de l'adversaire :", ennemi.Health)
                AfficherInventaire()
            }
        default:
            fmt.Println("Choix invalide.")
        }
 
        if ennemi.Health <= 0 {
            ennemi.Health = 0
            personnage.Money += ennemi.Recompense
            fmt.Println("Le monstre est vaincu !")
            fmt.Println("Tu gagnes", ennemi.Recompense, "d'or !")
            fmt.Println("Tu as maintenant", personnage.Money, "d'or.")
            AddExperience(personnage, ennemi.Experience)
 
            return false
        }
 
        degats := GoblinPattern(*ennemi, tour)
        personnage.Health -= degats
        if personnage.Health < 0 {
            personnage.Health = 0
        }
 
        fmt.Println(ennemi.Name, "te frappe pour", degats, "dégâts.")
        fmt.Println("PV du joueur :", personnage.Health)
 
        tour++
 
        if IsDead(*personnage) {
            fmt.Println("Tu as perdu !")
            return false
        }
 
    }
}
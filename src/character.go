package src

import (
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Name      string
	Classe    string
	Health    int
	Level     int
	HealthMax int
	Money     int
	Stuff     Stuff
    Experience int 
	ExperienceMax int 
}

type Objet struct {
	Nom     string
	Defense int
}

type Stuff struct {
	Casque   Objet
	Torse    Objet
	Jambiere Objet
}

var Personnage Character

func InitCharacter(name string, classe string, health int, level int, healthmax int, money int) *Character {
	return &Character{
		Name:      name,
		Classe:    classe,
		Health:    health,
		Level:     level,
		HealthMax: healthmax,
		Money:     money,
		Experience: 0,
		ExperienceMax: 100,
	}
}

func CharacterCreation() string {
	nom := AskName()
	nomFormate := FormatName(nom)

	return nomFormate
}

func AskName() string {
	for {
		var name string
		nameValide := true

		fmt.Print("comment s'appelle votre personnage: ")
		fmt.Scan(&name)

		for _, caractere := range name {
			if unicode.IsDigit(caractere) {
				nameValide = false
				break
			}
		}

		if nameValide {
			return name
		}

		fmt.Println("nom invalide : les chiffres sont interdit")
	}
}

func FormatName(name string) string {
	name = strings.ToLower(name)

	lettres := []rune(name)
	if len(lettres) == 0 {
		return ""
	}

	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

func DisplayInfo(personnage *Character) {
	fmt.Println("Nom :", personnage.Name)
	fmt.Println("Classe :", personnage.Classe)
	fmt.Println("PV :", personnage.Health)
	fmt.Println("Niveau :", personnage.Level)
	fmt.Println("PV max :", personnage.HealthMax)
	fmt.Println("Money :", personnage.Money)
}

func IsDead(personnage Character) bool {
	if personnage.Health <= 0 {
		return true
	}

	return false
}

func AddExperience(personnage *Character, experienceGagnee int) {
	if personnage.Level >= 10 {
		personnage.Level = 10
		personnage.Experience = 0
		return
	}

	personnage.Experience += experienceGagnee
	fmt.Println("XP gagnée :", experienceGagnee)

	for personnage.Experience >= personnage.ExperienceMax && personnage.Level < 10 {
		personnage.Experience -= personnage.ExperienceMax
		personnage.Level++
		personnage.ExperienceMax += 100
		personnage.HealthMax += 20
		personnage.Health = personnage.HealthMax

		fmt.Println("Niveau supérieur !")
		fmt.Println("Tu es maintenant niveau", personnage.Level)
		fmt.Println("Tes PV maximum sont maintenant de", personnage.HealthMax)
	}

	if personnage.Level == 10 {
		personnage.Experience = 0
		fmt.Println("Niveau maximum atteint !")
	}
}
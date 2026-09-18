package character 

import "fmt"

type character struct {
	Name      string
	Classe    string
	Health    int
	Level     int
	HealthMax int
}

func InitCharacter(name string, classe string, health int, level int, healthmax int) character {
	return character{
		Name:      name,
		Classe:    classe,
		Health:    health,
		Level:     level,
		HealthMax: healthmax,
	
	}
}

func AskName()string {
    var name string

    fmt.Print("comment s'appelle votre personnage: ")
	fmt.Scan(&name)

    return name
}

func DisplayInfo(personnage character) {
	fmt.Println("Nom :", personnage.Name)
    fmt.Println("Classe :", personnage.Classe)
    fmt.Println("PV :", personnage.Health)
    fmt.Println("Niveau :", personnage.Level)
    fmt.Println("PV max :", personnage.HealthMax)
}
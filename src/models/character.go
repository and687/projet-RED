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
	for{
    var name string
	nameValide := true

    fmt.Print("comment s'appelle votre personnage: ")
	fmt.Scan(&name)

	for _, caractere := range name {
		if unicode.IsDigit (caractere) {
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

	lettres [0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}


func DisplayInfo(personnage character) {
	fmt.Println("Nom :", personnage.Name)
    fmt.Println("Classe :", personnage.Classe)
    fmt.Println("PV :", personnage.Health)
    fmt.Println("Niveau :", personnage.Level)
    fmt.Println("PV max :", personnage.HealthMax)
}
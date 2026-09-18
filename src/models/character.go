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
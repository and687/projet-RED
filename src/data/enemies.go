package data

type Monster struct {
	Name       string
	Degats     int
	Health     int
	HealthMax  int
	Recompense int
}

var ennemi Monster

func InitGoblin(name string, degats int, health int, recompense int, healthmax int) Monster {
	return Monster{
		Name:       name,
		Health:     health,
		Degats:     degats,
		HealthMax:  healthmax,
		Recompense: recompense,
	}
}

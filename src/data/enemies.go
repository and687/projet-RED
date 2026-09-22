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

func GoblinPattern(monstre Monster, tour int) int {
	if tour%3 == 0 {
		return monstre.Degats * 2
	}
	return monstre.Degats
}

func GoblinPatern(monstre Monster, tour int) int {
	return GoblinPattern(monstre, tour)
}
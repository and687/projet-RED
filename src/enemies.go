package src

type Monster struct {
	Name       string
	Degats     int
	Health     int
	HealthMax  int
	Recompense int
	Experience int
}

var ennemi Monster

func InitGoblin(name string, degats int, health int, recompense int, healthmax int, experience int) Monster {
	return Monster{
		Name:       name,
		Health:     health,
		Degats:     degats,
		HealthMax:  healthmax,
		Recompense: recompense,
		Experience: experience,
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
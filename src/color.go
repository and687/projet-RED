package src

import "fmt"

const (
	Reset = "\033[0m"

	Noir   = "\033[30m"
	Rouge  = "\033[31m"
	Vert   = "\033[32m"
	Jaune  = "\033[33m"
	Bleu   = "\033[34m"
	Violet = "\033[35m"
	Cyan   = "\033[36m"
	Blanc  = "\033[37m"
	Gris   = "\033[90m"
	Orange = "\033[38;5;208m"

	Gras     = "\033[1m"
	Souligne = "\033[4m"
)

func Texte(couleur string, texte string) string {
	return couleur + texte + Reset
}

func PrintCouleur(couleur string, texte string) {
	fmt.Println(couleur + texte + Reset)
}

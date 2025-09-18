package main

import (
	"fmt"
	"strings"
	"time"
)

// =================
// visual.go : Visuels
// =================
// Ce fichier regroupe tout ce qui rend le terminal "beau" mais simple :
// - Couleurs ANSI (activables/désactivables via useColor)
// - ASCII Art (bannières simples en texte)
// - Effet "machine à écrire" (typewriter) pour la narration
// - Fonctions utilitaires : clear screen, petites pauses

var useColor = true // si l'affichage est bizarre, passer à false

// codes ANSI basiques
const (
	colReset   = "\033[0m"
	colRed     = "\033[31m"
	colGreen   = "\033[32m"
	colYellow  = "\033[33m"
	colBlue    = "\033[34m"
	colMagenta = "\033[35m"
	colCyan    = "\033[36m"
	colBold    = "\033[1m"
)

// color enveloppe un texte avec une couleur si useColor est true
func color(c, s string) string {
	if !useColor {
		return s
	}
	return c + s + colReset
}

// bold met un texte en gras (ANSI)
func bold(s string) string {
	if !useColor {
		return s
	}
	return colBold + s + colReset
}

// clear efface (quasi) l'écran : pratique avant d'afficher un écran important
func clear() {
	fmt.Print("\033[2J\033[H") // "effacer tout" + "curseur en haut à gauche"
}

// slow imprime un texte caractère par caractère (effet "machine à écrire")
func slow(s string, msPerChar int) {
	for _, r := range s {
		fmt.Printf("%c", r)
		time.Sleep(time.Duration(msPerChar) * time.Millisecond)
	}
}

// pause attend "dur" millisecondes (lisible à l'oral)
func pause(ms int) { time.Sleep(time.Duration(ms) * time.Millisecond) }

// title imprime un gros titre ASCII très simple
func title(text string) {
	fmt.Println(bold(color(colRed, "========================================")))
	fmt.Println(bold(color(colRed, "           "+text)))
	fmt.Println(bold(color(colRed, "========================================")))
}

// bannerCombat affiche un bandeau "COMBAT ENGAGÉ"
func bannerCombat() {
	art := `
(__  )/  \ (  ( \(  __)   / __)/  \ ( \/ )(  _ \ / _\(_  _)
 / _/(  O )/    / ) _)   ( (__(  O )/ \/ \ ) _ (/    \ )(  
(____)\__/ \_)__)(____)   \___)\__/ \_)(_/(____/\_/\_/(__) 
`
	// remet les vrais backticks, puis applique couleur + gras
	art = strings.ReplaceAll(art, "[BT]", "`")
	fmt.Print(centerBlock(bold(color(colRed, art)))) // change colCyan si tu veux (colMagenta, colRed, ...)
}

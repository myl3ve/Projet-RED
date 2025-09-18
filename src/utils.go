
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ===================
// utils.go : Entrées
// ===================

var reader = bufio.NewReader(os.Stdin)

func AskLine(prompt string) string {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println(color(colYellow, "Entrée vide. Réessaie."))
	}
}

func AskInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		raw, _ := reader.ReadString('\n')
		raw = strings.TrimSpace(raw)
		n, err := strconv.Atoi(raw)
		if err != nil {
			fmt.Println(color(colRed, "⛔ Entre un nombre valide."))
			continue
		}
		if n < min || n > max {
			fmt.Printf(color(colRed, "⛔ Choisis entre %d et %d.\n"), min, max)
			continue
		}
		return n
	}
}

// formatName met la première lettre en Majuscule et le reste en minuscule
func formatName(s string) string {
	if s == "" { return s }
	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}

// onlyLetters vérifie que le nom contient uniquement des lettres (Tâche 11)
func onlyLetters(s string) bool {
	ok, _ := regexp.MatchString("^[A-Za-zÀ-ÖØ-öø-ÿ]+$", s)
	return ok
}

// containsSkill / addSkill : gestion simple des compétences
func containsSkill(c *Character, skill string) bool {
	for _, s := range c.Skills {
		if s == skill { return true }
	}
	return false
}

func addSkill(c *Character, skill string) {
	if !containsSkill(c, skill) {
		c.Skills = append(c.Skills, skill)
	}
}

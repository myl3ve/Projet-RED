package ui

import (
	"fmt"
	"projet-red-monjeu/character"
	"projet-red-monjeu/combat"
)

const gameTitle = "Projet RED"

func ShowMenu() {
	for {
		fmt.Println("==============================")
		fmt.Println(" ", gameTitle)
		fmt.Println("==============================")
		fmt.Println("1) Nouvelle partie")
		fmt.Println("2) Entraînement (combat)")
		fmt.Println("3) Marchand")
		fmt.Println("4) Forgeron")
		fmt.Println("5) Quitter")

		switch Input("> ") {
		case "1":
			c := character.CharacterCreation() // Tâche 11/11-suite. :contentReference[oaicite:4]{index=4}
			character.MainLoop(&c)            // Petit sous-menu: infos/inventaire/retour (Tâche 6). :contentReference[oaicite:5]{index=5}
		case "2":
			c := character.CharacterCreation()
			combat.TrainingFight(&c) // Tâches 19–22 (gobelin, pattern, tour par tour). :contentReference[oaicite:6]{index=6}
		case "3":
			c := character.CharacterCreation()
			character.MerchantMenu(&c) // Tâches 7, 9, 14, 18. :contentReference[oaicite:7]{index=7}
		case "4":
			c := character.CharacterCreation()
			character.BlacksmithMenu(&c) // Tâche 15–17. :contentReference[oaicite:8]{index=8}
		case "5":
			fmt.Println("À plus !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

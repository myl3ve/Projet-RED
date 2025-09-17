package combat

import (
	"fmt"
	"projet-red-monjeu/character"
	"projet-red-monjeu/ui"
)

// Tâches 21–22: menu attaque/inventaire, alternance des tours, fin combat. :contentReference[oaicite:28]{index=28}
func TrainingFight(c *character.Character) {
	m := initGoblin()
	turn := 1
	fmt.Println("\n[ENTRAÎNEMENT]")

	for c.HP > 0 && m.HP > 0 {
		fmt.Printf("\n-- Tour %d --\n", turn)

		// Option Mission 1 : initiative (qui joue en premier) :contentReference[oaicite:29]{index=29}
		playerFirst := c.Initiative >= m.Initiative
		if playerFirst {
			if charTurn(c, &m) { break }
			if m.HP <= 0 { break }
			goblinPattern(turn, c, &m)
		} else {
			goblinPattern(turn, c, &m)
			if c.HP <= 0 { break }
			if charTurn(c, &m) { break }
		}
		turn++
	}

	if c.HP <= 0 {
		fmt.Println("\nTu es K.O. Retour au menu.")
		return
	}
	fmt.Println("\nLe gobelin est vaincu !")
	// Mission 2: expérience / niveau (facultatif à brancher). :contentReference[oaicite:30]{index=30}
	c.Exp += m.ExpReward
	fmt.Printf("Tu gagnes %d exp (total %d/%d)\n", m.ExpReward, c.Exp, c.ExpMax)
}

func charTurn(c *character.Character, m *Monster) bool {
	fmt.Println("Actions: 1) Attaquer  2) Inventaire  3) Fuir")
	switch ui.Input("> ") {
	case "1":
		// Attaque basique = 5 dégâts (Tâche 21). :contentReference[oaicite:31]{index=31}
		dmg := 5
		m.HP -= dmg
		if m.HP < 0 { m.HP = 0 }
		fmt.Printf("%s utilise Attaque basique : -%d PV | %s PV: %d/%d\n",
			c.Name, dmg, m.Name, m.HP, m.HPMax)
	case "2":
		character.AccessInventory(c) // consomme, puis tour du monstre. :contentReference[oaicite:32]{index=32}
	case "3":
		fmt.Println("Tu fuis l’entraînement.")
		return true
	default:
		fmt.Println("Commande inconnue.")
	}
	return false
}

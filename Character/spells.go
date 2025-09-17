package character

import "fmt"

// Tâche 10 + Missions 3/4 : sorts + mana. :contentReference[oaicite:25]{index=25}

func spellBook(c *Character) {
	for _, s := range c.Skills {
		if s == "boule de feu" {
			fmt.Println("Tu connais déjà Boule de feu.")
			return
		}
	}
	c.Skills = append(c.Skills, "boule de feu")
	fmt.Println("Nouveau sort appris: Boule de feu.")
}

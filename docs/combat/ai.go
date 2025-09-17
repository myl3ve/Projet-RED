package combat

import (
	"fmt"
	"projet-red-monjeu/character"
)

// Tâche 20: pattern gobelin (200% dmg tous les 3 tours). :contentReference[oaicite:27]{index=27}
func goblinPattern(turn int, c *character.Character, m *Monster) {
	dmg := m.ATK
	if turn%3 == 0 {
		dmg = 2 * m.ATK
	}
	c.HP -= dmg
	if c.HP < 0 { c.HP = 0 }
	fmt.Printf("%s inflige à %s %d dégâts | PV: %d/%d\n",
		m.Name, c.Name, dmg, c.HP, c.HPMax)
}

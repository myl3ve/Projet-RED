package character

import (
	"fmt"
	"projet-red-monjeu/ui"
	"strings"
	"time"
)

// Tâche 4/5/6/8/9/12/18. :contentReference[oaicite:14]{index=14}

func AccessInventory(c *Character) {
	for {
		fmt.Printf("\nInventaire (%d/%d) : %v\n", len(c.Inventory), c.InventoryCap, c.Inventory)
		fmt.Println("Choix: [use <nom>]  [retour]")
		cmd := ui.Input("> ")
		if cmd == "retour" {
			return
		}
		if strings.HasPrefix(cmd, "use ") {
			item := strings.TrimSpace(strings.TrimPrefix(cmd, "use "))
			useItem(c, item)
		}
	}
}

func useItem(c *Character, item string) {
	switch item {
	case "potion":
		takePot(c) // +50 PV, consomme l’item (Tâche 5). :contentReference[oaicite:15]{index=15}
	case "potion de poison":
		poisonPot(c) // 10 PV/s * 3s (Tâche 9). :contentReference[oaicite:16]{index=16}
	case "livre de sort : boule de feu":
		spellBook(c) // ajoute le sort si pas déjà appris (Tâche 10). :contentReference[oaicite:17]{index=17}
	case "augmentation d’inventaire":
		upgradeInventorySlot(c) // +10, max 3 fois (Tâche 18). :contentReference[oaicite:18]{index=18}
	default:
		fmt.Println("Objet inutilisable.")
	}
}

func takePot(c *Character) {
	if removeInventory(c, "potion") {
		c.HP += 50
		if c.HP > c.HPMax {
			c.HP = c.HPMax
		}
		fmt.Printf("Tu utilises une potion. PV: %d/%d\n", c.HP, c.HPMax)
	} else {
		fmt.Println("Aucune potion.")
	}
}

func poisonPot(c *Character) {
	if !removeInventory(c, "potion de poison") {
		fmt.Println("Pas de potion de poison.")
		return
	}
	for i := 0; i < 3; i++ {
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("Le poison te ronge... PV: %d/%d\n", c.HP, c.HPMax)
		time.Sleep(1 * time.Second) // Tâche 9. :contentReference[oaicite:19]{index=19}
	}
	isDead(c) // Tâche 8: si 0 → revive à 50% PV max. :contentReference[oaicite:20]{index=20}
}

func isDead(c *Character) {
	if c.HP <= 0 {
		fmt.Println("WASTED. Tu reviens à la vie à 50% PV max.")
		c.HP = c.HPMax / 2
	}
}

func addInventory(c *Character, item string) bool {
	if len(c.Inventory) >= c.InventoryCap {
		fmt.Println("Inventaire plein (max 10 par défaut).") // Tâche 12. :contentReference[oaicite:21]{index=21}
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func removeInventory(c *Character, item string) bool {
	for i, it := range c.Inventory {
		if it == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

var upgradesUsed = 0

func upgradeInventorySlot(c *Character) {
	if upgradesUsed >= 3 {
		fmt.Println("Tu as déjà atteint la limite d’augmentation d’inventaire (3).")
		return
	}
	c.InventoryCap += 10
	upgradesUsed++
	fmt.Printf("Capacité d’inventaire augmentée à %d (utilisations: %d/3)\n", c.InventoryCap, upgradesUsed)
}

// Petit sous-menu pour Tâche 6 (afficher infos / inventaire / retour)
func MainLoop(c *Character) {
	for {
		fmt.Println("\n1) Afficher informations")
		fmt.Println("2) Accéder à l’inventaire")
		fmt.Println("3) Retour menu principal")
		switch ui.Input("> ") {
		case "1":
			DisplayInfo(c)
		case "2":
			AccessInventory(c)
		case "3":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

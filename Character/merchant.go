package character

import (
	"fmt"
	"projet-red-monjeu/ui"
)

// Tâches 7, 9, 14, 18 — marchand + coûts. :contentReference[oaicite:22]{index=22}

func MerchantMenu(c *Character) {
	for {
		fmt.Printf("\n(Marchand) Or: %d | Inventaire: %d/%d\n", c.Money, len(c.Inventory), c.InventoryCap)
		fmt.Println("[1] Potion de vie (3 or)")
		fmt.Println("[2] Potion de poison (6 or)")
		fmt.Println("[3] Livre de Sort : Boule de feu (25 or)")
		fmt.Println("[4] Fourrure de loup (4 or)")
		fmt.Println("[5] Peau de troll (7 or)")
		fmt.Println("[6] Cuir de sanglier (3 or)")
		fmt.Println("[7] Plume de corbeau (1 or)")
		fmt.Println("[8] Augmentation d’inventaire (30 or)")
		fmt.Println("[0] Retour")
		switch ui.Input("> ") {
		case "1":
			buy(c, "potion", 3)
		case "2":
			buy(c, "potion de poison", 6)
		case "3":
			buy(c, "livre de sort : boule de feu", 25)
		case "4":
			buy(c, "fourrure de loup", 4)
		case "5":
			buy(c, "peau de troll", 7)
		case "6":
			buy(c, "cuir de sanglier", 3)
		case "7":
			buy(c, "plume de corbeau", 1)
		case "8":
			if buy(c, "augmentation d’inventaire", 30) {
				upgradeInventorySlot(c)
				removeInventory(c, "augmentation d’inventaire")
			}
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func buy(c *Character, item string, cost int) bool {
	if len(c.Inventory) >= c.InventoryCap {
		fmt.Println("Inventaire plein.")
		return false
	}
	if c.Money < cost {
		fmt.Printf("Pas assez d’or (coût: %d).\n", cost)
		return false
	}
	c.Money -= cost
	addInventory(c, item)
	fmt.Println("Achat:", item)
	return true
}

package character

import (
	"fmt"
	"projet-red-monjeu/ui"
)

// Tâche 15–17: craft + coûts + ressources + bonus PV max. :contentReference[oaicite:23]{index=23}

func BlacksmithMenu(c *Character) {
	for {
		fmt.Println("\n(Forgeron) 5 or par fabrication. Choisir :")
		fmt.Println("[1] Chapeau de l’aventurier (1 Plume de corbeau + 1 Cuir de sanglier)")
		fmt.Println("[2] Tunique de l’aventurier (2 Fourrure de loup + 1 Peau de troll)")
		fmt.Println("[3] Bottes de l’aventurier (1 Fourrure de loup + 1 Cuir de sanglier)")
		fmt.Println("[0] Retour")

		switch ui.Input("> ") {
		case "1":
			craft(c, "chapeau de l’aventurier", 5, map[string]int{
				"plume de corbeau": 1, "cuir de sanglier": 1,
			})
		case "2":
			craft(c, "tunique de l’aventurier", 5, map[string]int{
				"fourrure de loup": 2, "peau de troll": 1,
			})
		case "3":
			craft(c, "bottes de l’aventurier", 5, map[string]int{
				"fourrure de loup": 1, "cuir de sanglier": 1,
			})
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func craft(c *Character, equip string, gold int, req map[string]int) {
	if len(c.Inventory) >= c.InventoryCap {
		fmt.Println("Inventaire plein.")
		return
	}
	if c.Money < gold {
		fmt.Printf("Pas assez d’or (coût: %d).\n", gold)
		return
	}
	// check ressources
	for k, v := range req {
		if countItem(c, k) < v {
			fmt.Printf("Ressource manquante: %s x%d\n", k, v)
			return
		}
	}
	// consommer ressources
	for k, v := range req {
		for i := 0; i < v; i++ {
			removeInventory(c, k)
		}
	}
	c.Money -= gold
	addInventory(c, equip)
	fmt.Println("Fabriqué:", equip, "(ajouté à l’inventaire)")

	// possibilité d’équiper directement depuis l’inventaire
	EquipFromInventory(c, equip)
}

func countItem(c *Character, item string) int {
	n := 0
	for _, it := range c.Inventory {
		if it == item {
			n++
		}
	}
	return n
}

func EquipFromInventory(c *Character, equip string) {
	switch equip {
	case "chapeau de l’aventurier":
		swapEquip(c, &c.Equipment.Head, equip, 10)
	case "tunique de l’aventurier":
		swapEquip(c, &c.Equipment.Torso, equip, 25)
	case "bottes de l’aventurier":
		swapEquip(c, &c.Equipment.Feet, equip, 15)
	}
}

func swapEquip(c *Character, slot *string, equip string, bonus int) {
	// si déjà un equipement, on le remet à l’inventaire (remplacement) — Tâche 17. :contentReference[oaicite:24]{index=24}
	if *slot != "" {
		addInventory(c, *slot)
	}
	*slot = equip
	removeInventory(c, equip)
	c.HPMax += bonus
	if c.HP > c.HPMax {
		c.HP = c.HPMax
	}
	fmt.Printf("Équipé: %s | PV max = %d\n", equip, c.HPMax)
}

package main

import "fmt"

// =====================
// merchant.go : Boutique
// =====================
// Conformité PDF (Tâches 7,13,14,18)

var prices = map[string]int{
	"Potion de vie":       3,
	"Potion de poison":    6,
	"couteau":             25,
	"opinel 12":           4,
	"botte du daron":      7,
	"gilet pare ball":     3,
	"stylos plume":        1,
	"Augment. Inventaire": 30,
}

func showMerchant(c *Character) {
	for {
		fmt.Printf("\n=== Marchand (Or: %d) ===\n", c.Gold)
		options := []string{
			"Potion de vie", "Potion de poison", "Livre : katana",
			"opinel 12", "ciseau", "couteau papillon", "botte du darons",
			"Augment. Inventaire",
			"Retour",
		}
		for i, o := range options {
			if o != "Retour" {
				fmt.Printf("%d) %s (%d or)\n", i+1, o, prices[o])
			} else {
				fmt.Printf("%d) %s\n", i+1, o)
			}
		}
		choice := AskInt("Choix: ", 1, len(options))
		if options[choice-1] == "Retour" {
			return
		}
		itemName := options[choice-1]
		price := prices[itemName]

		if c.Gold < price {
			fmt.Println(color(colRed, "⛔ Pas assez d'or."))
			continue
		}
		if !checkInventoryCap(c) {
			fmt.Println(color(colRed, "⛔ Pas de place dans l'inventaire."))
			continue
		}
		c.Gold -= price

		typ := Consumable
		switch itemName {
		case "Livre :ciseau":
			typ = Misc
		case "Augment. Inventaire":
			typ = Upgrade
		case "ciseau", "couteau papillon", "stylos plume", "opinel 12":
			typ = Material
		}
		addInventory(c, Item{Name: itemName, Type: typ})
	}
}

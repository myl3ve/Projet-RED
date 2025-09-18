package main // Ce fichier fait partie du package exécutable "main".

import ( // Début du bloc d'import.
	"fmt"  // Package d'affichage formaté (Println, Printf, ...).
	"time" // Package pour gérer le temps (Sleep, durées, etc.).
) // Fin du bloc d'import.

// ========================
// inventory.go : Inventaire
// ========================

// checkInventoryCap vérifie la limite actuelle (Tâche 12)
func checkInventoryCap(c *Character) bool { // Fonction qui retourne vrai si on peut encore ajouter un objet.
	return len(c.Inventory) < c.InventoryCap // Compare le nombre d'objets avec la capacité maximale.
} // Fin checkInventoryCap.

// addInventory (Tâche 7) : ajoute si y'a de la place
func addInventory(c *Character, it Item) bool { // Ajoute un item dans l'inventaire si possible; renvoie succès/échec.
	if !checkInventoryCap(c) { // Si l'inventaire est plein...
		fmt.Println(color(colRed, "⛔ Inventaire plein.")) // ...on affiche un message d'erreur coloré.
		return false // ...et on indique l'échec.
	} // Fin du if plein.
	c.Inventory = append(c.Inventory, it) // Ajoute l'item à la fin de la slice Inventory.
	fmt.Println(color(colGreen, "✅ Ajouté : "), it.Name) // Message de confirmation (nom de l'item ajouté).
	return true // Indique que l'ajout a réussi.
} // Fin addInventory.

// removeInventory (Tâche 7) : supprime la 1ère occurrence
func removeInventory(c *Character, name string) bool { // Supprime la première occurrence d'un item par son nom.
	for i, it := range c.Inventory { // Parcourt l'inventaire avec index i et item it.
		if it.Name == name { // Si le nom correspond...
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...) // ...on retire l'élément par découpage/soudure de slices.
			return true // On confirme qu'on a bien supprimé quelque chose.
		} // Fin du if nom trouvé.
	} // Fin du for sur l'inventaire.
	return false // Rien trouvé à supprimer → échec.
} // Fin removeInventory.

// usePotionVie (Tâche 5)
func usePotionVie(c *Character) { // Utilise une potion de vie si disponible.
	if !removeInventory(c, "Potion de vie") { // Tente de retirer une "Potion de vie" de l'inventaire.
		fmt.Println(color(colYellow, "⛔ Pas de Potion de vie.")) // Message si aucune potion trouvée.
		return // On quitte sans effet.
	} // Fin du if pas de potion.
	c.HP += 50 // Soigne 50 PV.
	if c.HP > c.MaxHP { // Si on dépasse le maximum...
		c.HP = c.MaxHP // ...on borne au MaxHP.
	} // Fin du bornage PV.
	fmt.Printf(color(colGreen, "🧪 +50 PV → %d/%d\n"), c.HP, c.MaxHP) // Affiche l'état des PV après soin.
} // Fin usePotionVie.

// applyPoison (Tâche 9) : 10 PV/s pendant 3s avec time.Sleep
func applyPoison(c *Character) { // Applique un poison: retire 10 PV par seconde pendant 3 secondes.
	if !removeInventory(c, "Potion de poison") { // Vérifie et consomme une potion de poison avant d'appliquer l'effet.
		fmt.Println(color(colYellow, "⛔ Pas de Potion de poison.")) // Message si absente.
		return // On quitte sans effet.
	} // Fin du if pas de potion de poison.
	for i := 1; i <= 3; i++ { // Boucle 3 "ticks" de poison.
		time.Sleep(1 * time.Second) // EXIGÉ par le PDF : pause d'une seconde entre chaque tick.
		c.HP -= 10 // Retire 10 PV.
		if c.HP < 0 { c.HP = 0 } // Sécurité : borne à 0 pour éviter PV négatifs.
		fmt.Printf(color(colMagenta, "☠️ Poison %d/3 → PV: %d/%d\n"), i, c.HP, c.MaxHP) // Affiche l'état après chaque tick.
		if isDead(c) { // Si le personnage meurt (et éventuellement "ressuscité" par la logique isDead)...
			return // ...on arrête l'effet de poison.
		} // Fin test mort.
	} // Fin de la boucle poison.
} // Fin applyPoison.

// useSpellBook (Tâche 10) : apprend Boule de feu (sans doublon)
func useSpellBook(c *Character) { // Utilise un grimoire pour apprendre un sort (sans le dupliquer).
	if containsSkill(c, "Boule de feu") { // Vérifie si la compétence est déjà apprise.
		fmt.Println(color(colYellow, "ℹ️ Boule de feu déjà apprise.")) // Message d'info si doublon.
	} else { // Sinon, on peut l'apprendre.
		addSkill(c, "Boule de feu") // Ajoute la compétence au personnage.
		fmt.Println(color(colGreen, "📘 Sort appris : Boule de feu !")) // Message de confirmation.
	} // Fin du if/else.
	removeInventory(c, "Livre : Boule de feu") // Retire le livre de l'inventaire (consommé).
} // Fin useSpellBook.

// useItem route l'effet selon le nom / type
func useItem(c *Character, index int) { // Utilise l'objet à l'index donné dans l'inventaire.
	if index < 0 || index >= len(c.Inventory) { // Vérifie que l'index est dans les bornes.
		fmt.Println(color(colRed, "⛔ Index invalide.")) // Message d'erreur si invalide.
		return // On quitte.
	} // Fin du test de bornes.
	name := c.Inventory[index].Name // Récupère le nom de l'objet à utiliser.
	switch name { // Dispatch selon le nom de l'objet.
	case "Potion de vie": // Si c'est une potion de vie...
		usePotionVie(c) // ...on soigne.
	case "Potion de poison": // Si c'est une potion de poison...
		applyPoison(c) // ...on applique l'effet de poison (dégâts sur 3s).
	case "Livre : Boule de feu": // Si c'est un grimoire...
		useSpellBook(c) // ...on apprend le sort (si pas déjà appris).
	case "Augment. Inventaire": // Objet d'augmentation de capacité d'inventaire.
		// Tâche 18 : +10, max 3 fois
		if c.InventoryCap >= 40 { // Si on a déjà atteint le plafond (10 de base + 3*10).
			fmt.Println(color(colYellow, "⛔ Capacité déjà au maximum.")) // On informe que c'est au max.
		} else { // Sinon, on peut augmenter.
			c.InventoryCap += 10 // Ajoute 10 à la capacité.
			fmt.Println(color(colGreen, "📦 Capacité inventaire +10 → "), c.InventoryCap) // Affiche la nouvelle capacité.
		} // Fin du if/else d'augmentation.
		removeInventory(c, name) // Dans tous les cas, on consomme l'objet d'augmentation.
	default: // Par défaut: objet non utilisable directement.
		fmt.Println(color(colYellow, "ℹ️ Cet objet ne peut pas être utilisé directement : "), name) // Message informatif.
	} // Fin du switch.
} // Fin useItem.

// accessInventory (Tâche 4 + lien Tâche 7)
func accessInventory(c *Character) { // Ouvre l'interface d'inventaire en boucle (jusqu'à "Retour").
	for { // Boucle infinie de navigation dans l'inventaire.
		fmt.Printf("\n=== Inventaire (%d/%d) ===\n", len(c.Inventory), c.InventoryCap) // Affiche le compteur d'objets / capacité.
		for i, it := range c.Inventory { // Liste tous les items avec leur index.
			fmt.Printf("%d) %s\n", i+1, it.Name) // Affiche "1) Nom", "2) Nom", etc.
		} // Fin de l'affichage des items.
		fmt.Printf("%d) Retour\n", len(c.Inventory)+1) // Ajoute une option "Retour" à la fin.
		choice := AskInt("Choix: ", 1, len(c.Inventory)+1) // Demande un choix valide (1..N+1).
		if choice == len(c.Inventory)+1 { return } // Si l'utilisateur choisit "Retour" → on sort de la fonction.
		// équiper si c'est un équipement connu
		item := c.Inventory[choice-1] // Récupère l'item choisi (index 0-based).
		if item.Type == EquipmentItem && // Si c'est un type équipement...
			(item.Name == "Chapeau de l'aventurier" || item.Name == "Tunique de l'aventurier" || item.Name == "Bottes de l'aventurier") { // ...et l'un des 3 équipements gérés.
			equip(c, item.Name) // On équipe l'objet (gère les slots + recalcul des PV max).
		} else { // Sinon, ce n'est pas un équipement portable géré...
			useItem(c, choice-1) // ...on tente de l'utiliser (soin, poison, livre, etc.).
		} // Fin du if équipement.
	} // Fin de la boucle d'inventaire (revient au début après action).
} // Fin accessInventory.

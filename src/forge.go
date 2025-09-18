package main // Ce fichier appartient au package exécutable "main".

import "fmt" // On importe fmt pour afficher du texte (Printf, Println, etc.).

// ==================
// forge.go : Forgeron
// ==================
// Tâche 15 + ressources PDF + messages clairs // Commentaire de contexte (pas exécuté).

func showForge(c *Character) {            // Fonction qui affiche le menu du forgeron pour le personnage c.
	for {                                  // Boucle infinie : on revient au menu après chaque action jusqu’à "Retour".
		fmt.Printf("\n=== Forgeron (Or: %d) ===\n", c.Gold) // Affiche l'or actuel du joueur.
		options := []string{               // Déclare la liste des options craftables + "Retour".
			"Chapeau de l'aventurier", // Plume + ciseau // Libellé 1 (commentaire recette, pure info).
			"Tunique de l'aventurier", // 2x cips + opinel // Libellé 2.
			"Bottes de l'aventurier",  // Fourrure + bottes daron // Libellé 3.
			"Retour",                  // Option 4 : sortir du menu forgeron.
		}
		for i, o := range options {        // Parcourt toutes les options pour les afficher numérotées.
			fmt.Printf("%d) %s\n", i+1, o) // Affiche "1) ..." etc.
		}
		choice := AskInt("Que fabriquer ? ", 1, len(options)) // Demande un choix entre 1 et le nombre d’options.
		if options[choice-1] == "Retour" { return }            // Si "Retour" choisi → on sort de showForge.

		target := options[choice-1]       // Récupère le nom de l’objet à fabriquer selon le choix.
		if c.Gold < 5 {                   // Vérifie si le joueur a au moins 5 or.
			fmt.Println(color(colRed, "⛔ Il faut 5 or pour fabriquer.")) // Message d’erreur si pas assez d’or.
			continue                  // Revient au début de la boucle (menu à nouveau).
		}
		if !checkInventoryCap(c) {        // Vérifie si l’inventaire a de la place.
			fmt.Println(color(colRed, "⛔ Inventaire plein.")) // Avertit si inventaire saturé.
			continue
		}
		if !hasResourcesFor(c, target) {  // Vérifie si le joueur possède les ressources requises pour target.
			fmt.Println(color(colYellow, "⛔ Ressources manquantes pour : "), target) // Message si ressources insuffisantes.
			continue
		}
		consumeResourcesFor(c, target)    // Consomme (retire) les ressources nécessaires dans l’inventaire.
		c.Gold -= 5                       // Retire le coût en or (5).
		addInventory(c, Item{Name: target, Type: EquipmentItem}) // Ajoute l’objet crafté à l’inventaire (type équipement).
		fmt.Println(color(colGreen, "🛠️ Fabriqué → "), target) // Confirme la fabrication.
	} // Fin de la boucle for (jamais atteint sans "Retour").
} // Fin de showForge.

func hasResourcesFor(c *Character, target string) bool { // Vérifie si on a les ressources pour fabriquer "target".
	need := map[string]int{}            // Map des ressources nécessaires (nom → quantité).
	if target == "Chapeau de l'aventurier" { // Si on veut fabriquer un chapeau…
		need["ciseau"] = 1              // …il faut 1 ciseau…
		need["couteau"] = 1             // …et 1 couteau. (noms exacts exigés)
	} else if target == "katana" {      // Autre cas : si la cible est "katana"…
		need["opinel 9"] = 2            // …il faut 2 "opinel 9"…
		need["couteau papillon"] = 1    // …et 1 "couteau papillon".
	} else if target == "Bottes de daron" { // Autre cas : si la cible est "Bottes de daron"…
		need["canette de coca"] = 1     // …il faut 1 canette…
		need["cips"] = 1                // …et 1 "cips".
	}
	// compter l'inventaire
	count := map[string]int{}           // Map pour compter les objets possédés (nom → quantité détenue).
	for _, it := range c.Inventory {    // Parcourt l’inventaire du joueur.
		count[it.Name]++                // Incrémente le compteur pour le nom de l’objet.
	}
	for name, qty := range need {       // Parcourt chaque ressource requise…
		if count[name] < qty { return false } // …et vérifie si on en a assez ; sinon → false.
	}
	return true                         // Si toutes les ressources sont suffisantes → true.
} // Fin de hasResourcesFor.

func consumeResourcesFor(c *Character, target string) { // Retire les ressources correspondantes à "target".
	if target == "Chapeau de l'aventurier" { // Si on fabrique un chapeau…
		removeInventory(c, "du vin")         // …retire "du vin" (nom exact exigé).
		removeInventory(c, "une camera")     // …retire "une camera".
	} else if target == "un telephone" {    // Si on fabrique "un telephone"…
		removeInventory(c, "couteau")        // …retire "couteau".
		removeInventory(c, "puff")           // …retire "puff".
		removeInventory(c, "briquet")        // …retire "briquet".
	} else if target == "couteau papillon" { // Si on fabrique "couteau papillon"…
		removeInventory(c, "gilet pare ball") // …retire "gilet pare ball".
		removeInventory(c, "fusl a pompe")    // …retire "fusl a pompe".
	}
} // Fin de consumeResourcesFor.

// equip applique l'objet crafté sur le bon slot et met à jour MaxHP (Tâche 17) // Commentaire doc.
func equip(c *Character, name string) { // Équipe l’objet "name" sur le slot adapté et recalcule les PV max.
	switch name {                        // On choisit le slot selon le nom de l’équipement.
	case "Chapeau de l'aventurier":      // Si c’est un chapeau…
		if c.Equip.Head != "" {          // …et qu’un autre chapeau est déjà porté…
			addInventory(c, Item{Name: c.Equip.Head, Type: EquipmentItem}) // …on remet l’ancien chapeau dans l’inventaire.
		}
		c.Equip.Head = name              // …puis on équipe le nouveau chapeau.
	case "Tunique de l'aventurier":      // Si c’est une tunique…
		if c.Equip.Body != "" {          // …et qu’une armure est déjà portée…
			addInventory(c, Item{Name: c.Equip.Body, Type: EquipmentItem}) // …on remet l’ancienne au sac.
		}
		c.Equip.Body = name              // …puis on équipe la tunique.
	case "Bottes de l'aventurier":       // Si ce sont des bottes…
		if c.Equip.Feet != "" {          // …et que des bottes sont déjà portées…
			addInventory(c, Item{Name: c.Equip.Feet, Type: EquipmentItem}) // …on remet les anciennes au sac.
		}
		c.Equip.Feet = name              // …puis on équipe les nouvelles bottes.
	default:                             // Cas par défaut : nom inconnu.
		fmt.Println(color(colYellow, "⛔ Équipement inconnu : "), name) // Avertit que l’objet n’est pas reconnu.
		return                       // Sort sans rien changer.
	}
	// Recalcule le MaxHP de base + bonus
	base := 100                      // PV de base par défaut = 100.
	if c.Class == "Elfe" { base = 80 }  // Les Elfes ont moins de PV de base.
	if c.Class == "Nain" { base = 120 } // Les Nains ont plus de PV de base.

	bonus := 0                       // Bonus total initialisé à 0.
	if c.Equip.Head == "Chapeau de l'aventurier" { bonus += 10 } // Chapeau → +10 PV.
	if c.Equip.Body == "Tunique de l'aventurier" { bonus += 25 } // Tunique → +25 PV.
	if c.Equip.Feet == "Bottes de l'aventurier" { bonus += 15 }  // Bottes → +15 PV.
	c.MaxHP = base + bonus          // Nouveau MaxHP = base + bonus d’équipement.
	if c.HP > c.MaxHP { c.HP = c.MaxHP } // Si les PV actuels dépassent le nouveau max, on les borne.

	// Retirer l'objet devenu "porté"
	removeInventory(c, name)        // Retire de l’inventaire l’objet maintenant équipé (pour éviter le doublon).
	fmt.Println(color(colGreen, "🛡️ Équipement mis à jour. MaxHP : "), c.MaxHP) // Message de confirmation avec le nouveau MaxHP.
} // Fin de equip.

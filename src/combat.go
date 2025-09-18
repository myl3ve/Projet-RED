package main // Le fichier appartient au package exécutable "main".

import (          // Début du bloc d'importations.
	"fmt"       // Importe le package fmt pour l'affichage formaté (Println, Printf, etc.).
	"strings"   // Importe strings pour utiliser des fonctions sur les chaînes (Repeat, etc.).
) // Fin du bloc d'importations.

//
// ===================
// combat.go : Combat
// ===================
// Ce fichier gère TOUT ce qui concerne le COMBAT côté console.
// Objectifs :
// 1) Lisible à l’écran (couleurs + centrage + HUD) ✅
// 2) Conforme au sujet PDF (Tâches 19→22.2) ✅
//
// Dépendances utilisées (déjà présentes ailleurs dans le projet) :
// - centerLine(), color(), bold(), bannerCombat()  -> src/visual.go
// - AskInt(), containsSkill(), useItem()          -> src/utils.go / src/inventory.go
// - Structs Character, Monster                    -> src/types.go
//
// IMPORTANT : la campagne Histoire utilise sa propre fonction fight(...) dans story.go.
// Ici, on fournit l’ENTRAÎNEMENT + le tour joueur + le pattern ennemi + la détection mort.
//

// ==============================
// 19) Monstre d'entraînement
// ==============================
// On renvoie un Gobelin très simple : 40 PV / 5 dégâts / init 8 / +30 XP.
// (Parfait pour les tests et la démonstration en classe.)
func initGoblin() Monster { // Déclare une fonction qui construit et retourne un Monster d'entraînement.
	return Monster{           // Retourne un littéral de struct Monster avec les champs initialisés.
		Name:       "Gobelin d'entrainement", // Nom du monstre.
		MaxHP:      40,                       // Points de vie maximum.
		HP:         35,                       // Points de vie actuels (ici 35 pour commencer entamé).
		Attack:     5,                        // Dégâts de base par attaque.
		Initiative: 8,                        // Score d'initiative (détermine qui commence).
		ExpReward:  30,                       // Expérience gagnée par le joueur en cas de victoire.
	} // Fin du littéral Monster.
} // Fin de initGoblin.

// =====================================================
// (Aide visuelle) HUD = “barre d’état” centrée & colorée
// =====================================================
// Affiche clairement les PV du joueur et du monstre.
// On le rappelle à chaque tour pour guider le joueur visuellement.
func printHUD(c *Character, g *Monster) {                           // Affiche un encadré avec les PV des deux combattants.
	sep := strings.Repeat("─", 60)                                   // Crée une ligne de séparation (60 tirets).
	fmt.Println(centerLine(color(colCyan, sep)))                     // Affiche la ligne, centrée et en cyan.
	line := fmt.Sprintf("Toi: %d/%d PV  |  %s: %d/%d PV",            // Formate le texte avec les PV joueur/monstre.
		c.HP, c.MaxHP, g.Name, g.HP, g.MaxHP)                        // Valeurs injectées dans la ligne formatée.
	fmt.Println(centerLine(bold(color(colCyan, line))))              // Affiche la ligne d’infos en cyan et en gras, centrée.
	fmt.Println(centerLine(color(colCyan, sep)))                     // Ré-affiche une ligne de séparation pour fermer le HUD.
} // Fin de printHUD.

// =====================================
// 20) Pattern de l’ennemi (goblinPattern)
// =====================================
// Le monstre frappe à chaque tour.
// Tous les 3 tours (3, 6, 9, …) il met un COUP FORT (dégâts x2).
// On affiche un texte clair pour prévenir le joueur.
func goblinPattern(g Monster, c *Character, turn int) {             // Logique d'attaque ennemie en fonction du tour.
	dmg := g.Attack                                                  // Commence avec les dégâts de base du monstre.
	isBig := false                                                   // Indicateur si l'attaque est un "coup fort".
	if turn%3 == 0 {                                                 // Tous les 3 tours (si le tour est multiple de 3)...
		dmg = g.Attack * 2                                           // ...les dégâts sont doublés.
		isBig = true                                                 // Marque que c'est un coup fort (pour l'affichage).
	}                                                                // Fin du bloc conditionnel.

	// Message “attaque normale” vs “coup fort”
	if isBig {                                                       // Si c'est un coup fort...
		fmt.Println(centerLine(bold(color(colRed,                    // Affiche un avertissement centré, en rouge et en gras.
			"⚠️  COUP FORT de l'ennemi (x2 dégâts) !"))))
	}                                                                // Fin du if isBig.
	fmt.Printf("%s inflige à %s %d dégâts\n", g.Name, c.Name, dmg)   // Message standard d'attaque avec les dégâts infligés.

	// On retire les PV et on borne à 0
	c.HP -= dmg                                                      // Décrémente les PV du joueur des dégâts calculés.
	if c.HP < 0 {                                                    // Si les PV passent sous 0...
		c.HP = 0                                                     // ...on borne à 0 pour éviter des PV négatifs.
	}                                                                // Fin du bornage des PV.
	fmt.Printf("%s PV : %d/%d\n", c.Name, c.HP, c.MaxHP)             // Affiche l'état actuel des PV du joueur.

	// 8) Détection mort : on gère la “résurrection 50%” pour ne pas bloquer la démo
	isDead(c)                                                        // Appelle la fonction de détection mort (gère résurrection 50%).
} // Fin de goblinPattern.

// =======================
// 8) Mort / Résurrection
// =======================
// Si le joueur meurt (PV ≤ 0), on affiche un message, puis on le remet à 50% de ses PV max.
// C’est une exigence pédagogique du sujet : ne PAS “soft-locker” l’exercice.
func isDead(c *Character) bool {                                     // Vérifie si le joueur est mort; renvoie vrai/faux.
	if c.HP > 0 {                                                    // Si le joueur a encore des PV...
		return false                                                 // ...il n'est pas mort → on renvoie false.
	}                                                                // Fin du test de vie.
	fmt.Println(centerLine(color(colRed, "💀 Vous êtes mort.")))     // Affiche un message de mort centré et rouge.
	c.HP = c.MaxHP / 2                                               // “Résurrection” : restaure à 50% des PV max.
	// On montre le nouveau total de PV avec couleur et centrage
	fmt.Println(centerLine(fmt.Sprintf(                              // Affiche le nouveau total de PV de façon stylée.
		color(colGreen, "✨ Résurrection à %d/%d PV"), c.HP, c.MaxHP)))
	return true                                                      // Indique que le joueur était effectivement mort.
} // Fin de isDead.

// ==========================
// 21) Tour du joueur (menu)
// ==========================
// Le joueur choisit : Attaquer / Inventaire / (Armes s’il en a).
// → Attaque basique = 5 dégâts (conformité PDF).
// → Inventaire : on peut, par ex., boire une Potion de vie ; le tour se termine ensuite.
// → Armes (Fusil/AK47/MP5) : débloquées par l’Histoire, dégâts plus forts.
func charTurn(c *Character, g *Monster) {                            // Gère tout le tour du joueur (menu + action).
	fmt.Println()                                                    // Affiche une ligne vide (espacement).
	fmt.Println(centerLine(bold(color(colRed, "=== TON TOUR ===")))) // Titre centré "TON TOUR" en rouge et en gras.
	fmt.Println(centerLine(color(colYellow,                          // Petit conseil sur le pattern de l’ennemi.
		"Conseil : l'ennemi frappe x2 aux tours 3, 6, 9, ...")))

	// Construction dynamique du menu (numéros propres, même si pas toutes les armes)
	type option struct {                                             // Déclare un type interne pour une option de menu.
		label string                                                 // Texte affiché dans le menu.
		do    func()                                                 // Fonction exécutée lorsqu'on choisit l'option.
	}                                                                // Fin du type option.
	options := []option{                                             // Crée la tranche d'options de base (attaque, inventaire).
		{
			label: "Attaque basique (5 dégâts)",                     // Libellé de l'option 1.
			do: func() {                                             // Action : attaque de base.
				dmg := 5                                            // Dégâts fixes = 5.
				g.HP -= dmg                                          // On enlève 5 PV au monstre.
				if g.HP < 0 {                                        // Si PV < 0...
					g.HP = 0                                         // ...borne à 0 pour éviter négatifs.
				}                                                    // Fin bornage PV monstre.
				fmt.Printf("%s utilise Attaque basique → %d dégâts | %s PV: %d/%d\n", // Affiche le résumé de l'action.
					c.Name, dmg, g.Name, g.HP, g.MaxHP)
			},                                                       // Fin de la fonction do pour attaque basique.
		},
		{
			label: "Ouvrir l'inventaire",                             // Libellé de l'option 2.
			do: func() {                                             // Action : ouvrir l'inventaire et potentiellement utiliser un objet.
				if len(c.Inventory) == 0 {                           // Si l'inventaire est vide...
					fmt.Println(color(colYellow, "Inventaire vide."))// ...on informe et on termine l'action.
					return                                            // Retour (rien d'autre à faire).
				}                                                    // Fin test inventaire vide.
				// Petit sous-menu “quel objet utiliser ?”
				for i, it := range c.Inventory {                     // Parcourt les objets pour les lister numérotés.
					fmt.Printf("%d) %s\n", i+1, it.Name)            // Affiche l'index (1-based) et le nom de l'objet.
				}                                                    // Fin de la boucle d'affichage des objets.
				fmt.Printf("%d) Retour\n", len(c.Inventory)+1)       // Ajoute une option "Retour".
				idx := AskInt("→ Utiliser quoi ? ", 1,               // Demande à l'utilisateur de choisir un index...
					len(c.Inventory)+1)                              // ...entre 1 et nombre d'objets + 1 (Retour).
				if idx == len(c.Inventory)+1 {                       // S'il choisit "Retour"...
					return                                            // ...on quitte l'action inventaire sans utiliser.
				}                                                    // Fin du test "Retour".
				useItem(c, idx-1)                                    // Utilise l'objet sélectionné (converti en index 0-based).
				// NOTE : comme dans le PDF, après une utilisation d’objet,
				//        on laisse la main à l’ennemi (le tour est “consommé”).
			},                                                       // Fin de la fonction do pour l'inventaire.
		},
	} // Fin de l'initialisation des options de base.

	// Options d’armes (débloquées par l’Histoire)
	if containsSkill(c, "Fusil") {                                  // Si le perso possède la compétence "Fusil"...
		options = append(options, option{                           // ...on ajoute une entrée de menu supplémentaire.
			label: "Tir au Fusil (12 dégâts)",                      // Libellé pour l'arme Fusil.
			do: func() {                                            // Action : tirer au fusil.
				d := 12                                             // Dégâts du fusil.
				g.HP -= d                                           // On enlève ces PV au monstre.
				if g.HP < 0 {                                       // Si PV < 0...
					g.HP = 0                                        // ...on borne à 0.
				}                                                   // Fin bornage.
				fmt.Printf("🔫 Fusil → %d dégâts | %s PV: %d/%d\n",  // Affiche le résultat du tir.
					d, g.Name, g.HP, g.MaxHP)
			},                                                      // Fin action Fusil.
		})                                                          // Fin de l'append d'une option.
	} // Fin du test containsSkill(Fusil).
	if containsSkill(c, "mp5") {                                    // Si le perso possède "mp5" (attention à la casse)...
		options = append(options, option{                           // ...on ajoute l'option correspondante.
			label: "tir mp5 (14 dégâts)",                           // Libellé (note : casse et valeur d'affichage).
			do: func() {                                            // Action : tir mp5.
				d := 16                                             // Dégâts appliqués (ici 16).
				g.HP -= d                                           // On enlève d PV au monstre.
				if g.HP < 0 {                                       // Si PV < 0...
					g.HP = 0                                        // ...borne à 0.
				}                                                   // Fin bornage.
				fmt.Printf("🔫 AK47 → %d dégâts | %s PV: %d/%d\n",  // Message (NB: texte affiche "AK47" → incohérent avec mp5).
					d, g.Name, g.HP, g.MaxHP)
			},                                                      // Fin action mp5.
		})                                                          // Fin append option mp5.
	} // Fin du test containsSkill(mp5).
	if containsSkill(c, "ak45") {                                   // Si le perso possède "ak45"...
		options = append(options, option{                           // ...on ajoute l'option correspondante.
			label: "ak 45 (16dégâts)",                              // Libellé (note : orthographe/espaces).
			do: func() {                                            // Action : tir ak45.
				d := 14                                             // Dégâts appliqués (ici 14).
				g.HP -= d                                           // Enlève d PV au monstre.
				if g.HP < 0 {                                       // Si PV < 0...
					g.HP = 0                                        // ...borne à 0.
				}                                                   // Fin bornage.
				fmt.Printf("🔫 MP5 → %d dégâts | %s PV: %d/%d\n",   // Message (NB: texte affiche "MP5" → incohérent avec ak45).
					d, g.Name, g.HP, g.MaxHP)
			},                                                      // Fin action ak45.
		})                                                          // Fin append option ak45.
	} // Fin du test containsSkill(ak45).

	// Affichage du menu (centré + coloré)
	for i, op := range options {                                    // Parcourt toutes les options pour les afficher numérotées.
		fmt.Println(centerLine(color(colRed,                        // Affiche chaque option en rouge et centrée.
			fmt.Sprintf("%d) %s", i+1, op.label))))
	} // Fin de la boucle d’affichage du menu.
	choice := AskInt("→ Choix: ", 1, len(options))                  // Demande un choix valide (1..nombre d'options).

	// Exécute l’action choisie
	options[choice-1].do()                                          // Appelle la fonction de l’option sélectionnée (index 0-based).
} // Fin de charTurn.

// =========================================
// 22.1 / 22.2) Combat d'entraînement complet
// =========================================
// - Affiche les règles au début pour guider le joueur.
// - Montre le numéro du tour et le HUD (PV).
// - Respecte l’initiative : celui qui a l’initiative la plus haute commence.
// - Fin de combat : Victoire (gain d’XP) ou “Fin de l’entraînement”.
// - Retour au menu (laisser l’utilisateur reprendre la main).
func trainingFight(c *Character) {                                  // Lance un combat d'entraînement contre un gobelin.
	g := initGoblin()                                                // Crée le monstre d'entraînement.
	turn := 1                                                        // Initialise le compteur de tours à 1.
	playerTurnFirst := c.Initiative >= g.Initiative                  // Détermine qui joue en premier (initiative).

	bannerCombat()                                                   // Affiche une bannière ASCII/visuelle de combat.
	fmt.Println(centerLine(bold(color(colYellow,                     // Affiche un titre "Règles du combat d'entraînement".
		"Règles du combat d'entraînement"))))
	fmt.Println(centerLine(color(colYellow,                          // Règle 1 : ordre des tours.
		"• Ordre des tours selon l'Initiative.")))
	fmt.Println(centerLine(color(colYellow,                          // Règle 2 : dégâts des attaques.
		"• Attaque basique = 5 dégâts. Armes = dégâts supérieurs si débloquées.")))
	fmt.Println(centerLine(color(colYellow,                          // Règle 3 : coup fort ennemi tous les 3 tours.
		"• Tous les 3 tours (3, 6, 9...), l'ennemi frappe x2. Pense aux potions !")))

	for c.HP > 0 && g.HP > 0 {                                       // Boucle principale du combat tant que les deux sont vivants.
		fmt.Printf("\n— Tour %d —\n", turn)                          // Affiche le numéro du tour avec une ligne vide avant.
		printHUD(c, &g)                                              // Affiche le HUD (PV joueur/monstre).

		if playerTurnFirst {                                         // Si le joueur commence ce tour...
			// Joueur → Ennemi
			charTurn(c, &g)                                          // Exécute le tour du joueur.
			if g.HP <= 0 {                                           // Si le monstre est mort après l'action du joueur...
				break                                                // ...on sort de la boucle (victoire).
			}                                                        // Fin test mort monstre.
			goblinPattern(g, c, turn)                                // Puis l'ennemi joue selon son pattern.
		} else {                                                     // Sinon, l'ennemi joue d'abord...
			// Ennemi → Joueur
			goblinPattern(g, c, turn)                                // L'ennemi attaque selon son pattern.
			if c.HP <= 0 {                                           // Si le joueur meurt...
				break                                                // ...on sort (défaite).
			}                                                        // Fin test mort joueur.
			charTurn(c, &g)                                          // Puis c'est le tour du joueur.
		}                                                            // Fin alternance des tours.
		turn++                                                       // Incrémente le compteur de tours.
	} // Fin de la boucle de combat.

	// Écran de fin
	if g.HP <= 0 {                                                   // Si le monstre n'a plus de PV → victoire.
		fmt.Println(centerLine(color(colGreen, "🏆 Victoire !")))    // Affiche un message de victoire en vert.
		c.Experience += g.ExpReward                                  // Ajoute l'XP du monstre au personnage.
	} else {                                                         // Sinon, le combat s'est terminé autrement (défaite/arrêt).
		fmt.Println(centerLine(color(colYellow, "Fin de l'entrainement."))) // Message de fin neutre.
	}                                                                // Fin du if/else de fin de combat.
	fmt.Println(centerLine(color(colCyan, "Retour au menu...")))     // Invite visuellement à revenir au menu.
} // Fin de trainingFight.

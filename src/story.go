package main // Le fichier appartient au package exécutable "main".

import "fmt" // On importe fmt pour utiliser Println/Printf, etc.

// =================
// story.go : Histoire
// =================
// Adaptation fidèle au récit, avec ASCII + effets "slow". // Simple description du fichier.

func runStory(c *Character) { // Fonction principale de l'histoire, prend le perso joueur en paramètre (pointeur).
	clear()                                             // Nettoie l'écran/terminal avant d'afficher la scène.
	slow("...\n", 25)                                   // Affiche "..." lentement (effet dramatique).
	slow("Tu ouvres les yeux. Tu es assis à un bureau. Il fait nuit.\n", 18) // Texte narratif affiché lentement.
	slow("Tu es... dans ton lycée ? Quelque chose cloche.\n", 18)            // Suite de la narration.
	slow("Le néon clignote. Il y a du sang sur le sol.\n", 18)               // Détail d'ambiance.
	slow("Des cris lointains, comme sous l'eau.\n", 18)                      // Autre détail d'ambiance.
	slow("Une voix résonne : « QUI ES-TU ? »\n", 20)                         // Question clé affichée lentement.

	fmt.Println(centerLine(color(colRed, "[Choisissez votre genre] 1) Homme  2) Femme"))) // Affiche une ligne centrée et colorée en rouge.
	g := AskInt("Choix: ", 1, 2)                                                          // Demande un entier entre 1 et 2 et le stocke dans g.
	if g == 1 {                                                                           // Si le joueur choisit 1...
		fmt.Println(centerLine(color(colRed, "[ERREUR : FAUX.]  Recalcul en cours…")))     // Affiche un message d'erreur scénarisé.
		_ = AskInt("Rechoisis (1) Homme / (2) Femme: ", 1, 2)                              // Redemande un choix (on ignore la valeur avec _).
		fmt.Println(centerLine(color(colRed, "[ACCEPTÉ.] Traitement de l'identité en cours..."))) // Confirme l'acceptation après recalcule.
	}                                                                                     // Fin du if.
	slow("« TU N'ES PAS QUI TU PENSES ÊTRE. »\n", 20)                                      // Avance dans la narration.
	slow("« TU NE T'ES JAMAIS RÉVEILLÉ(E). »\n", 20)                                       // Renforce le mystère.
	slow("Casier. Numéro 666. Fermé. Quelque chose t'observe depuis l'intérieur.\n", 20)   // Détail horrifique.

	// NIVEAU 2 — Marchand // Commentaire logique pour structurer les scènes.
	clear()                                                              // Nettoie l'écran avant la nouvelle scène.
	title(color(colReset, "L'ORDINATEUR DU MARCHAND"))                   // Affiche un titre (sans couleur) avec une fonction de titre.
	slow("(Tu entres dans une salle plongée dans le noir. Un ordinateur s'allume tout seul.)\n", 18) // Description de scène.
	slow("L'écran grésille. Une vidéo apparaît, avec une connexion instable.\n", 18)                  // Ambiance vidéo.
	slow("L'image saute, la voix se coupe par moments.\n", 18)                                       // Ambiance parasite.
	slow("Vidéo — Homme inconnu (pixellisé, paniqué) :\n", 18)                                       // Introduction du PNJ en vidéo.
	slow("« Si tu vois ça… tu dois m'écouter… vite. »\n", 18)                                        // Réplique du PNJ.
	slow("[L'image coupe, puis revient]\n", 18)                                                      // Effet coupure.
	slow("« Va au distributeur. Prends une potion… ou… un poison… choisis bien. »\n", 18)            // Consigne du PNJ.

	showMerchant(c) // Ouvre l'interface/mécanique du marchand en utilisant le perso (achat/choix).

	// NIVEAU 3 — Combat zombie → Fusil + contamination // Nouvelle section.
	clear()                                                     // Nettoyage d'écran.
	title("NIVEAU 3 – COMBAT")                                  // Affiche le titre du niveau/combat.
	slow("Un zombie du lycée avance. Badge : ton prof de sport… mais plus humain.\n", 18) // Intro de l'ennemi.
	bannerCombat()                                              // Affiche une bannière ASCII/mise en forme de combat.
	zombie := Monster{Name: "Zombie du lycée", MaxHP: 50, HP: 50, Attack: 6, Initiative: 7, ExpReward: 35} // Instancie l'ennemi.
	fight(c, zombie, true) // contamination même en cas de victoire // Lance le combat (true = applique poison narratif si victoire).
	fmt.Println(color(colYellow, "[Fusil à pompes débloqué]"))  // Notifie le déblocage d'une arme.
	addSkill(c, "Fusil")                                        // Ajoute la compétence/arme "Fusil" au perso.

	slow("Le zombie tombe. « Je m’appelle Lucas. »\n", 18)      // Narration post-combat.
	slow("« Coincé ici, comme toi. On devrait faire équipe. »\n", 18) // Dialogue d'allié potentiel.
	fmt.Println(color(colCyan, "Vos forces se combinent : PV doublés.")) // Indique l'effet de groupe.
	c.MaxHP *= 2                                                // Double les PV max du perso.
	c.HP = c.MaxHP                                              // Soigne le perso au nouveau max.

	// NIVEAU 4 — Prof zombifié → AK47 + contamination // Nouvelle section combat.
	clear()                                                      // Nettoyage d'écran.
	title("NIVEAU 4 – COMBAT EMPOISONNÉ")                       // Titre du combat.
	slow("Prof : « Toujours… en retard… toujours toi ! »\n", 18) // Petite réplique du boss.
	prof := Monster{Name: "Professeur zombifié", MaxHP: 65, HP: 65, Attack: 8, Initiative: 8, ExpReward: 45} // Crée le boss.
	fight(c, prof, true)                                         // Combat contre le prof (contamination à la victoire).
	fmt.Println(color(colYellow, "[AK47 débloqué]"))             // Notifie l'arme débloquée.
	addSkill(c, "AK47")                                          // Ajoute l'AK47 au perso.

	// NIVEAU 5 — Doubles zombifiés // Nouvelle section.
	clear()                                                      // Nettoyage d'écran.
	title("NIVEAU 5 – DOUBLES ZOMBIFIÉS")                        // Titre de la scène.
	slow("Les doubles imitent tes attaques une fois sur deux.\n", 18) // Règle spéciale de l'ennemi (info au joueur).
	doubles := Monster{Name: "Doubles zombifiés", MaxHP: 80, HP: 80, Attack: 9, Initiative: 9, ExpReward: 60} // Crée l'ennemi.
	fight(c, doubles, false)                                     // Lance le combat (pas de contamination).

	// NIVEAU 6 — Remède (Violet = Rouge + Bleu) // Énigme alchimie.
	clear()                                                      // Nettoyage d'écran.
	title("NIVEAU 6 – REMÈDE")                                   // Titre de la scène d'énigme.
	slow("Laboratoire. Quatre fioles : 1) Rouge  2) Bleu  3) Vert  4) Jaune.\n", 18) // Présente les options.
	slow("Indice : VIOLET = Rouge + Bleu.\n", 18)                                     // Donne l'indice de combinaison.
	for {                                                        // Boucle jusqu'à ce que le joueur trouve la bonne combinaison.
		a := AskInt("Première fiole: ", 1, 4)                   // Demande le premier choix (1 à 4).
		b := AskInt("Deuxième fiole: ", 1, 4)                   // Demande le second choix (1 à 4).
		if (a == 1 && b == 2) || (a == 2 && b == 1) {          // Vérifie si la combinaison est Rouge + Bleu (dans n'importe quel ordre).
			fmt.Println(color(colGreen, "[REMÈDE RÉCUPÉRÉ – Poison neutralisé]")) // Message de succès.
			c.HP = c.MaxHP                                        // Soigne entièrement le perso.
			break                                                 // Sort de la boucle (énigme résolue).
		}                                                        // Fin du if de réussite.
		fmt.Println(color(colRed, "Mauvaise fiole... tu t'effondres (on te relève pour la démo).")) // Échec → message.
		c.HP = c.MaxHP / 2                                       // Pénalité : PV réduits de moitié.
	}                                                            // Fin de la boucle; elle recommence si mauvaise combinaison.
	fmt.Println(color(colYellow, "[MP5 débloqué]"))              // Notifie l'arme MP5 débloquée.
	addSkill(c, "MP5")                                           // Ajoute la MP5 au perso.

	// NIVEAU 7 — Gardien des Âmes // Boss final.
	clear()                                                      // Nettoyage d'écran.
	title("NIVEAU 7 – COMBAT FINAL")                             // Titre du combat final.
	slow("Gymnase. Le Gardien des Âmes, fusion de profs et du directeur.\n", 18) // Intro du boss final.
	gardien := Monster{Name: "Gardien des Âmes", MaxHP: 120, HP: 120, Attack: 11, Initiative: 10, ExpReward: 120} // Crée le boss final.
	fight(c, gardien, false)                                     // Combat final (pas de contamination).
	slow("Le Gardien se dissout en cendres. [ACCÈS AUTORISÉ – SORTIE DÉVERROUILLÉE]\n", 16) // Narration de victoire.
	slow("Lucas disparaît. Le casier est vide.\n", 16)                                   // Épilogue.
	slow("Voix : « TU NE T’ES JAMAIS RÉVEILLÉ(E). »\n\n", 16)                             // Dernière réplique mystérieuse.
} // Fin de runStory.

// fight pour l'histoire (réutilise le pattern & charTurn) // Commentaire de doc.
func fight(c *Character, m Monster, contaminateOnWin bool) { // Fonction de combat générique; bool = applique poison après victoire.
	g := m                                           // Copie du monstre (travaille sur une valeur locale).
	turn := 1                                        // Compteur de tour qui commence à 1.
	playerStarts := c.Initiative >= g.Initiative     // Détermine qui joue en premier selon l'initiative.
	bannerCombat()                                   // Affiche la bannière de combat.
	for c.HP > 0 && g.HP > 0 {                       // Boucle tant que les deux sont vivants.
		fmt.Printf("\n— Tour %d —\n", turn)           // Affiche le numéro du tour.
		if playerStarts {                             // Si le joueur commence...
			charTurn(c, &g)                           // Tour du joueur (attaque/compétence sur le monstre).
			if g.HP <= 0 {                            // Si le monstre est mort après l'action du joueur...
				break                                  // On sort de la boucle (victoire).
			}                                         // Fin du check mort monstre.
			goblinPattern(g, c, turn)                 // Tour du monstre (pattern d'attaque en fonction du tour).
		} else {                                      // Sinon le monstre commence...
			goblinPattern(g, c, turn)                 // Monstre attaque d'abord.
			if c.HP <= 0 {                            // Si le joueur meurt...
				break                                  // Sort de la boucle (défaite).
			}                                         // Fin du check mort joueur.
			charTurn(c, &g)                           // Puis le tour du joueur.
		}                                             // Fin de l'alternance des tours.
		turn++                                        // Incrémente le numéro de tour.
	}                                                 // Fin de la boucle de combat.
	if g.HP <= 0 {                                    // Si le monstre est mort → victoire du joueur.
		fmt.Println(color(colGreen, "🏆 Victoire !")) // Message de victoire coloré.
		c.Experience += g.ExpReward                   // Ajoute l'XP remportée au joueur.
		if contaminateOnWin {                         // Si on doit appliquer un poison narratif après victoire...
			// "poison narratif" simple               // Commentaire: 3 ticks de poison qui retirent 10 PV chacun.
			for i := 1; i <= 3; i++ {                 // Répète 3 fois.
				c.HP -= 10                            // Enlève 10 PV.
				if c.HP < 0 {                         // Si PV < 0...
					c.HP = 0                           // Clamp à 0 pour éviter des PV négatifs.
				}                                     // Fin clamp.
				fmt.Printf(color(colMagenta, "☠️ Poison %d/3 → PV: %d/%d\n"), i, c.HP, c.MaxHP) // Affiche l'état après chaque tick.
				if isDead(c) {                        // Si le joueur meurt du poison...
					break                              // On arrête d'appliquer le poison.
				}                                     // Fin check mort.
			}                                         // Fin de la boucle poison.
		}                                             // Fin du if contaminateOnWin.
	} else {                                          // Sinon, le joueur a perdu ou fuite/fin anormale.
		fmt.Println(color(colYellow, "Fin du combat.")) // Message neutre de fin de combat.
	}                                                 // Fin du else.
} // Fin de fight.

package main // Le fichier appartient au package exécutable "main".

import (          // Début du bloc d'import.
	"fmt"       // fmt : affichage formaté (Println, Printf, ...).
	"regexp"    // regexp : expressions régulières (ici pour enlever les codes ANSI).
	"strings"  // strings : fonctions utilitaires sur les chaînes.
	"time"     // time : gestion du temps (Sleep, etc.).
) // Fin du bloc d'import.

// ==================
// main.go : Entrée
// ==================
// Intro + création personnage (PDF Tâches 1–2), puis choix du mode:
// Histoire (ton scénario 1→7) ou Libre (menus et entraînement).

func characterCreation() Character {                                      // Fonction qui crée et retourne un personnage complet.
	// nom lettres uniquement (Tâche 11)
	var name string                                                        // Déclare une variable pour stocker le nom.
	for {                                                                  // Boucle jusqu'à obtenir un nom valide.
		name = AskLine(color(colReset, ("Ton nom (lettres uniquement) : "))) // Demande une ligne à l'utilisateur (stylée avec couleur).
		if onlyLetters(name) {                                              // Vérifie que le nom contient seulement des lettres.
			break                                                           // Si OK, on sort de la boucle.
		}                                                                   // Fin du test.
		fmt.Println(color(colRed, "⛔ Lettres uniquement."))                // Message d'erreur si non conforme.
	}                                                                       // Fin de la boucle de saisie du nom.
	name = formatName(name)                                                 // Formate le nom (ex: capitaliser, nettoyer).

	fmt.Println(centerLine(color(colRed, ("Classe : 1) Humain  2) Elfe  3) Nain")))) // Affiche le choix de classe, centré et en rouge.
	cl := AskInt("Choix: ", 1, 3)                                                    // Demande un entier entre 1 et 3 pour la classe.

	class := "Humain"                                                    // Valeur par défaut : Humain.
	base := 100                                                          // PV de base par défaut : 100.
	if cl == 2 {                                                         // Si l'utilisateur a choisi 2...
		class = "Elfe"                                                   // La classe devient Elfe.
		base = 80                                                        // PV de base de l'Elfe.
	}                                                                    // Fin du if Elfe.
	if cl == 3 {                                                         // Si l'utilisateur a choisi 3...
		class = "Nain"                                                   // La classe devient Nain.
		base = 120                                                       // PV de base du Nain.
	}                                                                    // Fin du if Nain.

	// initCharacter (Tâche 2) — valeurs simples et lisibles
	c := Character{                                                      // Initialise la structure Character avec des valeurs de départ.
		Name:          name,                                              // Nom du personnage (saisi/formatté).
		Class:         class,                                             // Classe choisie (Humain/Elfe/Nain).
		Level:         1,                                                 // Niveau initial.
		MaxHP:         base,                                              // PV max selon la classe.
		HP:            base / 2, // commence à 50% pour montrer les potions // PV actuels à 50% (pour inciter à soigner).
		MaxMana:       50,                                                // Mana max de départ.
		Mana:          50,                                                // Mana actuel.
		Gold:          100, // Tâche 13 fu pdf()                           // Or de départ.
		Experience:    0,                                                 // XP actuelle.
		ExperienceMax: 50,                                                // Seuil d'XP pour le prochain niveau.
		Inventory:     []Item{},                                          // Inventaire vide au départ.
		InventoryCap:  10, // Tâche 12 du pdf ()                           // Capacité d'inventaire de départ.
		Equip:         Equipment{},                                       // Emplacements d'équipement vides.
		Skills:        []string{"Coup de poing"}, // Tâche 10 (base)       // Compétence de base.
		Initiative:    10,                                                // Initiative de départ (ordre des tours).
	}                                                                     // Fin de l'initialisation du perso.
	return c                                                              // Retourne le personnage créé.
} // Fin de characterCreation.

func showInfos(c *Character) {                                           // Affiche un résumé des infos du perso.
	fmt.Printf("\n=== %s (%s) ===\n", c.Name, c.Class)                   // Titre avec Nom et Classe.
	fmt.Printf("Niveau: %d | PV: %d/%d | Mana: %d/%d | Or: %d | XP: %d/%d\n", // Stats principales (niveau, PV, mana, or, XP).
		c.Level, c.HP, c.MaxHP, c.Mana, c.MaxMana, c.Gold, c.Experience, c.ExperienceMax)
	fmt.Printf("Équipement: Tête[%s] Torse[%s] Pieds[%s]\n", c.Equip.Head, c.Equip.Body, c.Equip.Feet) // Slots d'équipement affichés.
	fmt.Printf("Compétences: %v\n", c.Skills)                             // Liste des compétences.
	fmt.Printf("Inventaire (%d/%d) items\n", len(c.Inventory), c.InventoryCap) // Nombre d'items et capacité.
} // Fin de showInfos.

func runMenus(c *Character) {                                            // Menu principal du mode "Libre".
	for {                                                                // Boucle du menu jusqu'à quitter.
		fmt.Println(centerLine("\n=== MENU PRINCIPAL ==="))              // Affiche le titre du menu, centré.
		fmt.Println(centerLine("1) Afficher les infos"))                 // Option 1 : infos perso.
		fmt.Println(centerLine("2) Accéder à l'inventaire"))             // Option 2 : inventaire.
		fmt.Println(centerLine("3) Marchand"))                           // Option 3 : marchand.
		fmt.Println(centerLine("4) Forgeron"))                           // Option 4 : forgeron.
		fmt.Println(centerLine("5) Entrainement"))                       // Option 5 : combat d’entraînement.
		fmt.Println(centerLine("6) Quitter"))                            // Option 6 : quitter le menu libre.

		choice := AskInt("Choix: ", 1, 6)                                // Demande un choix entre 1 et 6.
		switch choice {                                                   // Redirige vers la bonne action.
		case 1:                                                           // Si 1...
			showInfos(c)                                                  // ...affiche les infos.
		case 2:                                                           // Si 2...
			accessInventory(c)                                            // ...ouvre l'inventaire.
		case 3:                                                           // Si 3...
			showMerchant(c)                                               // ...ouvre le marchand.
		case 4:                                                           // Si 4...
			showForge(c)                                                  // ...ouvre le forgeron.
		case 5:                                                           // Si 5...
			trainingFight(c)                                              // ...lance le combat d'entraînement.
		case 6:                                                           // Si 6...
			return                                                        // ...sort de runMenus (retour au caller).
		}                                                                 // Fin du switch.
	}                                                                     // Fin de la boucle du menu.
} // Fin de runMenus.

var screenWidth = 120 // Largeur "virtuelle" de l'écran pour centrer les textes.


// retire les codes couleur ANSI pour compter la vraie longueur visible
func stripANSI(s string) string {                                        // Supprime les codes ANSI d'une chaîne.
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)                            // Compile une regex qui matche les codes de couleur ANSI.
	return re.ReplaceAllString(s, "")                                     // Remplace ces codes par une chaîne vide.
} // Fin de stripANSI.

func centerLine(s string) string {                                       // Centre une ligne visuellement selon screenWidth.
	plain := stripANSI(s)                                                 // Enlève les codes couleur pour mesurer la vraie longueur.
	if len(plain) >= screenWidth {                                        // Si la ligne est déjà plus large que la largeur cible...
		return s                                                          // ...on la retourne telle quelle (pas de centrage).
	}                                                                     // Fin du test de longueur.
	pad := (screenWidth - len(plain)) / 2                                 // Calcule le nombre d'espaces à ajouter à gauche.
	return strings.Repeat(" ", pad) + s                                   // Renvoie la ligne préfixée d'espaces.
} // Fin de centerLine.

func centerBlock(block string) string {                                   // Centre chaque ligne d'un bloc multi-lignes.
	lines := strings.Split(strings.TrimRight(block, "\n"), "\n")           // Découpe le bloc en lignes (sans le dernier \n).
	for i, ln := range lines {                                             // Parcourt chaque ligne...
		lines[i] = centerLine(ln)                                          // ...et la centre.
	}                                                                      // Fin de la boucle.
	return strings.Join(lines, "\n")                                       // Reconstruit le bloc centré.
} // Fin de centerBlock.

// Appelle printBigTitle() là où tu veux afficher le logo (intro, etc.)
func printBigTitle() {                                                    // Affiche un gros logo ASCII en couleur et gras.
	art := `                                                               // Déclare une string multi-lignes (backticks) contenant l'art ASCII.
   ____        ___ _       _       ____       _        _   _
 |  _ \  __ _| _ _ | | __  |  _ \  ___| |_ __ _| |_(_) ___  _ __
 | | | |/ _' | '_ \| |/ /  | | | |/ _ \ __/ _' | __| |/ _ \| '_ \
 | |_| | (_| | | | |   <   | |_| |  __/ || (_| | |_| | (_) | | | |
 |____/ \__,_|_| |_|_|\_\  |____/ \___|\__\__,_|\__|_|\___/|_| |_|
 												
` // Fin du littéral ASCII (note: les backticks gardent la mise en forme).
	// remet les vrais backticks, puis applique couleur + gras
	art = strings.ReplaceAll(art, "[BT]", "`")                            // Remplace [BT] par un backtick (utile si tu mets [BT] dans l'art).
	fmt.Print(centerBlock(bold(color(colRed, art))))                      // Affiche l'art centré, en rouge et en gras.
} // Fin de printBigTitle.

func main() {                                                             // Point d'entrée du programme.
	clear()                                                                // Nettoie l'écran au lancement.
	printBigTitle()                                                        // Affiche le gros titre/logo.

	time.Sleep(1 * time.Second)                                            // Petite pause d'une seconde (effet dramatique).
	fmt.Println(color(colRed, ("Une voix résonne : « QUI ES-TU ? »")))    // Affiche une réplique stylée en rouge.
	time.Sleep(1 * time.Second)                                            // Re-pause d'une seconde.

	// Création du personnage (Tâche 11 appliquée)
	player := characterCreation()                                          // Lance la création du personnage et stocke le résultat.

	// Deux modes : Histoire (1→7) ou Libre
	fmt.Println(centerLine(color(colRed, ("Modes : 1) Histoire (Niveaux 1→7)")))) // Affiche l'offre de modes (texte centré rouge).
	mode := AskInt("Choix: ", 1, 2)                                               // Demande un choix de mode entre 1 et 2.
	if mode == 1 {                                                                // Si l'utilisateur choisit 1...
		runStory(&player)                                                         // ...on lance le mode Histoire (scénario).
	} else {                                                                      // Sinon (mode 2)...
		runMenus(&player)                                                         // ...on lance le mode Libre (menus).
	}                                                                             // Fin du choix de mode.
	fmt.Println(color(colCyan, "Merci d'avoir joué !"))                           // Message de fin en cyan.
} // Fin de main.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Petit jeu texte inspiré du scénario fourni.
// Pour lancer : go run aventure_lycee.go

type Player struct {
	Name        string
	Gender      string
	HP          int
	MaxHP       int
	PoisonLevel int // 0..100
	HasLucas    bool
	Weapons     map[string]bool
	Alive       bool
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	clear()
	fmt.Println("Niveau 1 - Le réveil\n")
	fmt.Println("[ Écran noir... ]\n...\nTu ouvres les yeux.")
	fmt.Println("Tu es assis à un bureau. Il fait nuit.\nTu es... dans ton lycée ?\nQuelque chose cloche.")
	fmt.Println("Le néon au plafond clignote. Il y a du sang sur le sol.")
	fmt.Println("Une voix résonne dans ta tête.")
	fmt.Println("\n« QUI ES-TU ? »\n")

	fmt.Print("[Choisissez votre nom] > ")
	name := readLine(reader)
	if name == "" {
		name = "Prénom"
	}

	fmt.Printf("\nBien %s\n\n", name)
	fmt.Println("[Choisissez votre genre]")
	fmt.Println("1. Homme")
	fmt.Println("2. Femme")
	fmt.Print("\n> ")
	choice := readLine(reader)

	// En accord avec le scénario : si l'utilisateur choisit homme => erreur et recalcul
	gender := "Femme"
	if choice == "1" {
		fmt.Println("\n[ERREUR : FAUX.]")
		fmt.Println("> Recalcul en cours…\n")
		fmt.Println("« TU N’ES PAS QUI TU PENSES ÊTRE. »\n")
		fmt.Println("1. Femme")
		fmt.Print("> ")
		_ = readLine(reader)
		fmt.Println("\n[ACCEPTÉ.]")
	}

	fmt.Println("\n« TU NE T’ES JAMAIS RÉVEILLÉE. »\n")

	p := Player{
		Name:        name,
		Gender:      gender,
		HP:          100,
		MaxHP:       100,
		PoisonLevel: 0,
		HasLucas:    false,
		Weapons:     map[string]bool{"Ciseaux": true},
		Alive:       true,
	}

	pause(reader)

	// Niveau 2
	niveau2(reader, &p)
	if !p.Alive { return }

	// Niveau 3 combat zombie
	niveau3(reader, &p)
	if !p.Alive { return }

	// Niveau 4 combat empoisonné
	niveau4(reader, &p)
	if !p.Alive { return }

	// Niveau 5 doubles zombifiés
	niveau5(reader, &p)
	if !p.Alive { return }

	// Niveau 6 remède
	niveau6(reader, &p)
	if !p.Alive { return }

	// Niveau 7 combat final
	niveau7(reader, &p)
}

func niveau2(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 2 - L’Ordinateur du Marchand\n")
	fmt.Println("Tu entres dans une salle plongée dans le noir. Un ordinateur s’allume tout seul.")
	fmt.Println("L’écran grésille. Une vidéo apparaît, avec une connexion instable.")
	fmt.Println("\nVidéo – Homme inconnu (pixellisé, paniqué) :")
	fmt.Println("« …Si tu vois ça… tu dois m’écouter… vite. »")
	fmt.Println("…va au distributeur. Prends une potion… ou… un poison… choisis bien.")
	fmt.Println("\nL'image saccade, l'écran devient rouge. Un hurlement... puis tout s'éteint.")
	pause(r)
}

func niveau3(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 3 – Le combat\n")
	fmt.Println("Derrière les casiers, un zombie du lycée avance. Sa bouche est couverte de sang.")
	fmt.Println("Son badge porte le nom de ton prof de sport… mais il n’est plus humain.")

	combat := simpleCombat{EnemyName: "Zombie du lycée", EnemyHP: 40}
	result := runCombat(r, p, &combat)
	if !result { p.Alive = false; fmt.Println("Tu as été tuée par le zombie. Recommence le niveau...") ; return }

	fmt.Println("\nLe zombie tombe. Tu vois une personne qui se cache derrière un casier.")
	fmt.Println("Après un instant, elle sort et s’avance vers toi.")
	fmt.Println("\n« Je m'appelle Lucas. »")
	fmt.Println("Il te regarde avec sérieux.")
	fmt.Println("« Je suis coincé ici, comme toi. On devrait faire équipe. »")

	p.HasLucas = true
	p.MaxHP *= 2
	p.HP = p.MaxHP
	fmt.Printf("\nVos forces se combinent : Points de vie doublés → %d HP\n", p.MaxHP)
	p.Weapons["Fusil à pompes"] = true
	pause(r)
}

func niveau4(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 4 - combat empoisoné\n")
	fmt.Println("Dans un couloir, un professeur zombifié t’attend. Il parle comme si tu étais un élève connu.")
	fmt.Println("Professeur :\n\n« Toujours… en retard… toujours toi ! Tu ne changes jamais. »\n")
	fmt.Println("Lucas (tremblant) : « …Il te connaît. Mais il ne devrait pas… personne ne devrait. »")

	combat := simpleCombat{EnemyName: "Professeur zombifié", EnemyHP: 60}
	result := runCombat(r, p, &combat)
	if !result { p.Alive = false; fmt.Println("Tu as été vaincue...") ; return }

	// même si victoire, joueur contaminé
	p.PoisonLevel += 30
	if p.PoisonLevel > 100 { p.PoisonLevel = 100 }
	fmt.Println("\nTu sens ton sang bouillir.")
	fmt.Println("[Tu es empoisonnée]")
	p.Weapons["AK47"] = true
	pause(r)
}

func niveau5(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 5 - [COMBAT ENGAGÉ – Ennemis : Doubles Zombifiés]\n")
	fmt.Println("Des doubles imitent tes attaques une fois sur deux. Chaque attaque ratée augmente le niveau de poison de 10%.")

	enemyHP := 80
	for enemyHP > 0 && p.Alive {
		fmt.Printf("\nTon HP: %d/%d  |  Poison: %d%%\n", p.HP, p.MaxHP, p.PoisonLevel)
		fmt.Println("Choisis ton attaque :")
		weapons := availableWeapons(p)
		for i, w := range weapons { fmt.Printf("%d. %s\n", i+1, w) }
		fmt.Print("> ")
		sel := readLine(r)
		idx, _ := strconv.Atoi(sel)
		if idx < 1 || idx > len(weapons) { fmt.Println("Choix invalide, attaque ratée.") ; p.PoisonLevel += 10 ; checkPoisonDeath(p); continue }
		attack := weapons[idx-1]

		// attaque réussit 70% sauf si double imite
		if rand.Intn(100) < 70 {
			// réussite
			dmg := rand.Intn(20) + 10
			enemyHP -= dmg
			fmt.Printf("Tu utilises %s et infliges %d dégâts. Ennemi HP restant: %d\n", attack, dmg, max(0, enemyHP))
			// imitation chance 50% when enemy still alive
			if enemyHP > 0 && rand.Intn(2) == 0 {
				// double attaque imite → parfois touche toi
				if rand.Intn(100) < 50 {
					dmg2 := rand.Intn(15) + 5
					p.HP -= dmg2
					fmt.Printf("Le double imite ton attaque et te blesse de %d.\n", dmg2)
				}
			}
		} else {
			// échec
			fmt.Println("Ton attaque rate !")
			p.PoisonLevel += 10
			checkPoisonDeath(p)
		}

		// ennemi attaque
		if enemyHP > 0 {
			enemyDmg := rand.Intn(15) + 5
			p.HP -= enemyDmg
			fmt.Printf("Le double te contre-attaque et inflige %d dégâts.\n", enemyDmg)
		}

		if p.HP <= 0 { p.Alive = false ; fmt.Println("Tu t'es effondrée...") ; return }
		if p.PoisonLevel > 0 { // poison passif dégâts
			poisonD := p.PoisonLevel / 10
			p.HP -= poisonD
			fmt.Printf("Le poison te ronge (%d dégâts).\n", poisonD)
			checkPoisonDeath(p)
		}
	}

	if enemyHP <= 0 {
		fmt.Println("\nLes doubles sont vaincus.\n")
		p.Weapons["MP5"] = true
		pause(r)
	}
}

func niveau6(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 6 - Remède\n")
	fmt.Println("Tu arrives dans un petit laboratoire abandonné.")
	fmt.Println("Au centre, quatre fioles sont alignées : rouge, bleu, vert et jaune.")
	fmt.Println("Une note sur le pupitre indique :\n\"Une seule potion guérira le poison. Choisis avec soin.\"")
	fmt.Println("\nLucas te regarde :\n« Fais le bon choix… je ne sais pas laquelle est la bonne. »\n")

	fmt.Println("Choix des fioles")
	fmt.Println("1. Rouge")
	fmt.Println("2. Bleu")
	fmt.Println("3. Vert")
	fmt.Println("4. Jaune")
	fmt.Print("> ")
	choice := readLine(r)

	// bonne combinaison : rouge + bleu -> mais player chooses a single color in this simplified version
	// Nous interprétons : si joueur choisit Rouge ou Bleu => guérit; sinon mort.
	if choice == "1" || choice == "2" {
		p.PoisonLevel = 0
		fmt.Println("\n[REMÈDE RÉCUPÉRÉ – Poison neutralisé]")
		p.Weapons["MP5"] = true
		pause(r)
	} else {
		fmt.Println("\nLa fiole était toxique...")
		p.Alive = false
		fmt.Println("Tu meurs empoisonnée. Fin de la partie.")
		return
	}
}

func niveau7(r *bufio.Reader, p *Player) {
	clear()
	fmt.Println("Niveau 7 - Combat final\n")
	fmt.Println("Tu entres dans le gymnase principal. Au centre, une silhouette imposante :")
	fmt.Println("Le Gardien des Âmes, une fusion de plusieurs professeurs et du directeur.")
	fmt.Println("Le visage à moitié humain, à moitié zombie avec des membres disproportionnés.")

	combat := simpleCombat{EnemyName: "Gardien des Âmes", EnemyHP: 200}
	result := runCombat(r, p, &combat)
	if !result { p.Alive = false ; fmt.Println("Le Gardien t'a anéantie.") ; return }

	fmt.Println("\nLe Gardien hurle et se dissout en cendres.")
	fmt.Println("Une lumière s’allume sur la porte principale :")
	fmt.Println("[ACCÈS AUTORISÉ – SORTIE DÉVERROUILLÉE]")
	fmt.Println("→ silence. Lucas disparaît. Le casier est vide.")
	fmt.Println("\nVoix finale :\n\n« TU NE T’ES JAMAIS RÉVEILLÉ(E). »")
	fmt.Println("\n(Écran noir.)")
}

// Combat général

type simpleCombat struct {
	EnemyName string
	EnemyHP   int
}

func runCombat(r *bufio.Reader, p *Player, c *simpleCombat) bool {
	enemyHP := c.EnemyHP
	fmt.Printf("[ COMBAT ENGAGÉ ]  Ennemi : %s\n", c.EnemyName)
	for enemyHP > 0 && p.Alive {
		fmt.Printf("\nTon HP: %d/%d  |  Poison: %d%%\n", p.HP, p.MaxHP, p.PoisonLevel)
		fmt.Println("Choisissez votre arme :")
		weapons := availableWeapons(p)
		for i, w := range weapons { fmt.Printf("%d. %s\n", i+1, w) }
		fmt.Print("> ")
		sel := readLine(r)
		idx, _ := strconv.Atoi(sel)
		if idx < 1 || idx > len(weapons) { fmt.Println("Choix invalide, tu perds ton tour.") ;
			// ennemi attaque
			dmg := rand.Intn(12) + 5
			p.HP -= dmg
			fmt.Printf("L'ennemi profite et te blesse de %d.\n", dmg)
			if p.PoisonLevel > 0 { poisonTick(p) }
			checkPoisonDeath(p)
			continue
		}
		weapon := weapons[idx-1]
		// dégâts selon arme
		dmg := weaponDamage(weapon)
		enemyHP -= dmg
		fmt.Printf("Tu utilises %s et infliges %d dégâts. Ennemi HP restant: %d\n", weapon, dmg, max(0, enemyHP))

		// chance d'impact de l'ennemi
		if enemyHP > 0 {
			enemyDmg := rand.Intn(20) + 5
			p.HP -= enemyDmg
			fmt.Printf("L'ennemi riposte et inflige %d dégâts.\n", enemyDmg)
		}

		if p.PoisonLevel > 0 { poisonTick(p) }
		checkPoisonDeath(p)
	}

	return enemyHP <= 0
}

func availableWeapons(p *Player) []string {
	list := []string{}
	if p.Weapons["Ciseaux"] { list = append(list, "Ciseaux") }
	if p.Weapons["Objet improvisé 80%"] { list = append(list, "Objet improvisé 80%") }
	// always let player choose potion/poison if poison level < 100
	list = append(list, "Potion/Poison")
	if p.Weapons["Fusil à pompes"] { list = append(list, "Fusil à pompes") }
	if p.Weapons["AK47"] { list = append(list, "AK47") }
	if p.Weapons["MP5"] { list = append(list, "MP5") }
	return list
}

func weaponDamage(w string) int {
	switch w {
	case "Ciseaux":
		return rand.Intn(12) + 4
	case "Objet improvisé 80%":
		return rand.Intn(18) + 6
	case "Potion/Poison":
		// effet aléatoire : peut soigner ou empoisonner l'ennemi (ou soi-même si mal utilisé)
		if rand.Intn(100) < 60 {
			// potion soin
			return rand.Intn(8) + 2
		}
		return rand.Intn(20) + 5
	case "Fusil à pompes":
		return rand.Intn(25) + 15
	case "AK47":
		return rand.Intn(22) + 12
	case "MP5":
		return rand.Intn(20) + 10
	default:
		return rand.Intn(10) + 3
	}
}

func poisonTick(p *Player) {
	// Poison inflige dégâts selon niveau
	d := max(1, p.PoisonLevel/10)
	p.HP -= d
	fmt.Printf("Le poison te fait perdre %d HP.\n", d)
}

func checkPoisonDeath(p *Player) {
	if p.PoisonLevel >= 100 {
		fmt.Println("Le poison atteint un seuil mortel.")
		p.Alive = false
	}
	if p.HP <= 0 {
		p.Alive = false
	}
}

func readLine(r *bufio.Reader) string {
	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

func pause(r *bufio.Reader) {
	fmt.Print("\n(Entrez pour continuer...) ")
	_ = readLine(r)
}

func clear() {
	// simple séparation visuelle
	fmt.Print("\n----------------------------------------\n")
}

func max(a, b int) int { if a>b { return a } ; return b }

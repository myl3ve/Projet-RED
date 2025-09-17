package character

import "fmt"

// Tâches 1–3 + 11 (création perso) :contentReference[oaicite:9]{index=9}
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name        string
	Class       string
	Level       int
	HPMax       int
	HP          int
	Inventory   []string
	Money       int
	Skills      []string
	Equipment   Equipment
	InventoryCap int

	// Missions : initiative, exp, mana (peuvent être 0 au départ)
	Initiative int
	Exp        int
	ExpMax     int
	Mana       int
	ManaMax    int
}

func initCharacter(name, class string, level, hpMax, hp int, inv []string) Character {
	return Character{
		Name: name, Class: class, Level: level,
		HPMax: hpMax, HP: hp, Inventory: inv,
		Money: 100, // Tâche 13. :contentReference[oaicite:10]{index=10}
		Skills: []string{"coup de poing"}, // Tâche 10 de base. :contentReference[oaicite:11]{index=11}
		Equipment: Equipment{},
		InventoryCap: 10, // Tâche 12: limite (de base 10). :contentReference[oaicite:12]{index=12}
		ExpMax: 100, ManaMax: 50,
	}
}

func CharacterCreation() Character {
	// Tâche 11 : nom lettres + Majuscule/dreste en minuscules + choix de classe Humain/Elfe/Nain. :contentReference[oaicite:13]{index=13}
	name := sanitizeName(promptLetters("Ton nom (lettres seulement) : "))
	class := chooseClass()
	hpMax := 100
	switch class {
	case "elfe":
		hpMax = 80
	case "nain":
		hpMax = 120
	}
	hp := hpMax / 2
	return initCharacter(name, class, 1, hpMax, hp, []string{"potion"})
}

func DisplayInfo(c *Character) {
	fmt.Printf("Nom: %s | Classe: %s | Nv: %d | PV: %d/%d | Or: %d\n",
		c.Name, c.Class, c.Level, c.HP, c.HPMax, c.Money)
	fmt.Printf("Équipement: tête=%s, torse=%s, pieds=%s\n", c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet)
	fmt.Printf("Skills: %v | Mana: %d/%d | Exp: %d/%d | Initiative: %d\n",
		c.Skills, c.Mana, c.ManaMax, c.Exp, c.ExpMax, c.Initiative)
}



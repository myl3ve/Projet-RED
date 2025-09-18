
package main

// =======================
// types.go : Déclarations
// =======================

// --- Items ---
type ItemType int

const (
	Consumable ItemType = iota // Potion, etc.
	Material                    // Ressource de craft
	EquipmentItem              // Équipement (casque, tunique, bottes)
	Upgrade                    // Augmentation d'inventaire
	Misc                       // Livre de sort, etc.
)

type Item struct {
	Name string
	Type ItemType
}

// --- Équipement porté ---
type Equipment struct {
	Head string
	Body string
	Feet string
}

// --- Personnage (Tâches 1–2) ---
type Character struct {
	Name          string
	Class         string
	Level         int
	MaxHP         int
	HP            int
	MaxMana       int
	Mana          int
	Gold          int
	Experience    int
	ExperienceMax int

	Inventory    []Item
	InventoryCap int

	Equip       Equipment
	Skills      []string
	Initiative  int
}

// --- Monstre (Tâche 19) ---
type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	Attack     int
	Initiative int
	ExpReward  int
}

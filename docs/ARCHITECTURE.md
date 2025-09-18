
# ARCHITECTURE — Organisation du code

- `src/main.go` → entrée, création perso, menu principal.
- `src/visual.go` → couleurs ANSI, ASCII, typewriter, centrage.
- `src/types.go` → `Character`, `Item`, `Equipment`, `Monster`.
- `src/utils.go` → saisie (`AskLine`, `AskInt`), validation nom/skills.
- `src/inventory.go` → inventaire, potions, upgrade, `useItem`.
- `src/merchant.go` → marchand + prix.
- `src/forge.go` → recettes, fabrication, `equip` (bonus MaxHP).
- `src/combat.go` → **nouveau combat** (HUD, pattern x2, tour joueur, entraînement).
- `src/story.go` → campagne Niveaux 1→7 (débloque armes, remède, boss).

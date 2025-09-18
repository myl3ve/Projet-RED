
# RENDU — Couverture des Tâches (1 → 22.2)

Ce document relie **chaque Tâche du PDF** au **code** correspondant.

## 1–2. Personnage (struct + init)
- `src/types.go` → `Character`
- `src/main.go` → `characterCreation()`

## 3. Affichage infos
- `src/main.go` → `showInfos(c *Character)`

## 4. Inventaire / utilisation
- `src/inventory.go` → `accessInventory(c)` + `useItem(...)`

## 5. Potion de vie (cap PV)
- `src/inventory.go` → `usePotionVie(c)` (+50 PV, cap MaxHP)

## 6. Menu principal + “Retour”
- `src/main.go` → `runMenus(c)`

## 7. Ajout/Suppression inventaire (Marchand)
- `src/inventory.go` → `addInventory`, `removeInventory`
- `src/merchant.go` → `showMerchant`

## 8. Mort + Résurrection 50%
- `src/combat.go` → `isDead(c)`

## 9. Poison 10 PV/s ×3
- `src/inventory.go` → `applyPoison(c)` (avec `time.Sleep(1s)`)

## 10. Livre de sort (Boule de feu)
- `src/inventory.go` → `useSpellBook(c)`

## 11. Création (nom lettres, formatage)
- `src/utils.go` → `onlyLetters`, `formatName`
- `src/main.go` → `characterCreation()`

## 12. Limite inventaire
- `src/inventory.go` → `checkInventoryCap`

## 13–14. Or + Prix Marchand
- `src/merchant.go` → `prices` (3/6/25/4/7/3/1/30)

## 15. Forgeron (recettes + coût)
- `src/forge.go` → `showForge`, `hasResourcesFor`, `consumeResourcesFor`

## 16–17. Équipement + bonus PV
- `src/types.go` → `Equipment`
- `src/forge.go` → `equip(...)` (bonus +10/+25/+15)

## 18. Upgrade inventaire (+10, max 40)
- `src/inventory.go` → `useItem` pour `"Augment. Inventaire"`

## 19. Monstre d’entrainement
- `src/combat.go` → `initGoblin()`

## 20. Pattern (×2 tous les 3 tours)
- `src/combat.go` → `goblinPattern(...)`

## 21. Tour du joueur
- `src/combat.go` → `charTurn(...)` (Attaque/Inventaire/Armes)

## 22.1–22.2. Entrainement (boucle + fin)
- `src/combat.go` → `trainingFight(c)`

### BONUS — Histoire Niveaux 1→7
- `src/story.go` → `runStory(c)`
- `src/visual.go` → couleurs, ASCII, typewriter.


# ORAL 

## 1) Pourquoi ce design ?
- **Package unique** `main` => zéro erreur d'import, plus simple à expliquer.
- **Types clairs** (`Character`, `Item`, `Monster`, `Equipment`) => je montre que je sais modéliser.
- **Fonctions courtes** => chaque fonctionnalité est isolée et testable.

## 2) Références au PDF (pages & tâches)
- **Tâche 1–2** (p.10–11): `types.go` (Character) + `initCharacter()` dans `main.go`.
- **Tâche 3** (p.12): `showInfos()` affiche tous les attributs.
- **Tâche 4** (p.13): `accessInventory()` liste & utilisations.
- **Tâche 5** (p.14): `usePotionVie()` soigne en capant à MaxHP.
- **Tâche 6** (p.15): menu principal avec `switch` + “Retour”.
- **Tâche 7** (p.16): `showMerchant()` + `addInventory` / `removeInventory`.
- **Tâche 8** (p.17): `isDead()` => rez à 50% PV max.
- **Tâche 9** (p.18): `applyPoison()` **avec `time.Sleep(1s)`** (exigé).
- **Tâche 10** (p.19): `useSpellBook()` ajoute **Boule de feu** si non appris.
- **Tâche 11** (p.20): `characterCreation()` (nom lettres + Majuscule/minuscule).
- **Tâche 12** (p.21): `checkInventoryCap()` limite 10 items (puis upgrade).
- **Tâches 13–14** (p.24–25): **or** + **prix exacts PDF** dans `merchant.go`.
- **Tâche 15** (p.25–27): `showForge()` + recettes (Fourrure, Peau, Cuir, Plume).
- **Tâche 16** (p.27): `Equipment` struct + champ dans `Character`.
- **Tâche 17** (p.28): bonus PV (+10 tête, +25 torse, +15 pieds) + remplacement.
- **Tâche 18** (p.29): `Augment. Inventaire` (+10, max 3 fois, 30 or).
- **Tâche 19** (p.32–33): `Monster` + `initGoblin()` (40 PV, atk=5).
- **Tâche 20** (p.34): `goblinPattern()` (100% dmg / ×2 tous les 3 tours, logs).
- **Tâche 21** (p.35–36): `charTurn()` (Attaquer/Inventaire + logs).
- **Tâche 22.1** (p.37): `trainingFight()` (tour affiché, alternance).
- **Tâche 22.2** (p.38): fin si PV ≤ 0, retour menu.

## 3) “Beau terminal” —
- “J’utilise des **couleurs ANSI**, un **ASCII Art** pour le titre et le panneau *COMBAT ENGAGÉ*, et un effet **machine à écrire** pour la narration.”
- “Les **commentaires** expliquent chaque choix pour qu’un **débutant** comprenne.”

## 4) Démonstration en 60 secondes
1. Lancer → **Histoire**.
2. Niveau 2 : Marchand → acheter **Potion de vie**.
3. Niveau 3 : Combat zombie → victoire ⇒ **Fusil** débloqué + contamination.
4. Niveau 6 : Remède **Violet = Rouge + Bleu**.
5. Boss final → sortie.

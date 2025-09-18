
# COMBAT — Guide complet (version guidée)

## 1) Ordre des tours
- On compare l’**Initiative** du joueur et du monstre.
- Le plus haut **commence**.
- Ensuite, on alterne : *Joueur → Monstre* ou *Monstre → Joueur*.

## 2) Tour du joueur (`charTurn`)
- **Attaque basique** : 5 dégâts (conforme PDF).
- **Inventaire** : utiliser une potion/livre consomme le tour.
- **Armes** (si débloquées via l’histoire) :
  - Fusil (12), AK47 (16), MP5 (14).

## 3) Pattern ennemi (`goblinPattern`)
- Dégâts = `Attack` (**normaux**).
- Tous les **3 tours** (3, 6, 9, …) → **x2 dégâts** (coup fort).
- Affiche un warning `⚠️` quand c’est un **coup fort**.

## 4) Mort / Résurrection (`isDead`)
- Si PV ≤ 0 : message “💀 Vous êtes mort.” puis **réanimation à 50%**.
- Évite de bloquer la démo.

## 5) HUD
- À chaque tour : barre d’état **centrée** avec PV joueur / PV monstre.

## 6) Entraînement (`trainingFight`)
- Affiche les **règles** au début (aide au joueur).
- Incrémente le **tour**, affiche le **HUD**, applique **pattern** et **tour joueur**.
- Fin : **Victoire** (+XP) ou “Fin de l’entraînement”.


# ORAL — Antisèche

## Pitch (20s)
“*Locker 666* est un jeu console en **Go**. Il couvre toutes les **Tâches 1→22.2** (inventaire, marchand, forge, combat). J’ai intégré **mon histoire** avec des **couleurs**, **ASCII**, et un effet **machine à écrire**. Le code est **débutant-friendly** : une `package main`, fonctions courtes, commentaires partout.”

## Démo (1 min)
1. Lancer → **Histoire**.
2. **Marchand** (Niveau 2) : acheter **Potion de vie**.
3. **Combat** (Niveau 3) : victoire → **Fusil** débloqué + **contamination**.
4. **Remède** (Niveau 6) : Violet = **Rouge + Bleu** (PV full).
5. **Boss** : victoire → sortie déverrouillée.

## Points techniques
- `Character/Item/Monster` → modélisation claire.
- `applyPoison` avec `time.Sleep(1s)` (conformité PDF).
- `equip` → recalcul **MaxHP** avec bonus.
- **HUD** de combat centré + rappel **x2 au tour 3/6/9**.

## Q/R
- **Pourquoi `main` unique ?** Pour éviter les erreurs d’import et me concentrer sur la logique.
- **Limite d’inventaire ?** `checkInventoryCap` + upgrade +10 (max 40). 
- **Prix Marchand ?** Dans la map `prices`, selon le PDF.

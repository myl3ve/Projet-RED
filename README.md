# Projet RED – Jeu texte en Go

Projet développé dans le cadre de l’Ymmersion Ynov Campus.  
Jeu d’aventure textuel intégrant création de personnage, inventaire, marchand, craft et combats.

## 📦 Prérequis

- Go 1.22 ou version ultérieure

## 🚀 Installation

Clonez le projet et placez-vous à la racine (là où se trouve `go.mod`) :

```bash
git clone https://github.com/toncompte/projet-red-monjeu.git
cd projet-red-monjeu
go mod tidy
▶️ Lancement
Exécutez le jeu depuis la racine du projet :

bash
Copier le code
go run ./src
Le menu principal apparaîtra dans le terminal.

🗂 Structure du projet
bash
Copier le code
projet-red-monjeu/
├─ go.mod
├─ README.md
├─ docs/
│  └─ gestion_projet.md
└─ src/
   ├─ main.go           # point d’entrée du jeu
   ├─ ui/               # interface utilisateur (menu, entrées)
   ├─ character/        # gestion du personnage, inventaire, marchand, forge
   └─ combat/           # gestion des monstres et combats
✨ Fonctionnalités
Création et affichage du personnage

Inventaire et objets utilisables

Marchand et forge (achat / craft d’équipements)

Combat d’entraînement contre gobelin et système de tours

Gestion du poison et des PV
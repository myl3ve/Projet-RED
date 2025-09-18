
# VISUELS — Couleurs, ASCII, centrage

## Couleurs ANSI
- `color(colRed, "texte")`, `bold("texte")` — définis dans `src/visual.go`.
- Codes ANSI de base : Red/Green/Yellow/Blue/Magenta/Cyan.

## ASCII art
- Pour **gros** ASCII, utilise des **raw strings** `` `...` ``.
- Si l’ASCII contient des **backticks** (`), remplace-les par `[BT]` puis :
  ```go
  art = strings.ReplaceAll(art, "[BT]", "`")
  ```
- Si tu utilises des chaînes `"..."`, **double** les antislashs `\`.

## Centrer du texte
- Version simple (largeur fixe) : `centerLine(s)`, `centerBlock(s)` (voir `visual.go`).
- Version auto (taille console) : utilise `golang.org/x/term` (optionnel).

## Petites animations
- **Typewriter** : `slow("Texte...", 18)` (ms/caractère).
- **Pause** : `pause(300)` (ms).
- **Clear screen** : `clear()`.

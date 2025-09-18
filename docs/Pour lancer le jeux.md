
# INSTALL — Installation & Lancement

## Prérequis
- **Go 1.21+** : https://go.dev/dl/
- Terminal monospace (Windows Terminal / PowerShell / iTerm2).

## Lancer le jeu
```bash
cd projet-red_lycee-deluxe
go mod tidy
go run ./src
```
> Menu : **Histoire** (Niveaux 1→7) ou **Libre** (menus).

## Couleurs & ASCII (si affichage bizarre)
- Ouvre `src/visual.go` et règle `useColor = true/false`.
- **Windows** (élargir la fenêtre pour grands ASCII) :
  - PowerShell :
    ```powershell
    $Host.UI.RawUI.WindowSize = New-Object Management.Automation.Host.Size(150, 40)
    $Host.UI.RawUI.BufferSize = New-Object Management.Automation.Host.Size(150, 3000)
    ```
  - CMD :
    ```bat
    mode con: cols=150 lines=40
    ```

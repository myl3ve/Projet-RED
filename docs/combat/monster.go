package combat

type Monster struct {
	Name string
	HPMax int
	HP int
	ATK int
	Initiative int
	ExpReward int
}

func initGoblin() Monster { // Tâche 19. :contentReference[oaicite:26]{index=26}
	return Monster{
		Name: "Gobelin d’entrainement", HPMax: 40, HP: 40, ATK: 5,
		Initiative: 5, ExpReward: 30,
	}
}

package main

func InsertionSortScores(players []Score) {
	for i := 1; i < len(players); i++ {
		valeur := players[i]
		j := i - 1
		for j >= 0 && players[j].Score < valeur.Score {
			players[j+1] = players[j]
			j--
		}
		players[j+1] = valeur
	}
}

func SelectionSortScores(players []Score) {
	n := len(players)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if players[j].Score > players[max].Score {
				max = j
			}
		}
		players[i], players[max] = players[max], players[i]
	}
}

func IsStable(players []Score) bool {
	for i := 0; i < len(players)-1; i++ {
		if players[i].Score == players[i+1].Score && players[i].Player > players[i+1].Player {
			return false
		}
	}
	return true
}

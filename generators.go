package main

import (
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(69))

func RandomScores(n int) []int {
	scores := make([]int, n)
	for i := range scores {
		scores[i] = r.Intn(1001)
	}
	return scores
}

func SortedScores(n int) []int {
	scores := make([]int, n)
	for i := range scores {
		scores[i] = i
	}
	return scores
}

func ReversedScores(n int) []int {
	scores := SortedScores(n)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		scores[i], scores[j] = scores[j], scores[i]
	}
	return scores
}

func NearlySortedScores(n int) []int {
	scores := SortedScores(n)
	for k := 0; k < n/100; k++ {
		i, j := r.Intn(n), r.Intn(n)
		scores[i], scores[j] = scores[j], scores[i]
	}
	return scores
}

func RandomPlayers(n int) []Score {
	players := make([]Score, n)
	for i := range players {
		players[i] = Score{Player: fmt.Sprintf("Joueur%05d", i+1), Score: r.Intn(101)}
	}
	return players
}

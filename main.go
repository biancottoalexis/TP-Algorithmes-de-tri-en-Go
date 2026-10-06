package main

import (
	"errors"
	"fmt"
)

func main() {
	demoInt := []int{5, 2, 5, 1, 2}
	fmt.Println("Avant :", demoInt)
	if err := BubbleSort(demoInt); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Après :", demoInt)
	}

	var nilSlice []int
	err := BubbleSort(nilSlice)
	fmt.Println("BubbleSort(nil) :", err, "| ErrNilSlice ?", errors.Is(err, ErrNilSlice))

	err = BubbleSort([]int{})
	fmt.Println("BubbleSort([]) :", err, "| ErrEmptySlice ?", errors.Is(err, ErrEmptySlice))

	players := []Score{{"Léa", 12}, {"Tom", 15}, {"Ana", 12}, {"Max", 9}}
	if err := InsertionSortScores(players); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Joueurs triés :", players, "| stable ?", IsStable(players))
	}
}

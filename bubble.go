package main

import "errors"

var (
	ErrNilSlice   = errors.New("slice nil")
	ErrEmptySlice = errors.New("slice vide")
)

func BubbleSort(scores []int) error {
	if scores == nil {
		return ErrNilSlice
	}
	if len(scores) == 0 {
		return ErrEmptySlice
	}

	for lastIndex := len(scores) - 1; lastIndex > 0; lastIndex-- {
		swapped := false
		for i := 0; i < lastIndex; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
				swapped = true
			}
		}
		if !swapped {
			return nil
		}
	}
	return nil
}

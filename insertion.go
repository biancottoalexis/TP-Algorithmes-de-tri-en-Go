package main

func InsertionSort(scores []int) error {
	if scores == nil {
		return ErrNilSlice
	}
	if len(scores) == 0 {
		return ErrEmptySlice
	}

	for i := 1; i < len(scores); i++ {
		value := scores[i]
		j := i - 1
		for j >= 0 && scores[j] > value {
			scores[j+1] = scores[j]
			j--
		}
		scores[j+1] = value
	}
	return nil
}

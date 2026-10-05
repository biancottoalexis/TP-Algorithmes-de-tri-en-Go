package main

func SelectionSort(scores []int) {
	for i := 0; i < len(scores)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(scores); j++ {
			if scores[j] < scores[minIndex] {
				minIndex = j
			}
		}
		scores[i], scores[minIndex] = scores[minIndex], scores[i]
	}
}

package main

func QuickSort(scores []int) {
	if len(scores) <= 1 {
		return
	}

	lastIndex := len(scores) - 1
	pivot := scores[lastIndex]
	partitionIndex := 0

	for i := 0; i < lastIndex; i++ {
		if scores[i] < pivot {
			scores[i], scores[partitionIndex] = scores[partitionIndex], scores[i]
			partitionIndex++
		}
	}

	scores[partitionIndex], scores[lastIndex] = scores[lastIndex], scores[partitionIndex]

	QuickSort(scores[:partitionIndex])
	QuickSort(scores[partitionIndex+1:])
}

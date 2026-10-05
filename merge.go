package main

func MergeSort(scores []int) []int {
	if len(scores) <= 1 {
		return scores
	}
	middle := len(scores) / 2
	leftPacket := MergeSort(scores[:middle])
	rightPacket := MergeSort(scores[middle:])
	return merge(leftPacket, rightPacket)
}

func merge(leftPacket, rightPacket []int) []int {
	result := make([]int, 0, len(leftPacket)+len(rightPacket))
	i, j := 0, 0
	for i < len(leftPacket) && j < len(rightPacket) {
		if rightPacket[j] < leftPacket[i] {
			result = append(result, rightPacket[j])
			j++
		} else {
			result = append(result, leftPacket[i])
			i++
		}
	}
	result = append(result, leftPacket[i:]...)
	result = append(result, rightPacket[j:]...)
	return result
}

package main

import (
	"fmt"
	"slices"
	"testing"
)

// ---------- EX06 : grand comparatif ----------

func BenchmarkBubbleSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				BubbleSort(scores)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				SelectionSort(scores)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				InsertionSort(scores)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		var resultat []int
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultat, _ = MergeSort(base)
			}
		})
		_ = resultat
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				QuickSort(scores)
			}
		})
	}
}

func BenchmarkStdSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				slices.Sort(scores)
			}
		})
	}
}

// ---------- EX01 : bulles sur donnée triée vs inversée ----------

func BenchmarkBubbleSortOrdre(b *testing.B) {
	n := 10_000
	trie := SortedScores(n)
	inverse := ReversedScores(n)
	scores := make([]int, n)

	b.Run("Sorted", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, trie)
			b.StartTimer()
			BubbleSort(scores)
		}
	})
	b.Run("Reversed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, inverse)
			b.StartTimer()
			BubbleSort(scores)
		}
	})
}

// ---------- EX02 : sélection sur les 3 ordres ----------

func BenchmarkSelectionSortOrdre(b *testing.B) {
	n := 10_000
	aleatoire := RandomScores(n)
	trie := SortedScores(n)
	inverse := ReversedScores(n)
	scores := make([]int, n)

	b.Run("Random", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, aleatoire)
			b.StartTimer()
			SelectionSort(scores)
		}
	})
	b.Run("Sorted", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, trie)
			b.StartTimer()
			SelectionSort(scores)
		}
	})
	b.Run("Reversed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, inverse)
			b.StartTimer()
			SelectionSort(scores)
		}
	})
}

// ---------- EX03 : bulles vs insertion sur presque trié ----------

func BenchmarkPresqueTrie(b *testing.B) {
	n := 100_000
	base := NearlySortedScores(n)
	scores := make([]int, n)

	b.Run("Bubble", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, base)
			b.StartTimer()
			BubbleSort(scores)
		}
	})
	b.Run("Insertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, base)
			b.StartTimer()
			InsertionSort(scores)
		}
	})
}

// ---------- EX07 : insertion sur joueurs vs sur scores ----------

func BenchmarkInsertionSortScores(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomPlayers(n)
		players := make([]Score, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(players, base)
				b.StartTimer()
				InsertionSortScores(players)
			}
		})
	}
}

// ---------- EX08 : sélection sur joueurs ----------

func BenchmarkSelectionSortScores(b *testing.B) {
	n := 10_000
	base := RandomPlayers(n)
	players := make([]Score, n)
	b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(players, base)
			b.StartTimer()
			SelectionSortScores(players)
		}
	})
}

// ---------- EX08 : vérification de stabilité ----------

func TestStability(t *testing.T) {
	n := 10_000
	base := RandomPlayers(n)

	insertion := make([]Score, n)
	copy(insertion, base)
	InsertionSortScores(insertion)
	t.Logf("InsertionSortScores stable : %v", IsStable(insertion))

	selection := make([]Score, n)
	copy(selection, base)
	SelectionSortScores(selection)
	t.Logf("SelectionSortScores stable : %v", IsStable(selection))
}

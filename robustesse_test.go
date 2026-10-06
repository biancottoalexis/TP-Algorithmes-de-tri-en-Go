package main

import (
	"errors"
	"slices"
	"testing"
)

func cloneInts(s []int) []int {
	if s == nil {
		return nil
	}
	out := make([]int, len(s))
	copy(out, s)
	return out
}

func cloneScores(s []Score) []Score {
	if s == nil {
		return nil
	}
	out := make([]Score, len(s))
	copy(out, s)
	return out
}

var casLimitesInt = []struct {
	name    string
	input   []int
	want    []int
	wantErr error
}{
	{name: "SliceNil", input: nil, wantErr: ErrNilSlice},
	{name: "SliceVide", input: []int{}, wantErr: ErrEmptySlice},
	{name: "UnSeulElement", input: []int{7}, want: []int{7}},
	{name: "Doublons", input: []int{5, 2, 5, 1, 2}, want: []int{1, 2, 2, 5, 5}},
	{name: "DejaTrie", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
	{name: "TrieALEnvers", input: []int{4, 3, 2, 1}, want: []int{1, 2, 3, 4}},
	{name: "ValeursNegatives", input: []int{-3, 5, -1, 0, -8}, want: []int{-8, -3, -1, 0, 5}},
}

func TestBubbleSortCasLimites(t *testing.T) {
	for _, test := range casLimitesInt {
		t.Run(test.name, func(t *testing.T) {
			got := cloneInts(test.input)
			err := BubbleSort(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: BubbleSort(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: BubbleSort(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestSelectionSortCasLimites(t *testing.T) {
	for _, test := range casLimitesInt {
		t.Run(test.name, func(t *testing.T) {
			got := cloneInts(test.input)
			err := SelectionSort(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: SelectionSort(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: SelectionSort(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestInsertionSortCasLimites(t *testing.T) {
	for _, test := range casLimitesInt {
		t.Run(test.name, func(t *testing.T) {
			got := cloneInts(test.input)
			err := InsertionSort(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: InsertionSort(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: InsertionSort(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestQuickSortCasLimites(t *testing.T) {
	for _, test := range casLimitesInt {
		t.Run(test.name, func(t *testing.T) {
			got := cloneInts(test.input)
			err := QuickSort(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: QuickSort(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: QuickSort(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestMergeSortCasLimites(t *testing.T) {
	for _, test := range casLimitesInt {
		t.Run(test.name, func(t *testing.T) {
			got, err := MergeSort(test.input)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: MergeSort(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: MergeSort(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestMergeSortNeModifiePasLOriginal(t *testing.T) {
	original := []int{5, 2, 5, 1, 2}
	avant := cloneInts(original)
	MergeSort(original)
	if !slices.Equal(original, avant) {
		t.Errorf("MergeSort a modifié son paramètre : %v, attendu %v", original, avant)
	}
}

var casLimitesScores = []struct {
	name    string
	input   []Score
	want    []Score
	wantErr error
}{
	{name: "SliceNil", input: nil, wantErr: ErrNilSlice},
	{name: "SliceVide", input: []Score{}, wantErr: ErrEmptySlice},
	{name: "UnSeulJoueur", input: []Score{{"A", 10}}, want: []Score{{"A", 10}}},
	{
		name:  "DoublonsDeScore",
		input: []Score{{"Léa", 12}, {"Tom", 15}, {"Ana", 12}, {"Max", 9}},
		want:  []Score{{"Tom", 15}, {"Léa", 12}, {"Ana", 12}, {"Max", 9}},
	},
}

func TestInsertionSortScoresCasLimites(t *testing.T) {
	for _, test := range casLimitesScores {
		t.Run(test.name, func(t *testing.T) {
			got := cloneScores(test.input)
			err := InsertionSortScores(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: InsertionSortScores(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr == nil && !slices.Equal(got, test.want) {
				t.Errorf("%s: InsertionSortScores(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestSelectionSortScoresCasLimites(t *testing.T) {
	for _, test := range casLimitesScores {
		t.Run(test.name, func(t *testing.T) {
			got := cloneScores(test.input)
			err := SelectionSortScores(got)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: SelectionSortScores(%v) erreur = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}
			if test.wantErr != nil {
				return
			}
			if len(got) != len(test.want) {
				t.Fatalf("%s: SelectionSortScores(%v) longueur = %d; want %d", t.Name(), test.input, len(got), len(test.want))
			}
			for i := 0; i < len(got)-1; i++ {
				if got[i].Score < got[i+1].Score {
					t.Errorf("%s: SelectionSortScores(%v) = %v, pas trié par score décroissant", t.Name(), test.input, got)
				}
			}
		})
	}
}

func TestIsStableCasLimites(t *testing.T) {
	testCases := []struct {
		name  string
		input []Score
		want  bool
	}{
		{name: "SliceNil", input: nil, want: true},
		{name: "SliceVide", input: []Score{}, want: true},
		{name: "UnSeulJoueur", input: []Score{{"A", 5}}, want: true},
		{name: "CasStable", input: []Score{{"C", 9}, {"A", 5}, {"B", 5}}, want: true},
		{name: "CasInstable", input: []Score{{"C", 9}, {"B", 5}, {"A", 5}}, want: false},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := IsStable(test.input)
			if got != test.want {
				t.Errorf("%s: IsStable(%v) = %v; want %v", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestGenerateursNNegatif(t *testing.T) {
	if got := RandomScores(-5); len(got) != 0 {
		t.Errorf("RandomScores(-5) = %v; want un slice vide", got)
	}
	if got := SortedScores(-5); len(got) != 0 {
		t.Errorf("SortedScores(-5) = %v; want un slice vide", got)
	}
	if got := ReversedScores(-5); len(got) != 0 {
		t.Errorf("ReversedScores(-5) = %v; want un slice vide", got)
	}
	if got := NearlySortedScores(-5); len(got) != 0 {
		t.Errorf("NearlySortedScores(-5) = %v; want un slice vide", got)
	}
	if got := RandomPlayers(-5); len(got) != 0 {
		t.Errorf("RandomPlayers(-5) = %v; want un slice vide", got)
	}
}

func TestGenerateursValeursAttendues(t *testing.T) {
	if got := SortedScores(5); !slices.Equal(got, []int{0, 1, 2, 3, 4}) {
		t.Errorf("SortedScores(5) = %v; want [0 1 2 3 4]", got)
	}
	if got := ReversedScores(5); !slices.Equal(got, []int{4, 3, 2, 1, 0}) {
		t.Errorf("ReversedScores(5) = %v; want [4 3 2 1 0]", got)
	}
	if got := RandomScores(100); len(got) != 100 {
		t.Errorf("RandomScores(100) a une longueur de %d; want 100", len(got))
	}
	if got := RandomPlayers(10); len(got) != 10 {
		t.Errorf("RandomPlayers(10) a une longueur de %d; want 10", len(got))
	}
	if got := NearlySortedScores(200); len(got) != 200 {
		t.Errorf("NearlySortedScores(200) a une longueur de %d; want 200", len(got))
	}
}

func TestMainSmoke(t *testing.T) {
	main()
}

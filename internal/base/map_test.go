package base

import (
	"reflect"
	"testing"
)

func TestFrequencies(t *testing.T) {
    t.Parallel()
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{
			name:  "nil input",
			words: nil,
			want:  map[string]int{},
		},
		{
			name:  "empty input",
			words: []string{},
			want:  map[string]int{},
		},
		{
			name:  "single word",
			words: []string{"go"},
			want:  map[string]int{"go": 1},
		},
		{
			name:  "no duplicates",
			words: []string{"a", "b", "c"},
			want:  map[string]int{"a": 1, "b": 1, "c": 1},
		},
		{
			name:  "with duplicates",
			words: []string{"go", "go", "rust", "go", "rust", "c"},
			want:  map[string]int{"go": 3, "rust": 2, "c": 1},
		},
		{
			name:  "all same",
			words: []string{"x", "x", "x", "x"},
			want:  map[string]int{"x": 4},
		},
		{
			name:  "empty strings",
			words: []string{"", "", "a"},
			want:  map[string]int{"": 2, "a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		    t.Parallel()
			got := Frequencies(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Frequencies(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestMostFrequent(t *testing.T) {
    t.Parallel()
	tests := []struct {
		name      string
		freq      map[string]int
		wantWord  string
		wantCount int
		wantOK    bool
	}{
		{
			name:      "nil map",
			freq:      nil,
			wantWord:  "",
			wantCount: 0,
			wantOK:    false,
		},
		{
			name:      "empty map",
			freq:      map[string]int{},
			wantWord:  "",
			wantCount: 0,
			wantOK:    false,
		},
		{
			name:      "single entry",
			freq:      map[string]int{"go": 5},
			wantWord:  "go",
			wantCount: 5,
			wantOK:    true,
		},
		{
			name:      "clear winner",
			freq:      map[string]int{"go": 3, "rust": 1, "c": 2},
			wantWord:  "go",
			wantCount: 3,
			wantOK:    true,
		},
		{
			name: "equal max frequency — tie between two",
			freq: map[string]int{"go": 3, "rust": 3, "c": 1},
			wantWord:  "go",
			wantCount: 3,
			wantOK:    true,
		},
		{
			name: "equal max frequency — tie, shorter wins regardless of order",
			freq: map[string]int{"rust": 4, "c": 4, "go": 4},
			wantWord:  "c",
			wantCount: 4,
			wantOK:    true,
		},
		{
			name: "equal max frequency — same length tie",
			freq: map[string]int{"aa": 2, "bb": 2},
			wantWord:  "",
			wantCount: 2,
			wantOK:    true,
		},
		{
			name:      "all equal counts",
			freq:      map[string]int{"a": 1, "bb": 1, "ccc": 1},
			wantWord:  "a",
			wantCount: 1,
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		    t.Parallel()
			word, count, ok := MostFrequent(tt.freq)

			if count != tt.wantCount {
				t.Errorf("count = %d, want %d", count, tt.wantCount)
			}
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !tt.wantOK {
				if word != "" {
					t.Errorf("word = %q, want empty", word)
				}
				return
			}

			if tt.wantWord != "" {
				if word != tt.wantWord {
					t.Errorf("word = %q, want %q", word, tt.wantWord)
				}
				return
			}

			if tt.freq[word] != tt.wantCount {
				t.Errorf("word = %q (freq %d), want any key with freq %d",
					word, tt.freq[word], tt.wantCount)
			}
		})
	}
}
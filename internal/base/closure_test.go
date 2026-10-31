package base

import "testing"


func TestAccumulator_BasicCases(t *testing.T) {
    t.Parallel()

	acc := Accumulator()

	tests := []struct {
		input int
		want  int
	}{
		{2, 2},
		{-1, 1},
		{5, 6},
		{10, 16},
	}

	for _, tt := range tests {
		got := acc(tt.input)
		if got != tt.want {
			t.Errorf("acc(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestReaders_Basic(t *testing.T) {
    t.Parallel()

	rs := Readers([]int{10, 20})
	if len(rs) != 2 {
		t.Fatalf("len = %d, want 2", len(rs))
	}
	if got := rs[0](); got != 10 {
		t.Errorf("rs[0]() = %d, want 10", got)
	}
	if got := rs[1](); got != 20 {
		t.Errorf("rs[1]() = %d, want 20", got)
	}
}

func TestReaders_DoesNotChangeAfterInputMutation(t *testing.T) {
    t.Parallel()

	input := []int{10, 20}
	rs := Readers(input)

	input[0] = 999
	input[1] = 888

	if got := rs[0](); got != 10 {
		t.Errorf("rs[0]() after mutation = %d, want 10", got)
	}
	if got := rs[1](); got != 20 {
		t.Errorf("rs[1]() after mutation = %d, want 20", got)
	}
}
package dp_practice

import (
	"reflect"
	"testing"
)

// TODO - write test cases.

func TestFindAllUniqueCombinationsForTarget_TargetZero(t *testing.T) {
	// reset global state
	result = nil

	got := FindAllUniqueCombinationsForTarget([]int{2, 3, 6}, 0)
	want := [][]int{{}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %#v, got %#v", want, got)
	}
}

func TestFindAllUniqueCombinationsForTarget_PositiveTarget_NoResults(t *testing.T) {
	// reset global state
	result = nil

	got := FindAllUniqueCombinationsForTarget2([]int{2, 3, 6, 7}, 7)
	t.Fatalf("obtained value of got : %v", got)
	if len(got) != 0 {
		t.Fatalf("expected no combinations (len 0), got %#v", got)
	}
}

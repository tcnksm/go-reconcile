package reconcile

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stringCases = map[string]struct {
	desireds, currents []string
	want               *Action[string]
}{
	"Add": {
		[]string{"a", "b", "c"},
		[]string{"a", "b"},
		&Action[string]{
			Adds: []string{"c"},
		},
	},
	"Delete": {
		[]string{"a", "b"},
		[]string{"a", "b", "c"},
		&Action[string]{
			Deletes: []string{"c"},
		},
	},
	"AddAndDelete": {
		[]string{"a", "b", "d"},
		[]string{"a", "b", "c"},
		&Action[string]{
			Adds:    []string{"d"},
			Deletes: []string{"c"},
		},
	},
	"NoChange1": {
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		&Action[string]{},
	},
	"NoChange2": {
		[]string{},
		[]string{},
		&Action[string]{},
	},
}

var intCases = map[string]struct {
	desireds, currents []int
	want               *Action[int]
}{
	"Add": {
		[]int{1, 2, 3},
		[]int{1, 2},
		&Action[int]{
			Adds: []int{3},
		},
	},
	"Delete": {
		[]int{1, 2},
		[]int{1, 2, 3},
		&Action[int]{
			Deletes: []int{3},
		},
	},
	"AddAndDelete": {
		[]int{1, 2, 3},
		[]int{1, 2, 4},
		&Action[int]{
			Adds:    []int{3},
			Deletes: []int{4},
		},
	},
	"NoChange1": {
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		&Action[int]{},
	},
	"NoChange2": {
		[]int{},
		[]int{},
		&Action[int]{},
	},
}

func TestPlan(t *testing.T) {
	for name, tc := range stringCases {
		t.Run(fmt.Sprintf("String/%s", name), func(t *testing.T) {
			got := Plan(tc.desireds, tc.currents)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("Plan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
	for name, tc := range intCases {
		t.Run(fmt.Sprintf("Int/%s", name), func(t *testing.T) {
			got := Plan(tc.desireds, tc.currents)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("Plan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

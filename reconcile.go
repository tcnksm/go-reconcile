package reconcile

import "golang.org/x/exp/slices"

// Action is results returned by Plan.
type Action[T comparable] struct {
	Adds    []T
	Deletes []T
}

// Plan takes desired lists and current lists and check what to delete
// or to add. It returns results as Action.
func Plan[V comparable](desireds []V, currents []V) *Action[V] {
	var (
		adds    []V
		deletes []V
	)

	exists := make(map[V]struct{})
	for _, current := range currents {
		if slices.Contains(desireds, current) {
			exists[current] = struct{}{}
		} else {
			deletes = append(deletes, current)
		}
	}

	for _, desired := range desireds {
		if _, ok := exists[desired]; !ok {
			adds = append(adds, desired)
		}
	}

	return &Action[V]{
		Adds:    adds,
		Deletes: deletes,
	}
}

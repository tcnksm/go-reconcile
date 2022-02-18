package reconcile

import "golang.org/x/exp/slices"

// Action is results returned by Plan function. It contains 
// the items to add to or delete from the current observed list
// to reach to the desired state.
type Action[T comparable] struct {
	// The list of items to add to the current list.
	Adds    []T
	// The list of items to delete from the current list.
	Deletes []T
}

// Plan takes desired lists and current observed lists and decideds 
// what action to take (which item to add to or delete from the current 
// list) to reach to the given desired list. The results are returns as Action.
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

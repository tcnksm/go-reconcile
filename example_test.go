package reconcile

import "fmt"

func ExamplePlan_string() {
	// The list of the item you desire to have.
	desireds := []string{"alice", "bob", "charlie"}

	// The list of the item currently you observes.
	currents := []string{"alice", "bob", "dan", "frank"}

	actions := Plan(desireds, currents)
	fmt.Printf("Add: %v, Delete: %v", actions.Adds, actions.Deletes)
	// Output:
	// Add: [charlie], Delete: [dan frank]
}

func ExamplePlan_int() {
	// The list of the item you desire to have.
	desireds := []int{1, 2, 3}

	// The list of the item currently you observes.
	currents := []int{1, 2, 4, 5}

	actions := Plan(desireds, currents)
	fmt.Printf("Add: %v, Delete: %v", actions.Adds, actions.Deletes)
	// Output:
	// Add: [3], Delete: [4 5]
}

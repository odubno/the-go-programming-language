package tree

import (
	"fmt"
	"sort"
)

// Record is the format in which data is received
type Record struct {
	ID, Parent int
}

// Node helps branch is ensuing Child
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a nested family tree from a single slice of items
func Build(records []Record) (*Node, error) {
	// All records have a parent ID lower than their own ID, except for the root record,
	// which has a parent ID that's equal to its own ID.

	// if no records return nil
	if len(records) == 0 {
		return nil, nil
	}

	// Order records by ID. IDs act as parents to child nodes.
	// This ensures that all Parent nodes are created before adding Child nodes.
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// Create a map to help track results for each record
	family := map[int]*Node{}

	for index, record := range records {

		// The index should consecutively match the record
		if record.ID != index {
			return nil, fmt.Errorf("cycle dependency")
		} else if record.ID < record.Parent {
			return nil, fmt.Errorf("parent ID not in sequence")
		} else if record.ID > 0 && record.ID == record.Parent {
			return nil, fmt.Errorf("bad parent")
		}

		// Keep track of each node
		family[record.ID] = &Node{ID: record.ID}

		if record.ID > 0 {
			// ID is used as a Parent for the ensuing nodes.
			// Since the records are sorted by ID, the Parent node is created by preceding nodes.
			// The main key in the map is 0. The other keys are the children and they act as place holders,
			// until they're added to the main key
			p := family[record.Parent]
			p.Children = append(p.Children, family[record.ID])
		}
	}

	// 0 is the main key that stores the family tree
	return family[0], nil
}

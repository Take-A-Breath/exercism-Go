package tree

import (
	"errors"
	"sort"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 || records[0].ID != records[0].Parent {
		return &Node{}, errors.New("no root node")
	} else if records[0].ID < records[0].Parent {
		return &Node{}, errors.New("root node has parent")
	}

	for i := 1; i < len(records); i++ {
		if records[i].ID == records[i-1].ID {
			return &Node{}, errors.New("duplicate root")
		}
		if records[i].ID-records[i-1].ID != 1 {
			return &Node{}, errors.New("non-continuous")
		}
	}
	for _, v := range records {
		if v.Parent > v.ID {
			return &Node{}, errors.New("higher id parent of lower id")
		} else if v.ID != 0 && v.ID == v.Parent {
			return nil, errors.New("cycle directly")
		}
		nodes[v.ID] = &Node{ID: v.ID}
	}

	for _, v := range records {
		if v.ID == 0 {
			continue
		}
		if parent, ok := nodes[v.Parent]; ok {
			parent.Children = append(parent.Children, nodes[v.ID])
		}
	}
	return nodes[0], nil
}

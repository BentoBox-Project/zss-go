package zss

import (
	"fmt"
	"strings"
)

// Node is the unit of the graph to be compared
type Node struct {
	Label    string
	Children []*Node
}

// AddKid adds a given node to the Node's slice called Children
func (node *Node) AddKid(otherNode *Node, before bool) *Node {
	if before == false {
		var prependNodeSlice = []*Node{otherNode}
		node.Children = append(prependNodeSlice, node.Children...)
	} else {
		node.Children = append(node.Children, otherNode)
	}
	return node
}

// String retutns the label of the node
func (node *Node) String() string {
	s := fmt.Sprintf("%d:%s", len(node.Children), node.Label)
	nodeStrings := []string{}
	for _, c := range node.Children {
		nodeStrings = append(nodeStrings, c.String())
	}
	nodeStrings = append([]string{s}, nodeStrings...)
	s = strings.Join(nodeStrings, "\n")
	return s
}

// Contains returns true if the given label
func (node *Node) Contains(otherNode *Node) bool {
	if node.Label == otherNode.Label {
		return true
	}

	for _, n := range node.Children {
		if n.Label == otherNode.Label {
			return true
		}
	}
	return false
}

// ContainsWithLabel returns true if the given label is equals to the node's label
func (node *Node) ContainsWithLabel(otherNodeLabel string) bool {
	if node.Label == otherNodeLabel {
		return true
	}

	var containsList = []bool{}
	for _, n := range node.Children {
		containsList = append(containsList, n.ContainsWithLabel(otherNodeLabel))
	}

	for _, value := range containsList {
		if value == true {
			return true
		}
	}
	return false
}

// GetNode returns the label from a given Node
func GetNode(node *Node) string {
	return node.Label

}

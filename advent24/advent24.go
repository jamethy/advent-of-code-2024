package advent24

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	name        string
	leftSource  *Node
	rightSource *Node
	operation   string
	value       *bool
}

func (n *Node) GetValue() bool {
	if n.value != nil {
		return *n.value
	}
	leftValue := n.leftSource.GetValue()
	rightValue := n.rightSource.GetValue()

	var v bool
	switch n.operation {
	case "XOR":
		v = leftValue != rightValue
	case "OR":
		v = leftValue || rightValue
	case "AND":
		v = leftValue && rightValue
	}
	n.value = &v
	return *n.value
}

func Solution(inputFile string) (part1, part2 any) {
	fileParts := util.ReadFileSplitBy(inputFile, "\n\n")
	//lines := util.ReadFile(inputFile)

	initializationLines := strings.Split(fileParts[0], "\n")
	nodeLines := strings.Split(fileParts[1], "\n")
	nodes := make(map[string]*Node, len(nodeLines))

	for _, line := range nodeLines {
		parts := strings.Split(line, " ")
		leftNode, ok := nodes[parts[0]]
		if !ok {
			leftNode = &Node{name: parts[0]}
			nodes[parts[0]] = leftNode
		}
		rightNode, ok := nodes[parts[2]]
		if !ok {
			rightNode = &Node{name: parts[2]}
			nodes[parts[2]] = rightNode
		}
		node, ok := nodes[parts[4]]
		if !ok {
			node = &Node{name: parts[4]}
			nodes[parts[4]] = node
		}

		node.leftSource = leftNode
		node.rightSource = rightNode
		node.operation = parts[1]
	}

	for _, line := range initializationLines {
		parts := strings.Split(line, ": ")
		node := nodes[parts[0]]
		v := parts[1] == "1"
		node.value = &v
	}

	var zKeys []string
	for k := range nodes {
		if strings.HasPrefix(k, "z") {
			zKeys = append(zKeys, k)
		}
	}
	slices.Sort(zKeys)

	var u uint
	for _, k := range zKeys {
		n := nodes[k]
		pos, _ := strconv.Atoi(k[1:])
		value := n.GetValue()
		u = bitutil.SetBit(u, uint(pos), value)
	}

	return u, 0
}

func ptr[V any](v V) *V {
	return &v
}

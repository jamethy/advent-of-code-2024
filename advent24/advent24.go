package advent24

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"advent2024/util/mathutil"
	"fmt"
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
	default:
		panic("unrecognized operation")
	}
	n.value = &v
	return *n.value
}

func (n *Node) GetTree(indent string) string {
	sb := strings.Builder{}
	sb.WriteString(indent)
	sb.WriteString(n.name)

	if n.leftSource == nil {
		if *n.value {
			sb.WriteString("(1)")
			return sb.String()
		}
		sb.WriteString("(0)")
		return sb.String()
	}

	sb.WriteString("(\n")
	sb.WriteString(n.leftSource.GetTree(indent + " "))
	sb.WriteRune('\n')
	sb.WriteString(indent)
	sb.WriteRune(' ')
	sb.WriteString(n.operation)
	sb.WriteRune('\n')
	sb.WriteString(n.rightSource.GetTree(indent + " "))
	sb.WriteRune('\n')
	sb.WriteString(indent)
	sb.WriteRune(')')
	return sb.String()
}

func (n *Node) PrintInvolved() {
	involved := n.GetInvolved()
	slices.Sort(involved)
	fmt.Println(n.name)
	fmt.Println(strings.Join(involved, "\n"))
	fmt.Println()
}

func (n *Node) GetInvolved() []string {
	if n.leftSource == nil {
		return []string{}
	}
	involved := n.leftSource.GetInvolved()
	involved = append(involved, n.rightSource.GetInvolved()...)
	return append(involved, fmt.Sprintf("%s %s %s -> %s", n.leftSource.name, n.operation, n.rightSource.name, n.name))
}

func findZKeys(nodes map[string]*Node) []string {
	var zKeys []string
	for k := range nodes {
		if strings.HasPrefix(k, "z") {
			zKeys = append(zKeys, k)
		}
	}
	slices.Sort(zKeys)
	return zKeys
}

func Solution(inputFile string) (part1, part2 any) {

	nodes := parseNodes(inputFile)
	zKeys := findZKeys(nodes)

	u := runCalculation(nodes, zKeys)

	if inputFile == "sample.txt" {
		return u, nil
	}

	var swapped []string
	swapped = append(swapped, swapNodes("z16", "hmk", nodes)...) // z16
	swapped = append(swapped, swapNodes("z20", "fhp", nodes)...) // z20
	swapped = append(swapped, swapNodes("tpc", "rvf", nodes)...) // z27
	swapped = append(swapped, swapNodes("z33", "fcd", nodes)...) // z33
	for p := uint(1); p < 45; p++ {
		pp1 := uint(mathutil.IntPow(2, int(p)))
		setInputs(0, pp1, nodes)
		z1 := runCalculation(nodes, zKeys)

		if z1 != pp1 {
			fmt.Println(p, pp1, z1)
			nodes["z"+util.LeftPad(strconv.Itoa(int(p)), "0", 2)].PrintInvolved()
		}
	}
	slices.Sort(swapped)

	return u, strings.Join(swapped, ",")
}

func parseNodes(inputFile string) map[string]*Node {
	fileParts := util.ReadFileSplitBy(inputFile, "\n\n")

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

		if strings.HasPrefix(leftNode.name, "y") && strings.HasPrefix(rightNode.name, "x") {
			leftNode, rightNode = rightNode, leftNode
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
	return nodes
}

func swapNodes(a, b string, nodes map[string]*Node) []string {
	an := nodes[a]
	bn := nodes[b]
	an.leftSource, bn.leftSource = bn.leftSource, an.leftSource
	an.rightSource, bn.rightSource = bn.rightSource, an.rightSource
	an.operation, bn.operation = bn.operation, an.operation
	return []string{a, b}
}

func setInputs(x, y uint, nodes map[string]*Node) {
	for _, n := range nodes {
		switch n.name[0] {
		case 'y':
			pos, _ := strconv.Atoi(n.name[1:])
			v := bitutil.IsBitSet(y, uint(pos))
			n.value = &v
		case 'x':
			pos, _ := strconv.Atoi(n.name[1:])
			v := bitutil.IsBitSet(x, uint(pos))
			n.value = &v
		default:
			n.value = nil
		}
	}
}

func runCalculation(nodes map[string]*Node, zKeys []string) uint {
	var z uint
	for _, k := range zKeys {
		n := nodes[k]
		pos, _ := strconv.Atoi(k[1:])
		value := n.GetValue()
		z = bitutil.SetBit(z, uint(pos), value)
	}
	return z
}

func ptr[V any](v V) *V {
	return &v
}

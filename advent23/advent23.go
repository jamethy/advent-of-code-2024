package advent23

import (
	"advent2024/util"
	"advent2024/util/set"
	"slices"
	"strings"
)

type Node struct {
	name  string
	nodes []*Node
}

func (n *Node) String() string {
	return n.name
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	allNodes := make(map[string]*Node, len(lines))
	for _, line := range lines {
		nodeStrs := strings.Split(line, "-")
		aName, bName := nodeStrs[0], nodeStrs[1]

		a, ok := allNodes[aName]
		if !ok {
			a = &Node{name: aName}
			allNodes[aName] = a
		}

		b, ok := allNodes[bName]
		if !ok {
			b = &Node{name: bName}
			allNodes[bName] = b
		}

		a.nodes = append(a.nodes, b)
		b.nodes = append(b.nodes, a)
	}

	part1Networks := set.NewSet[string]()
	for _, n := range allNodes {
		if !strings.HasPrefix(n.name, "t") {
			continue
		}
		for _, a := range n.nodes {
			for _, b := range a.nodes {
				for _, c := range b.nodes {
					if c.name == n.name {
						networkName := []string{a.name, b.name, c.name}
						slices.Sort(networkName)
						part1Networks.Add(strings.Join(networkName, ","))
						break
					}
				}
			}
		}
	}

	return len(part1Networks), 0
}

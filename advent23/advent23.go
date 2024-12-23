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

func (n *Node) ConnectedDirectlyTo(other *Node) bool {
	if n == other {
		return true
	}
	for _, nn := range n.nodes {
		if nn == other {
			return true
		}
	}
	return false
}

type NetworkTriad struct {
	a, b, c *Node
}

func NewNetworkTriad(a, b, c *Node) NetworkTriad {
	if c.name < b.name {
		b, c = c, b
	}
	if b.name < a.name {
		a, b = b, a
	}
	if c.name < b.name {
		b, c = c, b
	}
	return NetworkTriad{a: a, b: b, c: c}
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

	networkTriads := set.NewSet[NetworkTriad]()
	for _, n := range allNodes {
		for _, a := range n.nodes {
			for _, b := range a.nodes {
				for _, c := range b.nodes {
					if c.name == n.name {
						networkTriads.Add(NewNetworkTriad(a, b, c))
						break
					}
				}
			}
		}
	}

	part1Count := 0
	for net := range networkTriads {
		for _, name := range []string{net.a.name, net.b.name, net.c.name} {
			if strings.HasPrefix(name, "t") {
				part1Count++
				break
			}
		}
	}

	triadsByNode := make(map[string][]NetworkTriad, len(allNodes))
	for net := range networkTriads {
		for _, name := range []string{net.a.name, net.b.name, net.c.name} {
			triadsByNode[name] = append(triadsByNode[name], net)
		}
	}

	var largestNetwork []NetworkTriad
	for _, n := range allNodes {
		nTriads := triadsByNode[n.name]
		if len(nTriads) <= len(largestNetwork) {
			continue
		}
		for i, ith := range nTriads {
			triadsConnectedToIth := []NetworkTriad{ith}
			iB, iC := ith.b, ith.c
			if iB == n {
				iB = ith.a
			} else if iC == n {
				iC = ith.a
			}
			for j, jth := range nTriads {
				if i == j {
					continue
				}
				connected := true
				for _, jNode := range []*Node{jth.a, jth.b, jth.c} {
					if jNode == n {
						continue
					}
					// is connected to both iB and iC
					if !jNode.ConnectedDirectlyTo(iB) {
						connected = false
						break
					}
					if !jNode.ConnectedDirectlyTo(iC) {
						connected = false
						break
					}
				}

				if connected {
					triadsConnectedToIth = append(triadsConnectedToIth, jth)
				}
			}
			if len(triadsConnectedToIth) > len(largestNetwork) {
				largestNetwork = triadsConnectedToIth
			}
		}
	}

	nodesInLargestNetwork := set.NewSet[string]()
	for _, t := range largestNetwork {
		nodesInLargestNetwork.Add(t.a.name)
		nodesInLargestNetwork.Add(t.b.name)
		nodesInLargestNetwork.Add(t.c.name)
	}
	networkName := nodesInLargestNetwork.Slice()
	slices.Sort(networkName)
	part2 = strings.Join(networkName, ",")

	return part1Count, part2
}

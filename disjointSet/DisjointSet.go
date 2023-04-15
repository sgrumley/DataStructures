package disjointSet

import "fmt"

type DisjointSet struct {
	rank   []int
	parent []int
	size   int
}

func New() DisjointSet {
	return DisjointSet{}
}

func (d *DisjointSet) Init(size int) {
	d.size = size + 1
	d.parent = append(d.parent, 0)
	d.rank = append(d.rank, 1)

	for i := 1; i < d.size; i++ {
		d.parent = append(d.parent, i)
		d.rank = append(d.rank, 1)
	}
}

// Path compression while searching list

func (d *DisjointSet) Find(value int) int {
	if value != d.parent[value] {
		d.parent[value] = d.Find(d.parent[value])
	}
	return d.parent[value]
}

// Union Set, appending smallest to biggest

func (d *DisjointSet) UnionSet(value1, value2 int) {
	parent1 := d.Find(value1)
	parent2 := d.Find(value2)

	if parent1 == parent2 {
		// Already in same set
		return
	}

	// Optimization using size
	if d.rank[parent1] > d.rank[parent2] {
		d.parent[parent2] = parent1
		// Update size
		d.rank[parent1] += d.rank[parent2]
	} else {
		d.parent[parent1] = parent2
		d.rank[parent2] += d.rank[parent1]
	}

}

func (d *DisjointSet) IsConnected(value1, value2 int) bool {
	if d.Find(value1) == d.Find(value2) {
		return true
	}
	return false
}

func (d *DisjointSet) PrintSet() {
	// print nodes
	fmt.Println("Nodes:")
	for i := 0; i < d.size; i++ {
		fmt.Print(i, " ")
	}

	// print parents
	fmt.Println("\nParents:")
	for i := 0; i < d.size; i++ {
		fmt.Print(d.parent[i], " ")
	}
	// print Root Node
	fmt.Println("\nRoot Node:")
	for i := 0; i < d.size; i++ {
		d.Find(i)
		fmt.Print(d.parent[i], " ")
	}
	// print Size
	fmt.Println("\nSize:")
	for i := 0; i < d.size; i++ {
		fmt.Print(d.rank[i], " ")
	}
	fmt.Println("\n------------------------------")
}

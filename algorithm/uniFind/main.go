package main

import "fmt"

type unionFind struct {
	parent []int
}

// 初始化并查集
func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &unionFind{parent: parent}
}

// 查找x所属的集合
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// 合并x和y所属的集合
func (uf *unionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		uf.parent[rootY] = rootX
	}
}

// 判断x和y是否属于同一个集合
func (uf *unionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func main() {
	uf := newUnionFind(10)
	uf.union(0, 1)
	uf.union(2, 3)
	uf.union(4, 5)
	uf.union(6, 7)
	uf.union(8, 9)
	uf.union(0, 9)
	fmt.Println("0 and 3 connected:", uf.connected(0, 3)) // false
	fmt.Println("0 and 9 connected:", uf.connected(0, 9)) // true
}

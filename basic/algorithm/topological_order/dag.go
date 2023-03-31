// DAG有向无环图的拓扑排序(采用邻接表存储)
package main

import "fmt"

// ArcNode is arc node
type ArcNode struct {
	adjvex int
	next   *ArcNode
}

// VNode is vertex node
type VNode struct {
	data  int
	first *ArcNode
}

// ALGraph is graph
type ALGraph struct {
	vertices       []VNode
	vexnum, arcnum int
}

const vexnum = 5

var graphRelation [vexnum][5]int

var al ALGraph

var inDegree [vexnum]int

// 入度为0的顶点
var emptyInDegreeVex []int

// 最后的拓扑排序结果
var topologicalVex []int

func initGraphRelation() {
	//有环图 vexnum = 6
	// graphRelation[0][0] = 2
	// graphRelation[0][1] = 4
	// graphRelation[1][0] = 5
	// graphRelation[2][0] = 6
	// graphRelation[2][1] = 5
	// graphRelation[3][0] = 2
	// graphRelation[4][0] = 4
	// graphRelation[5][0] = 6

	//无环图 vexnum = 5
	graphRelation[0][0] = 2
	graphRelation[0][1] = 4
	graphRelation[1][0] = 3
	graphRelation[1][1] = 4
	graphRelation[2][0] = 5
	graphRelation[3][0] = 3
	graphRelation[3][1] = 5
	graphRelation[4][0] = 0 //标识没有指向
}

// 初始化邻接表
func initALGraph() {
	for i, v := range graphRelation {
		node := VNode{data: i}
		arcnode := new(ArcNode)
		al.vexnum++
		for j, vv := range v {
			if vv != 0 {
				if j == 0 {
					arcnode = &ArcNode{adjvex: vv - 1}
					node.first = arcnode
				} else {
					arcnode.next = &ArcNode{adjvex: vv - 1}
				}
				al.arcnum++
			}
		}
		al.vertices = append(al.vertices, node)
	}
	fmt.Println(al)
}

// 初始化图的入度
func initGraphInDegree() {
	for _, v := range al.vertices {
		farc := v.first
		for {
			if farc == nil {
				break
			}
			inDegree[farc.adjvex]++
			farc = farc.next
		}
	}
}

// TopologicalSort is sort vex by topological rule
func TopologicalSort() {
	//将入度为0的顶点入栈
	for i, v := range inDegree {
		if v == 0 {
			emptyInDegreeVex = append(emptyInDegreeVex, i)
		}
	}
	count := 0
	for len(emptyInDegreeVex) > 0 {
		top := emptyInDegreeVex[0]
		topologicalVex = append(topologicalVex, top+1)
		//以点i为弧尾的边的弧头入度减一
		farc := al.vertices[top].first
		for {
			if farc == nil {
				break
			}
			inDegree[farc.adjvex]--
			if inDegree[farc.adjvex] == 0 {
				emptyInDegreeVex = append(emptyInDegreeVex, farc.adjvex)
			}
			farc = farc.next
		}
		count++
		emptyInDegreeVex = emptyInDegreeVex[1:]
	}
	if count == vexnum {
		fmt.Printf("topological sort order:%v\n", topologicalVex)
	} else {
		fmt.Println("topological sort fail, there is a loop graph")
	}
}

func main() {
	initGraphRelation()
	initALGraph()
	initGraphInDegree()
	TopologicalSort()
}

package substring_with_largest_variance

func largestVariance(s string) int {

	// 定义小根堆和大根堆
	smallh, bigh := NewMyheap(func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	}),
		NewMyheap(func(i, j int) bool {
			if i < j {
				return false
			}
			return true
		})
	res := 0
	var trackback func(i int)
	trackback = func(start int) {
		smallNode, bigNode := smallh.Peek(), bigh.Peek()
		if smallNode != nil && bigNode != nil && smallNode.Val != bigNode.Val {
			res = big(res, bigNode.Count-smallNode.Count)
		}
		for i := start; i < len(s); i++ {
			smallh.Add(int(s[i] - 'a'))
			bigh.Add(int(s[i] - 'a'))
			trackback(i + 1)
			smallh.Sub(int(s[i] - 'a'))
			bigh.Sub(int(s[i] - 'a'))
		}
	}
	trackback(0)
	return res
}

func big(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func NewMyheap(
	less func(i, j int) bool,
) *myheap {
	return &myheap{
		arr:  make([]*node, 0),
		less: less,
		m:    make(map[int]*node),
	}
}

type node struct {
	Val   int
	Count int
	Index int
}

// 实现自己的堆
type myheap struct {
	arr  []*node
	m    map[int]*node
	less func(i, j int) bool
}

func (h *myheap) Add(val int) {
	tmp, ok := h.m[val]
	if ok {
		tmp.Count++
		h.Swap(len(h.arr)-1, tmp.Index)
	} else {
		newNode := &node{
			Val:   val,
			Count: 1,
			Index: len(h.arr),
		}
		h.arr = append(h.arr, newNode)
		h.m[val] = newNode
	}
	h.up(len(h.arr) - 1)
}

func (h *myheap) Sub(val int) {
	tmp, ok := h.m[val]
	if ok {
		// 需要删除节点
		if tmp.Count == 1 {
			tmpIndex := tmp.Index
			lastIndex := len(h.arr) - 1
			h.Swap(tmp.Index, lastIndex)
			h.arr = h.arr[:lastIndex]
			delete(h.m, val)
			if tmpIndex < lastIndex {
				h.down(tmpIndex)
				h.up(tmpIndex)
			}
		} else {
			tmp.Count--
			h.down(tmp.Index)
			h.up(tmp.Index)
		}
	}
}

func (h *myheap) Peek() *node {
	if len(h.arr) == 0 {
		return nil
	}
	return h.arr[0]
}

func (h *myheap) up(index int) {
	for h.parent(index) >= 0 && h.less(h.arr[index].Count, h.arr[h.parent(index)].Count) {
		h.Swap(h.parent(index), index)
		index = h.parent(index)
	}
}

func (h *myheap) down(index int) {
	min := index
	for (h.leftChild(index) < len(h.arr) && h.less(h.arr[h.leftChild(index)].Count, h.arr[min].Count)) ||
		(h.rightChild(index) < len(h.arr) && h.less(h.arr[h.rightChild(index)].Count, h.arr[min].Count)) {
		if h.leftChild(index) < len(h.arr) && h.less(h.arr[h.leftChild(index)].Count, h.arr[min].Count) {
			min = h.leftChild(index)
		}
		if h.rightChild(index) < len(h.arr) && h.less(h.arr[h.rightChild(index)].Count, h.arr[min].Count) {
			min = h.rightChild(index)
		}
		if min == index {
			break
		}
		h.Swap(min, index)
		index = min
	}
}

func (h *myheap) parent(i int) int {
	return (i - 1) / 2
}
func (h *myheap) leftChild(i int) int {
	return i*2 + 1
}

func (h *myheap) rightChild(i int) int {
	return i*2 + 2
}

// 同时交换对应节点的索引
func (h *myheap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
	h.arr[i].Index, h.arr[j].Index = j, i
}

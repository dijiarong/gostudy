package array

// 环形数组

type CycleArray[T any] struct {
	arr   []T
	start int
	end   int
	count int
}

func NewCycleArray[T any](cap int) *CycleArray[T] {
	return &CycleArray[T]{
		arr:   make([]T, cap),
		start: 0,
		end:   0,
		count: 0,
	}
}

func (c *CycleArray[T]) resize(newSize int) {
	// 创建新的
	newArr := make([]T, newSize)
	// 复制
	for i := 0; i < c.count; i++ {
		newArr[i] = c.arr[(c.start+i)%len(c.arr)]
	}
	// 重置
	c.arr = newArr
	c.start = 0
	c.end = c.count
}

func (c *CycleArray[T]) AddFirst(val T) error {
	if c.IsFull() {
		c.resize(len(c.arr) * 2)
	}
	// start前移
	c.start = (c.start - 1 + len(c.arr)) % len(c.arr)
	// 赋值
	c.arr[c.start] = val
	c.count++
	return nil
}

func (c *CycleArray[T]) RemoveFirst() (T, error) {
	if c.isEmpty() {
		var t T
		return t, nil
	}
	// 取值
	val := c.arr[c.start]
	// start后移
	c.start = (c.start + 1) % len(c.arr)
	// 指向空
	c.arr[c.end] = *new(T)
	c.count--
	// 判断一下是否需要缩容
	if c.count != 0 && c.count == len(c.arr)/4 {
		c.resize(len(c.arr) / 2)
	}
	return val, nil
}

func (c *CycleArray[T]) AddLast(val T) error {
	if c.IsFull() {
		c.resize(len(c.arr) * 2)
	}
	// 赋值
	c.arr[c.end] = val
	// end后移
	c.end = (c.end + 1) % len(c.arr)
	c.count++
	return nil
}

func (c *CycleArray[T]) RemoveLast() (T, error) {
	if c.isEmpty() {
		var t T
		return t, nil
	}
	// end前移
	c.end = (c.end - 1 + len(c.arr)) % len(c.arr)
	// 取值
	val := c.arr[c.end]
	// 指向空
	c.arr[c.end] = *new(T)
	c.count--
	// 判断一下是否需要缩容
	if c.count != 0 && c.count == len(c.arr)/4 {
		c.resize(len(c.arr) / 2)
	}
	return val, nil
}

func (c *CycleArray[T]) IsFull() bool {
	return c.count == len(c.arr)
}

func (c *CycleArray[T]) isEmpty() bool {
	return c.count == 0
}

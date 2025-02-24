package lrucache

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.value, true
	}

	return 0, false
}

func (c *LRUCache) Put(key, value int) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.moveToHead(node)
		return
	}

	if len(c.cache) >= c.capacity {
		c.removeTail()
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	c.cache[key] = newNode
	c.moveToHead(newNode)
}

func (c *LRUCache) moveToHead(node *Node) {
	if node == c.head {
		return
	}

	// Remove from current position
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == c.tail {
		c.tail = node.prev
	}

	// Move to front
	node.next = c.head
	node.prev = nil
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
	if c.tail == nil {
		c.tail = node
	}
}

func (c *LRUCache) removeTail() {
	if c.tail == nil {
		return
	}

	delete(c.cache, c.tail.key)
	c.tail = c.tail.prev

	if c.tail != nil {
		c.tail.next = nil
	} else {
		c.head = nil
	}
}

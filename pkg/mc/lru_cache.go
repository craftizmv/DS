package mc

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Put(key int, val int) {

}

// Get - Returns -1 if key is not found.
func (l *LRUCache) Get(key int) int {
	return -1
}

func (l *LRUCache) Remove(key int) {

}

func (l *LRUCache) Size() int {
	return len(l.cache)
}

func (l *LRUCache) Add(key int, val int) {

}

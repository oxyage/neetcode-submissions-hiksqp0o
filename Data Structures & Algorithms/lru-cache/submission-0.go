// todo: Разобрать

/*
We want all operations to be O(1) while still following LRU (Least Recently Used) rules.

To do that, we combine:

Hash Map -> quickly find a node by its key in O(1).
Doubly Linked List -> quickly move nodes to the most recently used position and remove the least recently used node from the other end in O(1).
We keep:

The most recently used node near the right side.
The least recently used node near the left side.
Whenever we:

Get a key: move that node to the right (most recently used).
Put a key:
If it exists: update value and move it to the right.
If it's new:
If at capacity: remove the leftmost real node (LRU).
Insert the new node at the right.
Dummy left and right nodes make insert/remove logic cleaner.

Data Structures
A hash map cache that maps key -> node.
A doubly linked list with:
left dummy: before the least recently used node.
right dummy: after the most recently used node.
Helper: remove(node)
Unlink node from the list by connecting its prev and next nodes.
Helper: insert(node)
Insert node just before right (mark as most recently used).
get(key)
If key not in cache, return -1.
Otherwise:
Remove its node from the list.
Insert it again near right (mark as recently used).
Return the node's value.
put(key, value)
If key already exists:
Remove its old node from the list.
Create or update the node and store it in cache[key].
Insert the node near right.
If len(cache) > capacity:
Take the node right after left (this is LRU).
Remove it from the list.
Delete its key from the hash map.
This way, both get and put run in O(1) time, and the LRU policy is always maintained.


*/
type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	capacity int
	set map[int]*LRUNode
	left, right *LRUNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		capacity: capacity,
		set: make(map[int]*LRUNode),
		left: &LRUNode{},
		right: &LRUNode{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (this *LRUCache) remove(node *LRUNode) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *LRUNode) {
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev
}



func (this *LRUCache) Get(key int) int {
    if node, ok := this.set[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.set[key]; ok {
		this.remove(node)
		delete(this.set, key)
	}

	node := &LRUNode{
		key: key,
		value: value,
	}
	this.set[key] = node
	this.insert(node)

	if len(this.set) > this.capacity {
		lru := this.left.next
		this.remove(lru)
		delete(this.set, lru.key)
	}
}

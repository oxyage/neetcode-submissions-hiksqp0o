type Node struct {
	key int 
	val int
	prev *Node
	next *Node
}

type LRUCache struct {
	cap int 			// следим за объемом кеша
	cache map[int]*Node // храним указатели на ноды, чтобы получать за O(1)
	leastRU *Node // фиктивный узел, слева в списке, менее используемые
	mostRU *Node // фиктивный узел, справа в списке, более используемые
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
		cap: capacity,
		cache: make(map[int]*Node),
		leastRU: &Node{},
		mostRU: &Node{},
	}
	
	// инициализируем кеш с фиктивными узлами
	lruCache.leastRU.next = lruCache.mostRU // [leastRU].next -> [mostRU]
	lruCache.mostRU.prev = lruCache.leastRU // [leastRU] <- [mostRU].prev

	return lruCache
}

func (this *LRUCache) Get(key int) int {
    // если есть в мапе, вернем его и поместим узел в mostRecentUsed
	if node, ok := this.cache[key]; ok {
		// перемещаем узел в правую часть списка
		// удаляем из списка, вставляем в правую часть
		this.remove(node)
		this.insert(node)
		return node.val
	}	

	return -1
}

func (this *LRUCache) remove(node *Node) {
	// что значит удалить узел из двусвязного списка?
	// текущая картина
	// [prevNode] <- prev - [node] - next -> [nextNode] 
	// [prevNode] - next -> [node] <- prev - [nextNode] 

	prevNode, nextNode := node.prev, node.next

	// у prevNode поменять указатель next на nextNode
	// [prevNode] - next -> [nextNode]
	prevNode.next = nextNode

	// у nextNode поменять указатель prev на prevNode 
	// [prevNode] <- prev - [nextNode]
	nextNode.prev = prevNode
}

func (this *LRUCache) insert(node *Node) {
	// что значит вставить узел в двусвязный список
	// текущая картина
	// [prevNode] <- -> [mostRU] 

	// целевая картина
	// [prevNode] <- prev - [node] - next -> [mostRU]
	// [prevNode] - next -> [node] <- prev - [mostRU]

	// вставлять будем в правую часть, то есть в mostRecentUsed
	// 1. сохранить указатели на prevNode, mostRU
	tmpMostRUprev := this.mostRU.prev

	// на данном шаге
	// [prevNode] <- prev - [node] - next -> [mostRU]
	node.next = this.mostRU
	node.prev = this.mostRU.prev


	// [prevNode] - next -> [node] <- prev - [mostRU]
	tmpMostRUprev.next = node
	this.mostRU.prev = node
}


func (this *LRUCache) Put(key int, value int) {
	// если узел есть, его надо удалить
	if node, ok := this.cache[key]; ok {
		// удаляем из списка
		this.remove(node)
		// удаляем из мапы
		delete(this.cache, key)
	}

	newNode := &Node{
		key: key,
		val: value,
	}
	// вставляем в мапу
	this.cache[key] = newNode
	// вставляем в список
	this.insert(newNode)

	// если текущий размер мапы превышает установленный объем, то надо удалить leastRecentUsed
	if len(this.cache) > this.cap {
		leastUsedNode := this.leastRU.next
		// удалить из мапы
		delete(this.cache, leastUsedNode.key)
		// удалить из списка
		this.remove(leastUsedNode)
	}
}

type MinStack struct {
	items []int
	set map[int]int
	minItems []int
}

func Constructor() *MinStack {
	return &MinStack{
		items: make([]int, 0),
		set: make(map[int]int),
		minItems: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	// insert
	idx := len(this.items)
	this.items = append(this.items, val)
	this.set[idx] = val
	// fmt.Printf("pushed value %d to index %d\n", val, idx)

	// update min
	this.updateMin(val)

}

func (this *MinStack) Pop() {
	idx := len(this.items)	
	if idx == 0 {
		return 
	}

	delete(this.set, idx - 1)
	this.items = this.items[0:idx - 1]
	this.minItems = this.minItems[0:idx - 1]

	// fmt.Printf("Pop() idx: %d, items: %v, minItems: %v, set: %v\n", idx, this.items, this.minItems, this.set)
}

func (this *MinStack) Top() int {
	idx := len(this.items)	
	if idx == 0 {
		return -1
	}

	// fmt.Printf("Top() idx: %d, items: %v, minItems: %v, set: %v \n", idx, this.items, this.minItems, this.set)

	return this.items[idx - 1]
}

func (this *MinStack) GetMin() int {
	if len(this.minItems) == 0 {
		return -2
	}
	idx := len(this.minItems)
	minimum := this.minItems[idx - 1]

	// fmt.Printf("GetMin() minimum: %d idx: %d, items: %v, minItems: %v, set: %v\n", minimum, idx, this.items, this.minItems, this.set)

	return minimum
}

func (this *MinStack) updateMin(val int) {
	idx := len(this.minItems)
	if idx == 0 {
		this.minItems = []int{val}
		return
	}
	
	newMin := min(val, this.minItems[idx - 1])
	this.minItems = append(this.minItems, newMin)
	// fmt.Printf("update min: %d\n", newMin)
}

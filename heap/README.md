# Creating Heaps in go

## basics
- heaps are a basic data structure vs heapsort a sorting algorithm that basically creates a heap once finished ( and thus sorting our input array/list)
- useful for ex. priority queues, selection algorithms and graph algorithms ...
- can be thought of as a binary tree, all but the lowest levels full and last row fills from left to right. 
- **Max heap** largest key on top -> quickly able to "pop" that element
- **Min Heap** as expected smallest is on the root. 
- insert and extract from the heap are really heap: O(log n)

## using an array actually (for maxheap)
- instead of going building an actual binary-tree we can just store our heap in an array.
    - `a[0]` -> root, `a[1-2]` -> one Level lower, `a[3-6]` -> one level lower again, `a[7-14]` etc...
    - the **left node**  of any point is: `(2*index +1)`
    - the **right node** of any point is: `(2*index +2)`

## inserting into our heap (for maxheap)
- element gets inserted into as last place.
- if bigger than its parent-node we swap em.
- rinse and repeat till done. (till that element doesnt swap up or reached top)

## extract keys (for maxheap)
- we get the root, our target biggest value and extract it
- afterwards we put the LAST (lowest row, last element) node on the empty top
- now we compare the next row, if bigger node is bigger than our inserted Element we swap.
- rinse and repeat going down for that new element.

```go

// our processes we want to queue (bigger prio -> do first)
type Process struct{
	prio int
}

// our heap structure (max heap in this case)
type Heap struct{
	arr []Process
}


// public function to add a element to the heap
func (h *Heap) Insert(proc Process){
	h.arr =  append(h.arr, proc)
	h.heapifyUp(len(h.arr)-1)
}
// bring heap back into heap-state after a Input()
// does so by swapping with parent till uptop or not bigger anymore
func (h *Heap)heapifyUp(idx int){
	for h.arr[idx].prio > h.arr[parent(idx)].prio {			// while( node>parent )
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}


// public function to "pop()" the largest/root node
func (h *Heap) Extract() (Process, error) {
	length := len(h.arr) -1
	if length < 0 {
		return Process{}, fmt.Errorf("Heap is Empty, can not remove anything")
	}
	popElement := h.arr[0]
	h.arr[0] = h.arr[length]	// swap last element to first
	h.arr = h.arr[:length]		// remove last slice element (but does not reallocate in go if i understand correctly)

	h.heapifyDown(0)			// start our sort-shuffle from index 0
	return popElement, nil
}
// bring heap back into heap-state after a Extract()
// does so by potentially swapping with bigger child, moving down till bottom/no more swap
func (h *Heap)heapifyDown(idx int){
	current := idx
	last 	:= len(h.arr)-1
	l, r 	:= left(idx), right(idx)
	for l <= last {
		if l == last {
			current = l
		} else if h.arr[l].prio > h.arr[r].prio{
			current = l
		} else {
			current = r
		}
		if h.arr[idx].prio < h.arr[current].prio{
			h.swap(idx, current)
			idx = current
			l, r = left(idx) , right(idx)
		} else { return }
	}
}


/*
*	helpers
*/

// returns the equivalent parent/left/right node of our "thought off binary-tree"
func parent(idx int) int {
	return (idx -1) / 2
}

func left(idx int) int {
	return 2*idx +1
}

func right(idx int) int {
	return 2*idx +2
}

func (h *Heap)swap(i1 int, i2 int){
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}
```
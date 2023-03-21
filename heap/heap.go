package main

import (
	"fmt"
	"time"
)

func main (){
	fmt.Println("heap starting")
	h := &Heap{}
	createDailyTasks(h)
	fmt.Printf("full queue: %+v \n", *h)
	doTask(h)
	fmt.Printf("now empty que: %+v \n", *h)
	doTask(h)
}

func createDailyTasks(h *Heap){
	// just some testdata to fill in:
	h.Push(Process{"walking the dog       |",5,20})
	h.Push(Process{"feeding the cat       |",1,80})
	h.Push(Process{"cleaning the bathroom |",2,10})
	h.Push(Process{"cooking dinner        |",1,25})
	h.Push(Process{"shopping food         |",3,30})
	h.Push(Process{"going to work         |",1,70})
	h.Push(Process{"brushing teeth        |",1,85})
	h.Push(Process{"walk home             |",2,60})
	h.Push(Process{"doing work            |",4,65})
}

func doTask(h *Heap){
	task, err := h.Pop()
	if err != nil{
		fmt.Println("Nothing left to do!")
		return
	}
	fmt.Printf("starting %v with priority: %d. Will take %d seconds...\n",  task.name, task.prio, task.duration)
	time.Sleep(time.Duration(task.duration) * time.Second)
	doTask(h)
}



/*
*	Max-Heap Priority queue
*/

// our processes we want to queue (bigger prio -> do first)
type Process struct{
	name string
	duration int
	prio int		// only "important attribute"
}

// our heap structure (max heap in this case)
type Heap struct{
	arr []Process
}


// public function to add a element to the heap
func (h *Heap) Push(proc Process){
	h.arr =  append(h.arr, proc)
	h.heapifyUp(len(h.arr)-1)
}
// bring heap back into heap-state after Inputing element
// does so by swapping with parent till uptop or not bigger anymore
func (h *Heap) heapifyUp(idx int){
	for h.arr[idx].prio > h.arr[parent(idx)].prio {			// while( node>parent )
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}


// public function to "pop()" the root node(=largest). Error if empty
func (h *Heap) Pop() (Process, error) {
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
// bring heap back into heap-state after a Pop()
// does so by potentially swapping with bigger child, moving down till bottom/no more swap
func (h *Heap) heapifyDown(idx int){
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
*	helpers for Max-Heap
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
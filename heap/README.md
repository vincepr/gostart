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
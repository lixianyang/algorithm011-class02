package Week_02

// 暴力：生成丑数
func nthUglyNumber(n int) int {
	var num int
	visited := make(map[int]struct{}, n)
	heap := []int{1}
	for i := 0; i < n; i++ {
		num = heap[0]
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		heapify(heap, len(heap), 0)
		for _, prime := range []int{2, 3, 5} {
			tmp := num * prime
			if _, ok := visited[tmp]; ok {
				continue
			}
			visited[tmp] = struct{}{}
			heap = append(heap, tmp)
			c, p := len(heap)-1, len(heap)/2-1
			for p >= 0 && heap[c] < heap[p] {
				heap[c], heap[p] = heap[p], heap[c]
				c, p = p, (p-1)/2
			}
		}
	}
	return num
}
func heapify(arr []int, length, i int) {
	var min, l, r int
	for {
		min, l, r = i, 2*i+1, 2*i+2
		if l < length && arr[l] < arr[min] {
			min = l
		}
		if r < length && arr[r] < arr[min] {
			min = r
		}
		if min == i {
			break
		}
		arr[min], arr[i] = arr[i], arr[min]
		i = min
	}
}

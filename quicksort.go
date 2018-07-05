package learngo

/*
QuickSort sort the input array

Given input int array
8 1 3 2 5 4 7 6
							^
							pivot

Partitioning
---------->
1 3 2 5 4 [8] 7 6

swap [8] and 6 pivot
1 3 2 5 4 [6] 7 8

do begin <- [6] and [6] -> end
*/
func QuickSort(a []int) {
	if a == nil {
		return
	}
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)

		qsort(a, low, p-1)
		qsort(a, p+1, high)
	}
}

func partition(a []int, low, high int) int {
	p := a[high]
	i := low
	for j := low; j < high; j = j + 1 {
		if a[j] <= p {
			a[j], a[i] = a[i], a[j]
			i = i + 1
		}
	}
	a[i], a[high] = a[high], a[i]
	return i
}

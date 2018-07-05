package learngo

/*
MergeSort sort the input array

Divde the array into first half and second half

MergeSort the first half
MergeSort the second half
Merge the two halves sorted
*/
func MergeSort(a []int) {
	if a == nil {
		return
	}

	mSort(a, 0, len(a)-1)
}

func mSort(a []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mSort(a, low, mid)
		mSort(a, mid+1, high)
		merge(a, low, mid, high)
	}
}

func merge(a []int, low, mid, high int) {
	l, r := mid-low+1, high-mid
	tmpL := make([]int, l)
	tmpR := make([]int, r)
	copy(tmpL, a[low:])
	copy(tmpR, a[mid+1:])

	i, j, k := 0, 0, low
	for i < l && j < r {
		if tmpL[i] <= tmpR[j] {
			a[k] = tmpL[i]
			i = i + 1
		} else {
			a[k] = tmpR[j]
			j = j + 1
		}
		k = k + 1
	}

	for i < l {
		a[k] = tmpL[i]
		i, k = i+1, k+1
	}
	for j < r {
		a[k] = tmpR[j]
		j, k = j+1, k+1
	}
}

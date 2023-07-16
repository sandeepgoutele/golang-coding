package modules

import "log"

func sort012(arr []int) {
	l, m, r := -1, 0, len(arr)
	for m < r {
		switch arr[m] {
		case 0:
			l++
			arr[l], arr[m] = arr[m], arr[l]
			m++
		case 1:
			m++
		case 2:
			r--
			arr[m], arr[r] = arr[r], arr[m]
			m++
		}
	}
}

func Sort012() {
	arr := []int{0, 2, 1, 2, 0}
	sort012(arr)
	log.Print("{0 2 1 2 0}: ", arr)
	arr = []int{0, 1, 0}
	sort012(arr)
	log.Print("{0 1 0}: ", arr)
}

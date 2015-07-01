package sort

func Insert(list []interface{}, comparison func(va, vb interface{}) bool) {
	for i := 1; i < len(list); i++ {
		var start, end = 0, i
		var key = list[i]

		for start < end {
			var mid = (start + end) / 2

			if comparison(key, list[mid]) {
				end = mid
			} else {
				start = mid + 1
			}
		}

		for j := i; j > start; j-- {
			list[j] = list[j-1]
		}

		list[start] = key
	}
}

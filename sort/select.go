package sort

func Select(list []interface{}, comparison func(va, vb interface{}) bool) {
	for i := 0; i < len(list)-1; i++ {
		var pos = i

		for j := i + 1; j < len(list); j++ {
			if comparison(list[j], list[pos]) {
				pos = j
			}
		}

		list[pos], list[i] = list[i], list[pos]
	}
}

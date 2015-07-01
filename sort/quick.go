package sort

func Quick(list []interface{}, comparison func(va, vb interface{}) bool) {
	if len(list) <= 1 {
		return
	}

	mid, i := list[0], 1
	head, tail := 0, len(list)-1

	for head < tail {
		if comparison(list[i], mid) {
			list[i], list[tail] = list[tail], list[i]
			tail--
		} else {
			list[i], list[head] = list[head], list[i]
			head++
			i++
		}
	}

	list[head] = mid
	Quick(list[:head], comparison)
	Quick(list[head+1:], comparison)
}

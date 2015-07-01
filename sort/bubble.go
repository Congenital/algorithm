package sort

func Bubble(list []interface{}, comparison func(va, vb interface{}) bool) {
	for i := 0; i < len(list)-1; i++ {
		for j := len(list) - 1; j > i; j-- {
			if comparison(list[j], list[j-1]) {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

package utils

type OrderedIntArray []int

func (arr *OrderedIntArray) Append(value int) {
	pos := ff((*arr)[:], value, 0)
	if pos == len(*arr) {
		(*arr) = append(*arr, value)
	} else {
		*arr = append((*arr)[:pos+1], (*arr)[pos:]...)
		(*arr)[pos] = value
	}
}

func ff(arr []int, value int, margin int) int {
	if len(arr) == 1 {
		if arr[0] > value {
			return margin
		} else {
			return margin + 1
		}
	}

	midPoint := Middle(arr)
	if arr[midPoint] < value {
		return ff(arr[midPoint:], value, margin+midPoint)
	} else if arr[midPoint] > value {
		return ff(arr[:midPoint], value, margin)
	} else {
		return margin + midPoint
	}
}

package packsize

import (
	"sort"
)

var packSizes = []int{250, 500, 1000, 2000, 5000}

func NumberOfPacks(order int) map[int]int {
	result := map[int]int{}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	for order > 0 {
		for index, packSize := range packSizes {
			div := order / packSize
			if div > 0 {
				result[packSize] = div
				order = order - div*packSize
				if order == 0 {
					break
				}
			} else if index+2 == len(packSizes) && order > packSizes[index+1] {
				result[packSizes[index]] += 1
				order = 0
				break
			} else if order < packSizes[len(packSizes)-1] {
				result[packSizes[len(packSizes)-1]] += 1
				order = 0
				break
			}
		}
	}
	return result
}

func AddPackSize(packSize int) {
	packSizes = append(packSizes, packSize)
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
}

func RemovePackSize(packSize int) {
	for i, v := range packSizes {
		if v == packSize {
			packSizes = append(packSizes[:i], packSizes[i+1:]...)
		}
	}
}

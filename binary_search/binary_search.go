package binary_search

func BinarySearch(numberList []int, item int) int {

	lowNumber := 0
	highNumber := len(numberList) - 1

	for lowNumber <= highNumber {
		mid := (lowNumber + highNumber) / 2

		if numberList[mid] == item {
			return mid
		} else if numberList[mid] < item {
			lowNumber = mid + 1
		} else {
			highNumber = mid - 1
		}
	}

	return 0
}

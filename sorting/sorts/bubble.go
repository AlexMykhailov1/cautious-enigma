package sorts

// Bubble sorts integers using nested for loop
// Best complexity is O(n) when an array is already sorted
// In all other cases has complexity of O(n^2)
func Bubble(ints []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(ints)-1; i++ {
			j := i + 1
			// compare
			if ints[j] < ints[i] {
				ints[i], ints[j] = ints[j], ints[i]
				swapped = true
			}
		}
	}
	return ints
}

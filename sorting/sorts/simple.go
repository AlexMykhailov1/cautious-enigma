package sorts

// Simple sorts integers using nested for loop iteration.
// It has time complexity of O(n^2) in all cases.
func Simple(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			// compare
			if ints[j] < ints[i] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	return ints
}

package sorts

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Simple sort with random values",
			input: []int{376564, 132, 2345, 4321345, 135, -2346, 154, -9284, 235, 0},
			want:  []int{-9284, -2346, 0, 132, 135, 154, 235, 2345, 376564, 4321345},
		},
		{
			name:  "Simple sort with random values, some of them repeat",
			input: []int{-234, 10, 314, -213, 10, 3458, 2_359_043_245},
			want:  []int{-234, -213, 10, 10, 314, 3458, 2_359_043_245},
		},
		{
			name:  "Simple sort best, average && worst case: not sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Simple sort best, average && worst case: sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Simple sort with no values",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Simple sort with nil slice passed",
			input: nil,
			want:  nil,
		},
	}
	// run tests
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Simple(tc.input)
			require.Equal(t, tc.want, got)
		})
	}
}

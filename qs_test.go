package qs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = [][]int{
	{},
	{42},
	{85, 74, 43, 83, 31},
	{
		6, 76, 4, 79, 21, 87, 15, 43, 29, 98,
		19, 11, 57, 95, 43, 51, 59, 1, 23, 55,
		94, 84, 60, 98, 95,
	},
	{
		85, 24, 45, 21, 44, 91, 27, 56, 40, 20,
		89, 31, 75, 39, 99, 68, 78, 99, 96, 20,
		74, 98, 57, 49, 19, 6, 1, 89, 66, 46,
		86, 28, 82, 32, 40, 61, 68, 25, 77, 36,
		72, 82, 28, 34, 9, 87, 40, 90, 64, 78,
		65, 85, 68, 81, 72, 89, 26, 14, 16, 26,
		91, 34, 78, 100, 28, 55, 87, 70, 59, 61,
		50, 88, 41, 81, 80, 65, 1, 43, 5, 41,
		58, 99, 61, 42, 84, 92, 76, 69, 25, 42,
		95, 31, 46, 49, 44, 56, 6, 97, 30, 45,
	},
}

func TestQuickSort(t *testing.T) {
	for _, xs := range testData {
		t.Run(fmt.Sprintf("len() = %d", len(xs)), func(t *testing.T) {
			a := append(nums(nil), xs...)
			b := append(nums(nil), xs...)
			quicksort(nums(b))
			sort.Sort(a)
			assert.Equal(t, a, b)
		})
	}
}

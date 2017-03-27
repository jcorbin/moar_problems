package qs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: you can easily change or add testing data using `jot -r N` on a BSD
// system (TODO GNU equiv?)

var testData = []struct {
	name string
	xs   []int
}{
	{"empty", []int{}},
	{"the one", []int{42}},
	{"rand 5", []int{85, 74, 43, 83, 31}},
	{"rand 25", []int{
		6, 76, 4, 79, 21, 87, 15, 43, 29, 98,
		19, 11, 57, 95, 43, 51, 59, 1, 23, 55,
		94, 84, 60, 98, 95,
	}},
	{"rand 100", []int{
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
	}},
	{"tied median", []int{
		// this sample has a 5-element cohort tied for median
		44, 27, 62, 23, 33, 44, 12, 44, 60, 56,
		90, 96, 44, 29, 29, 74, 66, 18, 65, 2,
		97, 44, 33, 82, 13,
	}},
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			// the standard way
			a := append(nums(nil), tc.xs...)
			sort.Sort(a)

			// our way
			b := append(nums(nil), tc.xs...)
			quicksort(nums(b))

			// same?
			assert.Equal(t, a, b)
		})
	}
}

// // TODO: uncomment when implementing quickselect
// func TestQuickSelect(t *testing.T) {
// 	for _, pcts := range [][]float64{
// 		{0.5},
// 		{0.25, 0.5, 0.75},
// 		{0.25, 0.5, 0.75, 0.9, 0.95, 0.99},
// 	} {
// 		for _, tc := range testData {
// 			t.Run(tc.name, func(t *testing.T) {
// 				for _, pct := range pcts {
// 					t.Run(fmt.Sprintf("quickselect %0.2f", pct), func(t *testing.T) {
// 						i := int(pct * float64(len(tc.xs)))

// 						// the long way round
// 						a := append(nums(nil), tc.xs...)
// 						sort.Sort(a)

// 						// QET
// 						b := append(nums(nil), tc.xs...)
// 						quickselect(nums(b), i)

// 						// same?
// 						assert.Equal(t, a[i], b[i])
// 					})
// 				}
// 			})
// 		}
// 	}
// }

// // TODO: uncomment when implementing quickselectmany
// func TestQuickSelectMany(t *testing.T) {
// 	for _, pcts := range [][]float64{
// 		{0.5},
// 		{0.25, 0.5, 0.75},
// 		{0.25, 0.5, 0.75, 0.9, 0.95, 0.99},
// 	} {
// 		for _, tc := range testData {
// 			t.Run(tc.name, func(t *testing.T) {
// 				is := make([]int, len(pcts))
// 				for i, pct := range pcts {
// 					is[i] = int(pct * float64(len(tc.xs)))
// 				}

// 				// the long way round
// 				a := append(nums(nil), tc.xs...)
// 				sort.Sort(a)
// 				avs := make([]int, len(is))
// 				for i := range is {
// 					avs[i] = a[is[i]]
// 				}

// 				// QET
// 				b := append(nums(nil), tc.xs...)
// 				quickselectmany(nums(b), is...)
// 				bvs := make([]int, len(is))
// 				for i := range is {
// 					bvs[i] = b[is[i]]
// 				}

// 				// same?
// 				assert.Equal(t, avs, bvs)
// 			})
// 		}
// 	}
// }

// // TODO: uncomment when implementing quickselectfirst
// func TestQuickSelectFirst(t *testing.T) {
// 	for _, pct := range []float64{0.25, 0.5, 0.75, 0.9, 0.95, 0.99} {
// 		for _, tc := range testData {
// 			t.Run(tc.name, func(t *testing.T) {
// 				t.Run(fmt.Sprintf("quickselect %0.2f", pct), func(t *testing.T) {
// 					i := int(pct * float64(len(tc.xs)))

// 					// the long way round
// 					a := append(nums(nil), tc.xs...)
// 					sort.Sort(a)
// 					j := i
// 					for j > 0 && a[j-1] == a[j] {
// 						j--
// 					}

// 					// QET
// 					b := append(nums(nil), tc.xs...)
// 					k := quickselectfirst(nums(b), i)

// 					// same?
// 					assert.Equal(t, j, k)
// 					assert.Equal(t, a[j], b[k])
// 				})
// 			})
// 		}
// 	}
// }

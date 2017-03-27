package qs

import "sort"

type nums []int

func (s nums) Len() int           { return len(s) }
func (s nums) Less(i, j int) bool { return s[i] < s[j] }
func (s nums) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// quicksort sorts a sortable using the quick sort algorithm.
func quicksort(si sort.Interface) {
	q := [][2]int{{0, si.Len()}}
	for len(q) > 0 {
		i := len(q) - 1
		lo, hi := q[i][0], q[i][1]
		q = q[:i]
		if hi-lo <= 1 {
			continue
		}
		pi := lo/2 + hi/2
		pi = partition(si, lo, pi, hi-1)
		if pi-lo > 1 {
			q = append(q, [2]int{lo, pi})
		}
		if r := pi + 1; hi-r > 1 {
			q = append(q, [2]int{r, hi})
		}
	}
}

// partition is the core of the quicksort family.
func partition(si sort.Interface, lo, pi, hi int) int {
	si.Swap(pi, hi)
	i := lo
	for j := lo; j < hi; j++ {
		if si.Less(j, hi) {
			si.Swap(i, j)
			i++
		}
	}
	si.Swap(hi, i)
	return i
}

// TODO
// // quickselect does enough partial sorting to fix the position of the i-th
// // element in the final sorted order.
// func quickselect(si sort.Interface, i int)

// TODO
// // quickselectmany is just quickselect for more than one index in a single
// // pass.
// func quickselectmany(si sort.Interface, is ...int)

// TODO
// // quickselectfirst is like quickselect, except after fixing the given index it
// // seeks backwards over all values equal to the fixed one (in sorted order),
// // returning the index of the first occurrence.
// func quickselectfirst(si sort.Interface, i int) int

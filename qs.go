package qs

import "sort"

type nums []int

func (s nums) Len() int           { return len(s) }
func (s nums) Less(i, j int) bool { return s[i] < s[j] }
func (s nums) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

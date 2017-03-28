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

// quickselect does enough partial sorting to fix the position of the i-th
// element in the final sorted order.
func quickselect(si sort.Interface, i int) {
	lo, hi := 0, si.Len()
	if lo == hi {
		return
	}
	if i < lo || i >= hi {
		panic("i out of range")
	}
	for {
		pi := lo/2 + hi/2
		pi = partition(si, lo, pi, hi-1)
		if i == pi {
			return
		}
		if i < pi {
			if pi-lo <= 1 {
				return
			}
			hi = pi
		} else if i > pi {
			r := pi + 1
			if hi-r <= 1 {
				return
			}
			lo = r
		}
	}
}

type quickCore struct {
	q [][2]int

	// known positions
	k []bool

	// TODO si sort.Interface
}

func (qs *quickCore) pop() (int, int) {
	i := len(qs.q) - 1
	lo, hi := qs.q[i][0], qs.q[i][1]
	qs.q = qs.q[:i]
	return lo, hi
}

func (qs *quickCore) fix(si sort.Interface, lo, pi, hi int) int {
	pi = partition(si, lo, pi, hi-1)
	qs.k[pi] = true
	return pi
}

// quickselectmany is just quickselect for more than one index in a single
// pass.
func quickselectmany(si sort.Interface, is ...int) {
	n := len(is)
	l := si.Len()
	qs := quickCore{
		q: [][2]int{{0, l}},
		k: make([]bool, l),
	}

	for len(qs.q) > 0 {
		lo, hi := qs.pop()
		if hi-lo <= 1 {
			continue
		}
		pi := qs.fix(si, lo, lo/2+hi/2, hi)

		if is[0] > pi {
			// all are to the right
			// XXX
			continue
		}

		if j := len(is) - 1; is[j] < pi {
			// all are to the left
			// XXX
			continue
		}

		needLeft, needRight := false, false

		for _, i := range is {
			if i == pi {
				n--
				if n <= 0 {
					return
				}
			}
			if i < pi {
				needLeft = true
			}
			if i > pi {
				needRight = true
			}
		}

		if needLeft {
			if pi-lo <= 1 {
				n--
				if n <= 0 {
					return
				}
			}
			qs.q = append(qs.q, [2]int{lo, pi})
		}

		if needRight {
			r := pi + 1
			if hi-r <= 1 {
				n--
				if n <= 0 {
					return
				}
			}
			qs.q = append(qs.q, [2]int{r, hi})
		}

	}

}

// TODO
// // quickselectfirst is like quickselect, except after fixing the given index it
// // seeks backwards over all values equal to the fixed one (in sorted order),
// // returning the index of the first occurrence.
// func quickselectfirst(si sort.Interface, i int) int

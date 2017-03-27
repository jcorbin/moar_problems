A Go practice problem based on quicksort.

An implementation and test for `quicksort` is given.

Tests for some `quicksort` cousins are given:

- `quickselect` should select a single value from the sorted order
- `quickselectmany` should select several values from the sorted order
- `quickselectfirst` should select the first occurrence of a value, returning
  its index; in other words it's `quickselect` that finds the position of the
  first value tied for the given position.

For some color on `quickselectfirst`, it's an efficient solution to "split an
array in roughly half, but don't split tied values between the halves".

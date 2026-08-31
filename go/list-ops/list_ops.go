package listops

// IntList is a list of integers with custom operations.
type IntList []int

// Foldl reduces the list from left to right.
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
	for i := range s {
		result = fn(initial, s[i])
		initial = result
	}
	return result
}

// Foldr reduces the list from right to left.
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
	N := s.Length()
	for i := N - 1; i >= 0; i-- {
		result = fn(s[i], initial)
		initial = result
	}
	return result
}

// Filter keeps elements that satisfy fn.
func (s IntList) Filter(fn func(int) bool) IntList {
	// Count matches to allocate the result size.
	count := 0
	for _, val := range s {
		if fn(val) {
			count++
		}
	}

	result := make(IntList, count)

	// Copy matching elements into the result.
	j := 0
	for _, val := range s {
		if fn(val) {
			result[j] = val
			j++
		}
	}
	return result
}

func (s IntList) Length() int {
	count := 0
	for range s {
		count++
	}
	return count
}

// Map applies fn to every element.
func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, s.Length())
	for i, val := range s {
		result[i] = fn(val)
	}
	return result
}

func (s IntList) Reverse() IntList {
	N := s.Length()
	for i := range N / 2 {
		s[i], s[N-1-i] = s[N-1-i], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	total_len := s.Length() + lst.Length()
	appended := make(IntList, total_len)
	copy(appended, s)

	for j := range lst {
		appended[s.Length()+j] = lst[j]
	}
	return appended
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
	for _, list := range lists {
		result = result.Append(list)
	}
	return result
}

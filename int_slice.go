package jsarrayext

// IntSlice ...
type IntSlice []int

// Every ...
func (s IntSlice) Every(fn func(element int, index int) bool) bool {
	return every(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}

// Fill fills elements from a start index to an end index (not included) with a
// static value.
func (s IntSlice) Fill(value int, start int, end int) IntSlice {
	return fill(s, value, start, end).(IntSlice)
}

// Filter ...
func (s IntSlice) Filter(fn func(element int, index int) bool) IntSlice {
	return filter(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	}).(IntSlice)
}

// Find ...
func (s IntSlice) Find(fn func(element int, index int) bool) int {
	r := find(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})

	if r == nil {
		return 0
	}
	return r.(int)
}

// FindIndex ...
func (s IntSlice) FindIndex(fn func(element int, index int) bool) int {
	return findIndex(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}

// ForEach ...
func (s IntSlice) ForEach(fn func(element int, index int)) {
	forEach(s, func(element interface{}, index int) {
		fn(element.(int), index)
	})
}

// Map ...
func (s IntSlice) Map(fn func(element int, index int) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int) interface{} {
		return fn(element.(int), index)
	})
}

// Some ...
func (s IntSlice) Some(fn func(element int, index int) bool) bool {
	return some(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}

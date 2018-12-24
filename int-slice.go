package jsarrayext

// IntSlice ...
type IntSlice []int

// Every ...
func (s IntSlice) Every(fn func(element int, index int, slice IntSlice) bool) bool {
	return every(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element.(int), index, slice.(IntSlice))
	})
}

// Filter ...
func (s IntSlice) Filter(fn func(element int, index int, slice IntSlice) bool) IntSlice {
	return filter(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element.(int), index, slice.(IntSlice))
	}).(IntSlice)
}

// Find ...
func (s IntSlice) Find(fn func(element int, index int, slice IntSlice) bool) int {
	return find(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element.(int), index, slice.(IntSlice))
	}).(int)
}

// FindIndex ...
func (s IntSlice) FindIndex(fn func(element int, index int, slice IntSlice) bool) int {
	return findIndex(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element.(int), index, slice.(IntSlice))
	})
}

// ForEach ...
func (s IntSlice) ForEach(fn func(element int, index int, slice IntSlice)) {
	forEach(s, func(element interface{}, index int, slice interface{}) {
		fn(element.(int), index, slice.(IntSlice))
	})
}

// Map ...
func (s IntSlice) Map(fn func(element int, index int, slice IntSlice) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int, slice interface{}) interface{} {
		return fn(element.(int), index, slice.(IntSlice))
	})
}

// Some ...
func (s IntSlice) Some(fn func(element int, index int, slice IntSlice) bool) bool {
	return some(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element.(int), index, slice.(IntSlice))
	})
}

package jsarrayext

// Slice ...
type Slice []interface{}

// Every ...
func (s Slice) Every(fn func(element interface{}, index int) bool) bool {
	return every(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// Filter ...
func (s Slice) Filter(fn func(element interface{}, index int) bool) Slice {
	return filter(s, func(element interface{}, index int) bool {
		return fn(element, index)
	}).(Slice)
}

// Find ...
func (s Slice) Find(fn func(element interface{}, index int) bool) interface{} {
	return find(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// FindIndex ...
func (s Slice) FindIndex(fn func(element interface{}, index int) bool) int {
	return findIndex(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// ForEach ...
func (s Slice) ForEach(fn func(element interface{}, index int)) {
	forEach(s, func(element interface{}, index int) {
		fn(element, index)
	})
}

// Map ...
func (s Slice) Map(fn func(element interface{}, index int) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int) interface{} {
		return fn(element, index)
	})
}

// Some ...
func (s Slice) Some(fn func(element interface{}, index int) bool) bool {
	return some(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

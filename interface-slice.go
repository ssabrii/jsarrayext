package jsarrayext

// Slice ...
type Slice []interface{}

// Every ...
func (s Slice) Every(fn func(element interface{}, index int, slice Slice) bool) bool {
	return every(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element, index, slice.(Slice))
	})
}

// Filter ...
func (s Slice) Filter(fn func(element interface{}, index int, slice Slice) bool) Slice {
	return filter(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element, index, slice.(Slice))
	}).(Slice)
}

// Find ...
func (s Slice) Find(fn func(element interface{}, index int, slice Slice) bool) interface{} {
	return find(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element, index, slice.(Slice))
	})
}

// FindIndex ...
func (s Slice) FindIndex(fn func(element interface{}, index int, slice Slice) bool) int {
	return findIndex(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element, index, slice.(Slice))
	})
}

// ForEach ...
func (s Slice) ForEach(fn func(element interface{}, index int, slice Slice)) {
	forEach(s, func(element interface{}, index int, slice interface{}) {
		fn(element, index, slice.(Slice))
	})
}

// Map ...
func (s Slice) Map(fn func(element interface{}, index int, slice Slice) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int, slice interface{}) interface{} {
		return fn(element, index, slice.(Slice))
	})
}

// Some ...
func (s Slice) Some(fn func(element interface{}, index int, slice Slice) bool) bool {
	return some(s, func(element interface{}, index int, slice interface{}) bool {
		return fn(element, index, slice.(Slice))
	})
}

package jsarrayext

// Slice ...
type Slice []interface{}

// From ...
func From(slice []interface{}) Slice {
	return Slice(slice)
}

// Every ...
func (s Slice) Every(fn func(interface{}, int, Slice) bool) bool {
	for i, v := range s {
		if fn(v, i, s) == false {
			return false
		}
	}

	return true
}

// Filter ...
func (s Slice) Filter(fn func(interface{}, int, Slice) bool) Slice {
	filteredSlice := make([]interface{}, 0, len(s))
	s.ForEach(func(element interface{}, index int, slice Slice) {
		if fn(element, index, slice) == true {
			filteredSlice = append(filteredSlice, element)
		}
	})

	return filteredSlice
}

// Find ...
func (s Slice) Find(fn func(interface{}, int, Slice) bool) interface{} {
	if index := s.FindIndex(fn); index != -1 {
		return s[index]
	}

	return nil
}

// FindIndex ...
func (s Slice) FindIndex(fn func(interface{}, int, Slice) bool) int {
	for i, v := range s {
		if fn(v, i, s) == true {
			return i
		}
	}

	return -1
}

// ForEach ...
func (s Slice) ForEach(fn func(interface{}, int, Slice)) {
	for i, v := range s {
		fn(v, i, s)
	}
}

// Map ...
func (s Slice) Map(fn func(interface{}, int, Slice) interface{}) Slice {
	newSlice := make([]interface{}, len(s))
	for i, v := range s {
		newSlice[i] = fn(v, i, s)
	}

	return newSlice
}

// Some ...
func (s Slice) Some(fn func(interface{}, int, Slice) bool) bool {
	for i, v := range s {
		if fn(v, i, s) == true {
			return true
		}
	}

	return false
}

package jsarrayext

import "reflect"

func every(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index, slice) == false {
			return false
		}
	}

	return true
}

func filter(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) bool,
) interface{} {
	filtered := reflect.MakeSlice(reflect.TypeOf(slice), 0, reflect.ValueOf(slice).Len())

	forEach(slice, func(element interface{}, index int, slice interface{}) {
		if fn(element, index, slice) == true {
			filtered = reflect.AppendSlice(filtered, reflect.ValueOf(element))
		}
	})

	return filtered.Interface()
}

func find(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) bool,
) interface{} {
	if index := findIndex(slice, fn); index != -1 {
		element := reflect.ValueOf(slice).Index(index).Interface()
		return element
	}

	return nil
}

func findIndex(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) bool,
) int {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index, slice) == true {
			return index
		}
	}

	return -1
}

func forEach(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}),
) {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		fn(element, index, slice)
	}
}

func mapToInterfaceSlice(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) interface{},
) []interface{} {
	mapped := make([]interface{}, reflect.ValueOf(slice).Len())

	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		mapped[index] = fn(element, index, slice)
	}

	return mapped
}

func some(
	slice interface{},
	fn func(element interface{}, index int, slice interface{}) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index, slice) == true {
			return true
		}
	}

	return false
}

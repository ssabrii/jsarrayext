package jsarrayext

import "reflect"

func every(
	slice interface{},
	fn func(element interface{}, index int) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == false {
			return false
		}
	}

	return true
}

func filter(
	slice interface{},
	fn func(element interface{}, index int) bool,
) interface{} {
	filtered := reflect.MakeSlice(reflect.TypeOf(slice), 0, reflect.ValueOf(slice).Len())

	forEach(slice, func(element interface{}, index int) {
		if fn(element, index) == true {
			if element == nil {
				// reflect.Append(filtered, nil) won't work.
				nilSlice := reflect.MakeSlice(reflect.TypeOf(slice), 1, 1)
				filtered = reflect.AppendSlice(filtered, nilSlice)
			} else {
				filtered = reflect.Append(filtered, reflect.ValueOf(element))
			}
		}
	})

	return filtered.Interface()
}

func find(
	slice interface{},
	fn func(element interface{}, index int) bool,
) interface{} {
	if index := findIndex(slice, fn); index != -1 {
		element := reflect.ValueOf(slice).Index(index).Interface()
		return element
	}

	return nil
}

func findIndex(
	slice interface{},
	fn func(element interface{}, index int) bool,
) int {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == true {
			return index
		}
	}

	return -1
}

func forEach(
	slice interface{},
	fn func(element interface{}, index int),
) {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		fn(element, index)
	}
}

func mapToInterfaceSlice(
	slice interface{},
	fn func(element interface{}, index int) interface{},
) []interface{} {
	mapped := make([]interface{}, reflect.ValueOf(slice).Len())

	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		mapped[index] = fn(element, index)
	}

	return mapped
}

func some(
	slice interface{},
	fn func(element interface{}, index int) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == true {
			return true
		}
	}

	return false
}

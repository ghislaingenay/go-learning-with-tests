package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// This is a closure designed to take a nested reflect.Value, 
	// extract its raw Go interface using .Interface()
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
		// Objects ans structures
	case reflect.Struct:
		// NumField: Counts how many fields the struct has.
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		// MapKeys: Returns a slice of reflect.Value objects representing the keys in the map.
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			// Recv: Reads data from the channel dynamically.
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
		case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}


func getValue(x interface{}) reflect.Value {
	// Converts the empty interface interface{} into a reflect.Value object, 
	// allowing Go's reflection engine to inspect its type and structure.
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		// get the actual value the pointer is pointing to
		val = val.Elem()
	}

	return val
}
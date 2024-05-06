package gempty

import (
	"reflect"
)

type UnsupportedError struct {
	kind reflect.Kind
}

func (e *UnsupportedError) Error() string {
	return "unsupported kind: " + e.kind.String()
}

// Clone creates an empty "copy" of the calues passed.
// A new instance containing default value will be returned. or an error
// if unable to clone the passed type.
func Clone[T comparable](s T) (T, error) {
	var cln T
	sType := reflect.TypeOf(s)
	switch kind := sType.Kind(); kind {
	case reflect.Slice:
		sTypElm := sType.Elem()
		cln = reflect.MakeSlice(reflect.SliceOf(sTypElm), 0, 0).Interface().(T)
	case reflect.Map:
		sTypElm := sType.Elem() // deref
		sKeyTypElm := sType.Key()
		cln = reflect.MakeMap(reflect.MapOf(sKeyTypElm, sTypElm)).Interface().(T)
	case reflect.Chan, reflect.Func:
		return cln, &UnsupportedError{kind}
	default:
		var n reflect.Value
		if sType.Kind() == reflect.Ptr {
			n = reflect.New(sType.Elem())
		} else {
			n = reflect.New(sType).Elem()
		}
		cln = n.Interface().(T)
	}
	return cln, nil
}

// IsPtr will return true, if the passed value is a pointer
func IsPtr[T comparable](v T) bool {
	return reflect.TypeOf(v).Kind() == reflect.Pointer
}

// TODO: make this work!
// func AsPtr[T comparable](v T) T {
// 	if !IsPtr(v) {
// 		return reflect.PointerTo(v).(T)
// 	}
// 	return v
// }

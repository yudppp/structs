package structs

import (
	"reflect"
	"strconv"
	"strings"
)

// NewExample is returnned including example value
func NewExample(in interface{}) interface{} {
	return parse(reflect.TypeOf(in), "example").Interface()
}

// NewDefault is returnned including default value
func NewDefault(in interface{}) interface{} {
	return parse(reflect.TypeOf(in), "default").Interface()
}

func parse(rt reflect.Type, tagName string) reflect.Value {
	rv := reflect.New(rt).Elem()
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = reflect.New(rt).Elem()
		defer func() {
			rv = rv.Addr()
		}()
	}
	if rt.Kind() != reflect.Struct {
		rv = parseWithValue(rt, tagName, "")
		return rv
	}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		val := reflect.ValueOf(field.Tag.Get(tagName))

		if field.Type.Kind() == reflect.Ptr {
			fv := parseWithValue(field.Type, tagName, val.String())
			fvPtr := reflect.New(fv.Type())
			fvPtr.Elem().Set(fv)
			rv.Field(i).Set(fvPtr)
		} else {
			fv := parseWithValue(field.Type, tagName, val.String())
			rv.Field(i).Set(fv)
		}
	}
	return rv
}

func parseWithValue(rt reflect.Type, tagName string, value string) reflect.Value {
	rv := reflect.New(rt).Elem()
	if isUncommonKind(rt.Kind()) &&
		rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = reflect.New(rt).Elem()
		defer func() {
			rv = rv.Addr()
		}()
	}
	switch rt.Kind() {
	case reflect.String:
		rv.SetString(value)
	case reflect.Bool:
		b, _ := strconv.ParseBool(value)
		rv.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, _ := strconv.ParseInt(value, 10, 64)
		rv.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, _ := strconv.ParseUint(value, 10, 64)
		rv.SetUint(n)
	case reflect.Float64, reflect.Float32:
		x, _ := strconv.ParseFloat(value, 0)
		rv.SetFloat(x)
	case reflect.Struct:
		rv.Set(parse(rt, tagName))
	case reflect.Slice:
		sl := reflect.MakeSlice(rt, 0, 1)
		if isCommonKind(rt.Elem().Kind()) {
			if value != "" {
				vals := strings.Split(value, ",")
				for _, val := range vals {
					sl = reflect.Append(sl, parseWithValue(rt.Elem(), tagName, val))
				}
			}
		} else if rt.Elem().Kind() == reflect.Ptr {
			if isCommonKind(rt.Elem().Elem().Kind()) {
				if value != "" {
					vals := strings.Split(value, ",")
					for _, val := range vals {
						fv := parseWithValue(rt.Elem(), tagName, val)
						fvPtr := reflect.New(fv.Type())
						fvPtr.Elem().Set(fv)
						sl = reflect.Append(sl, fvPtr)
					}
				}
			} else if rt.Elem().Kind() != reflect.Interface {
				fv := parse(rt.Elem(), tagName)
				fvPtr := reflect.New(fv.Type())
				fvPtr.Elem().Set(fv)
				sl = reflect.Append(sl, fvPtr)
			}
		} else if rt.Elem().Kind() != reflect.Interface {
			sl = reflect.Append(sl, parse(rt.Elem(), tagName))
		}
		rv.Set(sl)
	}
	return rv
}

func isCommonKind(kind reflect.Kind) bool {
	commons := []reflect.Kind{
		reflect.String,
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
	}
	for _, v := range commons {
		if v == kind {
			return true
		}
	}
	return false
}

func isUncommonKind(kind reflect.Kind) bool {
	uncommons := []reflect.Kind{
		reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice,
	}
	for _, v := range uncommons {
		if v == kind {
			return true
		}
	}
	return false
}

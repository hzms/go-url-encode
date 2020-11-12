package urlencode

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func Encode(valueMap map[string]interface{}) string {
	data := url.Values{}
	kv := make(map[string]string)
	for key, value := range valueMap {
		v := reflect.ValueOf(value)
		kind := v.Kind()
		switch kind {
		case reflect.String:
			kv[key] = value.(string)
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			kv[key] = fmt.Sprint(value)
		case reflect.Map:
			appendFromMap(kv, key, v)
		case reflect.Slice, reflect.Array:
			appendFromArray(kv, key, v)
		default:
			data.Set(key, "")
		}
	}
	for k, v := range kv {
		data.Set(k, v)
	}
	return data.Encode()
}

func appendFromMap(data map[string]string, keyPrefix string, m reflect.Value) map[string]string {
	for _, key := range m.MapKeys() {
		v := m.MapIndex(key)
		kind := v.Kind()
		newKey := keyPrefix + "[" + key.String() + "]"
		switch kind {
		case reflect.String:
			data[newKey] = v.String()
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			data[newKey] = fmt.Sprint(v.Int())
		case reflect.Map:
			appendFromMap(data, newKey, v)
		case reflect.Slice, reflect.Array:
			appendFromArray(data, newKey, v)
		default:
			data[newKey] = ""
		}
	}
	return data
}

func appendFromArray(data map[string]string, keyPrefix string, m reflect.Value) map[string]string {
	len := m.Len()
	for i := 0; i < len; i++ {
		v := m.Index(i)
		kind := v.Kind()
		newKey := keyPrefix + "[" + strconv.Itoa(i) + "]"
		switch kind {
		case reflect.String:
			data[newKey] = v.String()
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			data[newKey] = fmt.Sprint(v.Int())
		case reflect.Map:
			appendFromMap(data, newKey, v)
		case reflect.Slice, reflect.Array:
			appendFromArray(data, newKey, v)
		default:
			data[newKey] = ""
		}
	}
	return data
}

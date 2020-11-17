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
		appendKeyValue(kv, key, v)
	}
	for k, v := range kv {
		data.Set(k, v)
	}
	return data.Encode()
}

func appendKeyValue(data map[string]string, key string, value reflect.Value) map[string]string {
	kind := value.Kind()
	switch kind {
	case reflect.String:
		data[key] = value.String()
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		data[key] = fmt.Sprint(value.Int())
	case reflect.Map:
		appendFromMap(data, key, value)
	case reflect.Slice, reflect.Array:
		appendFromArray(data, key, value)
	case reflect.Struct:
		appendFromStruct(data, key, value)
	case reflect.Ptr:
		pv := reflect.Indirect(value)
		if pv.Kind() == reflect.Struct {
			appendFromStruct(data, key, pv)
		}
	default:
		data[key] = ""
	}
	return data
}

func appendFromMap(data map[string]string, keyPrefix string, m reflect.Value) map[string]string {
	for _, key := range m.MapKeys() {
		v := m.MapIndex(key)
		newKey := keyPrefix + "[" + key.String() + "]"
		appendKeyValue(data, newKey, v)
	}
	return data
}

func appendFromArray(data map[string]string, keyPrefix string, arr reflect.Value) map[string]string {
	l := arr.Len()
	for i := 0; i < l; i++ {
		v := arr.Index(i)
		newKey := keyPrefix + "[" + strconv.Itoa(i) + "]"
		appendKeyValue(data, newKey, v)
	}
	return data
}

func appendFromStruct(data map[string]string, keyPrefix string, s reflect.Value) map[string]string {
	// attrNum := s.NumField()
	t := s.Type()
	attrNum := t.NumField()
	for i := 0; i < attrNum; i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get("urlencode")
		name := field.Name
		if tagValue != "" {
			name = tagValue
		}
		newKey := keyPrefix + "[" + name + "]"
		appendKeyValue(data, newKey, s.Field(i))
	}
	return data
}

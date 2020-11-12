package main

import (
	"fmt"
	"urlencode"
)

func main() {
	data := map[string]interface{}{
		"aa": 1,
		"bb": []string{"1", "pp"},
		"cc": []map[string]string{map[string]string{"hh": "mm"}},
		"dd": "88",
		"tt": map[string]string{"uu": "9", "kk": "88"}}
	fmt.Println(urlencode.Encode(data))
}

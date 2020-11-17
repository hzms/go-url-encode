package main

import (
	"fmt"
	"urlencode"
)

type Book struct {
	Name string `urlencode:"book_name"`
	Id   int
}

func main() {
	data := map[string]interface{}{
		"int":    1,
		"arr":    []string{"tom", "lucy"},
		"string": "text",
		"obj":    map[string]string{"uu": "9", "kk": "88"},
		"objarr": []map[string]string{map[string]string{"name": "namevalue"}},
		"struct": Book{Name: "test", Id: 1},
	}
	fmt.Println(urlencode.Encode(data))
}

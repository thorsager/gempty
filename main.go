package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/thorsager/gempty/gempty"
)

type wrapper struct {
	v any
}

type Foo struct {
	A string `json:"a"`
	B int    `json:"b"`
	C bool   `json:"c"`
}

func main() {
	testUnmarshal(`{"a":"a","b":2,"c":true}`, &Foo{"a", 2, true})
	testUnmarshal(`{"a":"a","b":2,"c":true}`, Foo{"a", 2, true})
}

// testUnmarshal a simpel function that demonstrates how it is possible to
// test that unmarshalling end up with a specific set of values, without
// having to pass types or placeholders.
func testUnmarshal[T comparable](s string, r T) {
	var err error
	var c T
	c, err = gempty.Clone(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("(%T/%+v) --> (%T/%+v)\n", r, r, c, c)
	if !gempty.IsPtr(c) {
		err = json.Unmarshal([]byte(s), &c)
	} else {
		err = json.Unmarshal([]byte(s), c)
	}
	// err = json.Unmarshal([]byte(s), gempty.AsPtr(c))
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(r, c) {
		fmt.Printf("(%T/%+v) != (%T/%+v)\n", r, r, c, c)
	} else {
		fmt.Printf("(%T/%+v) == (%T/%+v)\n", r, r, c, c)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type Obj struct {
	Id   int
	Name string
	O    *Obj
	L    []*Obj
	M    map[string]*Obj
}

func main() {

	var o1 = &Obj{
		Id:   1,
		Name: "aaa",
		O: &Obj{
			Id:   2,
			Name: "bbb",
		},
		L: []*Obj{{
			Id:   3,
			Name: "ccc",
		}},
		M: map[string]*Obj{
			"ddd": {
				Id:   4,
				Name: "ddd",
			},
		},
	}

	var deepCopy = func(src, desc interface{}) {
		data, _ := json.Marshal(src)
		json.Unmarshal(data, desc)
		fmt.Println(desc)
	}

	fmt.Println(o1)
	var o2 = &Obj{}
	deepCopy(o1, o2)
	fmt.Println(*o2)

}

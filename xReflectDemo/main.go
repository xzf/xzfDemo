package main

import (
	"xzf/xReflect"
	"fmt"
	"encoding/json"
)

type Obj struct {
	Str1 string
	Str2 string
	Str3 string
	Str4 string
	Str5 string
}

func main() {
	var obj Obj
	sourceValue := []string{"1", "2", "3", "4", "5"}
	for index, str := range sourceValue {
		xReflect.SetField(xReflect.SetFieldReq{
			Obj:   &obj,
			Index: index,
			Value: str,
		})
	}
	b, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(b))
}

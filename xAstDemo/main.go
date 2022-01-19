package main

import (
	"xzf/xAst"
	"fmt"
)

func main() {
	groupSlice, err := xAst.GetConstList(`package main

const(
 A = "A"
 B = "B"
 C = "C"
)

const(
 D = "D"
 E = "E"
 F = "F"
)

const G = "G"
const H = 1
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range groupSlice {
		fmt.Println(item.Name, "=", item.Value)
	}
}

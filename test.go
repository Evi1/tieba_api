package main

import (
	"encoding/json"
	"fmt"
)

type RPlugin interface {
	RunPlugin(string) string
}

func main() {
	var rPlugin RPlugin = ARPlugin

	res := rPlugin.RunPlugin("1")
	fmt.Print(res)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	fmt.Print(int(m["Age"].(float64)))
}

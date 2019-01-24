package main

import (
	"fmt"
	"github.com/Evi1/gotools/json"
)

type RPlugin interface {
	RunPlugin(string) string
}

func main() {
	var rPlugin RPlugin = ARPlugin
	m := map[string]interface{}{"func": "check_cookies", "cookies": "\n" +
		""}
	jstr, _ := json.MapToJsonStr(m)
	fmt.Println(jstr)
	res := rPlugin.RunPlugin(jstr)
	fmt.Println(res)
	m = map[string]interface{}{"func": "get_liked", "cookies": "\n" +
		""}
	jstr, _ = json.MapToJsonStr(m)
	fmt.Println(jstr)
	res = rPlugin.RunPlugin(jstr)
	fmt.Println(res)
}

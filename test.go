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
	m := map[string]interface{}{"func": "check_cookies", "cookies": "BDUSS=0NQdmpwbX4wWmR2MFNBRlEzNVJZeFRkcDFXRGtQLTVzQm92Sy1zZk5teXA5c2xZSVFBQUFBJCQAAAAAAAAAAAEAAACDqdI5utrK~dfWyrHJ0LjnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKlpolipaaJYZ\n" +
		"TOKEN=f78835ccc7f13b431103e28415f47f12f398c17e2cf229f730b2b39f4aa9a2aa"}
	jstr, _ := json.MapToJsonStr(m)
	fmt.Println(jstr)
	res := rPlugin.RunPlugin(jstr)
	fmt.Println(res)
	m = map[string]interface{}{"func": "get_liked", "cookies": "BDUSS=0NQdmpwbX4wWmR2MFNBRlEzNVJZeFRkcDFXRGtQLTVzQm92Sy1zZk5teXA5c2xZSVFBQUFBJCQAAAAAAAAAAAEAAACDqdI5utrK~dfWyrHJ0LjnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKlpolipaaJYZ\n" +
		"TOKEN=f78835ccc7f13b431103e28415f47f12f398c17e2cf229f730b2b39f4aa9a2aa"}
	jstr, _ = json.MapToJsonStr(m)
	fmt.Println(jstr)
	res = rPlugin.RunPlugin(jstr)
	fmt.Println(res)
}

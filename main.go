package main

import (
	"errors"
	"github.com/Evi1/gotools/json"
)

type rPlugin int

func error2jsonStr(err error) string {
	m := map[string]interface{}{"ok": "false", "info": err.Error()}
	jstr, _ := json.MapToJsonStr(m)
	return jstr
}

var FuncMap = make(map[string]func(map[string]interface{}) string)

func (rp rPlugin) RunPlugin(jArgs string) string {
	m, e := json.JsonStrToSIMap(jArgs)
	if e != nil {
		return error2jsonStr(e)
	}
	functionName, ok := m["func"]
	if !ok {
		return error2jsonStr(errors.New("no func in json code"))
	}
	function, ok := FuncMap[functionName.(string)]
	if !ok {
		return error2jsonStr(errors.New("func not supported"))
	}
	return function(m)
}

// exported as symbol named "Greeter"
var ARPlugin rPlugin

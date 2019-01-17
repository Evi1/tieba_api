package main

import (
	"github.com/Evi1/gotools/json"
)

func init() {
	FuncMap["check_cookies"] = checkCookies
}

func checkCookies(m map[string]interface{}) string {
	cookieJar, errStr := makeBaiduCJFromeMap(m)
	if len(errStr) != 0 {
		return errStr
	}
	if GetCookieState(cookieJar) {
		m := map[string]interface{}{"ok": "true"}
		jstr, _ := json.MapToJsonStr(m)
		return jstr
	} else {
		m := map[string]interface{}{"ok": "false", "info": "cookie login failed"}
		jstr, _ := json.MapToJsonStr(m)
		return jstr
	}
}

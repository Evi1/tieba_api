package main

import (
	"fmt"
	"github.com/Evi1/gotools/json"
	"os"
	"regexp"
)

type LikedTieba struct {
	TiebaId     int
	Name        string
	UnicodeName string
	Exp         int
}

func init() {
	FuncMap["get_liked"] = getLiked
}

func getLiked(m map[string]interface{}) string {
	cookieJar, errStr := makeBaiduCJFromeMap(m)
	fmt.Println(cookieJar)
	if len(errStr) != 0 {
		return errStr
	}
	pn := 0
	likedTiebaList := make([]LikedTieba, 0)
	for {
		pn++
		url := "http://tieba.baidu.com/f/like/mylike?pn=" + fmt.Sprintf("%d", pn)
		body, fetchErr := Fetch(url, nil, cookieJar)
		if fetchErr != nil {
			return error2jsonStr(fetchErr)
		}
		fmt.Println(body)
		os.Exit(0)
		reg := regexp.MustCompile("<tr><td>.+?</tr>")
		allTr := reg.FindAllString(body, -1)
		for _, line := range allTr {
			likedTieba, err := ParseLikedTieba(line)
			if err != nil {
				continue
			}
			print(likedTieba.Name)
			likedTiebaList = append(likedTiebaList, likedTieba)
		}
		if allTr == nil {
			break
		}
	}
	str, e := json.MapToJsonStr(map[string]interface{}{"list": likedTiebaList})
	if e != nil {
		return error2jsonStr(e)
	}
	return str
}

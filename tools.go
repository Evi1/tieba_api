package main

import (
	"bytes"
	"errors"
	"github.com/Evi1/gotools/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func makeBaiduCJFromeMap(m map[string]interface{}) (*cookiejar.Jar, string) {
	rawCookie, ok := m["cookies"]
	if !ok {
		return nil, error2jsonStr(errors.New("no cookies in json code"))
	}
	cookies := makeBaiduCookiesFromStr(rawCookie.(string))
	cookieJar := makeBaiduCJ(cookies)
	return cookieJar, ""
}
func ParseLikedTieba(html string) (LikedTieba, error) {
	likedTieba := LikedTieba{}
	exp := regexp.MustCompile("<a href=\"/f\\?kw=(.*?)\" title=\"(.*?)\"")
	names := exp.FindStringSubmatch(html)
	if names == nil {
		return likedTieba, errors.New("Cannot get parse string")
	}
	likedTieba.UnicodeName = names[1]
	likedTieba.Name = names[2]
	exp = regexp.MustCompile("<a class=\"cur_exp\".+?>(\\d+)</a>")
	likedTieba.Exp, _ = strconv.Atoi(exp.FindStringSubmatch(html)[1])
	exp = regexp.MustCompile("balvid=\"(\\d+)\"")
	likedTieba.TiebaId, _ = strconv.Atoi(exp.FindStringSubmatch(html)[1])
	return likedTieba, nil
}
func makeBaiduCookiesFromStr(cookieStr string) (cookies []*http.Cookie) {
	cookies = make([]*http.Cookie, 0)
	rawCookieList := strings.Split(strings.Replace(cookieStr, "\r\n", "\n", -1), "\n")
	for _, rawCookieLine := range rawCookieList {
		rawCookieInfo := strings.SplitN(rawCookieLine, "=", 2)
		if len(rawCookieInfo) < 2 {
			continue
		}
		cookies = append(cookies, &http.Cookie{
			Name:   rawCookieInfo[0],
			Value:  rawCookieInfo[1],
			Domain: ".baidu.com",
		})
	}
	return
}

func makeBaiduCJ(cookies []*http.Cookie) *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	URL, _ := url.Parse("http://baidu.com")
	cookieJar.SetCookies(URL, cookies)
	return cookieJar
}

func fetch_tbs(ptrCookieJar *cookiejar.Jar) string {
	body, err := Fetch("http://tieba.baidu.com/dc/common/tbs", nil, ptrCookieJar)
	if err != nil {
		return ""
	}
	return body
}

func GetCookieState(ptrCookieJar *cookiejar.Jar) bool {
	str := fetch_tbs(ptrCookieJar)
	m, e := json.JsonStrToSIMap(str)
	if e != nil {
		return false
	}
	lstate, ok := m["is_login"]
	if !ok || int(lstate.(float64)) != 1 {
		return false
	}
	return true
}

func Fetch(targetUrl string, postData map[string]string, ptrCookieJar *cookiejar.Jar) (string, error) {
	var request *http.Request
	httpClient := &http.Client{
		Jar: ptrCookieJar,
	}
	if nil == postData {
		request, _ = http.NewRequest("GET", targetUrl, nil)
	} else {
		postParams := url.Values{}
		for key, value := range postData {
			postParams.Set(key, value)
		}
		postDataStr := postParams.Encode()
		postDataBytes := []byte(postDataStr)
		postBytesReader := bytes.NewReader(postDataBytes)
		request, _ = http.NewRequest("POST", targetUrl, postBytesReader)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	//fmt.Println("fetching")
	response, fetchError := httpClient.Do(request)
	//fmt.Println("fetching done")
	if fetchError != nil {
		return "", fetchError
	}
	defer response.Body.Close()
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return "", readError
	}
	return string(body), nil
}

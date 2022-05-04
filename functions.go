if strings.ToUpper(Method) == "POST" {
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
req.Header.Set("User-Agent", UserAgent)
req.Header.Set("Cookie", "rememberMe="+RememberMe) // 设置请求头
return !strings.Contains(SetCookieAll, "rememberMe=deleteMe;"), nil // 检测是否包含"deleteMe"
```

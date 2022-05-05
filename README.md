# Shiro_killer 

#### 正文

#### 1. 用法
在项目文件夹使用 `go build` 编译
`ShiroKeyCheck.exe -f urls.txt` 批量扫描 `urls.txt` 中的目标

*可选参数*
```
*-ua User-Agent `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36`*
*-m 发送请求的方式 `GET/POST`*
*-content 以 `POST` 方式发送的内容(-m POST时有效)*
*-timeout 每个请求的超时时间 `3`*
*-interval 请求之间间隔的时间 `0`*
*-proxy HTTP代理如 `http://127.0.0.1:8080`*
*-key 指定需要检测的KEY文件*
*-t 并发数量 `50`*
*-k 标签指定keys文件*
*-rm rememberMe 关键字的别名 `rememberMe`*
```

#### 2. 优点

- 单个目标爆破时间短，多目标并发检测平均速度更快
- 检测准确率高
- 内置大量已公开KEY且可自行拓展

#### 3. 关键代码分析

```
"main.go"

func KeyCheck(TargetUrl string) (bool, string) {
    Content, _ := base64.StdEncoding.DecodeString(CheckContent)
    isFind, Result := false, ""
    if SKey != "" {
    time.Sleep(time.Duration(Interval) * time.Second)
    isFind, Result = FindTheKey(SKey, Content, TargetUrl)
} else {
    isFind = false
    for i := range ShiroKeys 
{ // 遍历Key列表
    time.Sleep(time.Duration(Interval) * time.Second)
    isFind, Result = FindTheKey(ShiroKeys[i], Content, TargetUrl)
if isFind {
    break // 找到任意Key既返回结果
}
}
}
    return isFind, Result
}
    f, err := os.Open(UrlFile)
if err != nil {
    panic(err)
}
    defer f.Close()
    rd := bufio.NewReader(f)
    startTime := time.Now()
for {
    UnFormatted, _, err := rd.ReadLine() // 逐行读取目标
    if err == io.EOF {
    break
}
    TargetUrl := string(UnFormatted)
    if !strings.Contains(TargetUrl, "http://") && !strings.Contains(TargetUrl, "https://") {
    TargetUrl = "https://" + TargetUrl
}
    wg.Add(1)
    pool.Submit(func() { // 提交并发爆破任务
    StartTask(string(TargetUrl))
    wg.Done()
})
}
    wg.Wait()

"functions.go"

if strings.ToUpper(Method) == "POST" {
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
    req.Header.Set("User-Agent", UserAgent)
    req.Header.Set("Cookie", "rememberMe="+RememberMe) // 设置请求头
    return !strings.Contains(SetCookieAll, "rememberMe=deleteMe;"), nil // 检测是否包含"deleteMe"
```

![图片](https://github.com/Peony2022/shiro_killer/blob/main/%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE1.png)
![图片](https://github.com/Peony2022/shiro_killer/blob/main/%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE2.png)
![图片](https://github.com/Peony2022/shiro_killer/blob/main/%E6%9C%AC%E8%8D%89%E7%BA%B2%E7%9B%AE.jpg)


注：后面可能会想成立一个知识星球，持续更新与分享更多自开发各种工具，以及各种红队姿势。
by:铁皮石斛，白术。

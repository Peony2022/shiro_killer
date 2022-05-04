# shiro_killer

1. 用法

`ShiroKeyCheck.exe -f urls.txt` 批量扫描 `urls.txt` 中的目标

*可选参数*

> - *-ua User-Agent `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36`*
> - *-m 发送请求的方式 `GET/POST`*
> - *-content 以 `POST` 方式发送的内容(-m POST时有效)*
> - *-timeout 每个请求的超时时间 `3`*
> - *-interval 请求之间间隔的时间 `0`*
> - *-proxy HTTP代理如 `http://127.0.0.1:8080`*
> - *-key 指定需要检测的KEY*
> - *-t 并发数量 `50`*
> - *-k 标签指定keys文件*



2. 优点

- 单个目标爆破时间短，多目标并发检测平均速度更快
- 检测准确率高
- 内置大量已公开KEY且可自行拓展



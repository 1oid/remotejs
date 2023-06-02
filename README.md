# remotejs
remote execute js when debugger.paused

## example
```shell
前置条件, 需要安装chrome

# 查看帮助
./remotejs -h
GLOBAL OPTIONS:
   --url value, -u value            open url when open chrome, default blank url
   --chrome-path value, --cp value  use specified chrome path
   --proxy value                    set proxy for browser
   --remote-debug-address value     use remote chrome debugging
   --web-listen value               web server port (default: "8088")
   --help, -h                       show help
   

./remotejs                                              # 打开一个空白的浏览器
./remotejs -u [URL]                                     # 打开一个浏览器，并加载指定url
./remotejs --remote-debug-address "ws://127.0.0.1:9222" # 指定一个远程浏览器(需要目标开remote-debugger-port)

# 其他看 --help
```

## Todos (下次一定的事情)
- [ ] 多tab的debugPauseEvent捕获
- [ ] 配置文件
- [ ] 其他

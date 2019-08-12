# go 编写 shellcode 加载器

引用
> https://github.com/brimstone/go-shellcode

## 使用
git clone https://github.com/jax777/shellcode-launch
copy 至 $GOPATH/src 路径下
### windows 下
生成shellcode
修改 `winlaunch.go` 文件sc 变量
在当前目录打开cmd
- 32 位 运行 
```
set CGO_ENABLED=0
set GOOS=windows 
set GOARCH=386 
go build -ldflags="-s -w" winlaunch.go
```
- 64 位 运行 win_64.bat
```
set CGO_ENABLED=0
set GOOS=windows 
set GOARCH=amd64 
go build -ldflags="-s -w" winlaunch.go
```

### linux 下
生成shellcode
修改 `linuxlaunch.go` 文件sc 变量
- 32 位 运行
`CGO_ENABLED=1 GOOS=linux GOARCH=386  go build -ldflags="-s -w" linuxlaunch.go`
- 64 位 运行
`CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" linuxlaunch.go`
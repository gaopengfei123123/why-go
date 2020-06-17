首先编译文件
```sh
$ GOOS=js GOARCH=wasm go build -o main.wasm
```
然后导入wasm.js文件
```
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.{html,js}" .
```

服务器上运行index.html文件看一下效果

[文档地址](https://github.com/golang/go/wiki/WebAssembly)


由于正常 go build 打包出来的文件最少都2mb, 对网页加载来说不太友好, 需要gzip压缩, 或者现有的工具 tinygo 什么的

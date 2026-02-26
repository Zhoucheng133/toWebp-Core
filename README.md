# toWebp Core

![License](https://img.shields.io/badge/License-MIT-dark_green)

## 简介

这是[toWebp](https://github.com/Zhoucheng133/toWebp)核心模块

## 构建

如果你要自行构建该动态库，可以使用下面的命令进行构建：

```bash
#  macOS
go build -buildmode=c-shared -ldflags="-s -w" -o build/core.dylib
# Windows
go build -buildmode=c-shared -ldflags="-s -w" -o build/core.dll
```
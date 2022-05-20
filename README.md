# 简介

作者：京城郭少

# 关于项目

* 兼容golang原生error库
* gerror会自动生成traceID，方便排查整条链路的错误。
* gerror会自动获取错误发生的文件名+函数名+行号，方便问题的定位。

# Demo

* 详情可看example目录

```go
package main

import (
	"fmt"
	"github.com/GuoFlight/gerror"
)

func main() {
	err := gerror.NewErr("错误")
	fmt.Println(err.TraceID)
	fmt.Println(err.ErrFile)
	fmt.Println(err.ErrFunc)
	fmt.Println(err.ErrLine)
	fmt.Println(err.Error())
}
```

/*
 * @file: basic2.go
 * @brief: Go基础语法-常量定义
 * @author: feihu1996.cn
 * @date: 18-7-23
 * @version: 1.0
 */

package main

import (
    "math"
    "fmt"
)

const filename = "basic2.go"  // 未声明类型时，常量的作用相当于简单的文本替换， 其值可作为各种类型使用

func consts() {
    const (
        a = 3
        b = 4
    )
    const d string = "const"
    var c int
    c = int(math.Sqrt(a*a + b*b))
    fmt.Println(filename, c)
    fmt.Println(d)
}

func enums() {
    /*
    使用常量实现枚举类型
    */
    const (
      x = 0
      y = 1
      z = 2
    )
    const (
       cpp = iota  // 从0开始自动加1
       _
       java
       python
       golang
       javascript
    )
    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(x, y, z)
    fmt.Println(cpp, java, python, golang, javascript)
    fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
    consts()
    enums()
}

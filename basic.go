/*
 * @file: basic.go
 * @brief:Go基础语法-变量定义
 * @author: feihu1996.cn
 * @date: 18-7-23
 * @version: 1.0
 */

package main

import "fmt"

var (
    aa = 3
    ss = "kkk"
    bb = true
)

func variableZeroValue() {
    /*
    变量声明
    变量初始值
    */
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
    /*
    变量定义：声明并赋值
    */
    var a, b int = 3, 4
    var s string = "abc"
    fmt.Println(a, b, s)
}

func variableTypeDeduction() {
    /*
    变量定义：类型推断
    */
    var a, b, c, d = 3, 4, true, "def"
    fmt.Println(a, b, c, d)
}

func variableShorter() {
    /*
    变量定义：快速声明(仅限函数内部使用)
    */
    a, b, c, d := 3, 4, true, "def"
    fmt.Println(a, b, c, d)
}

func main() {
    fmt.Println("Hello, World!")
    variableZeroValue()
    variableInitialValue()
    variableTypeDeduction()
    variableShorter()
    fmt.Println(aa, ss, bb)
}

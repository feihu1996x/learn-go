/*
 * @file: strings.go
 * @brief: Go内建容器——字符和字符串处理
 * @author: feihu1996.cn
 * @date: 18-8-6
 * @version: 1.0
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func main(){
    s := "诗意代码"

    // 打印字符串s的长度（单位是字节）
    // 一个字节==1个内存单元==8位二进制==2位十六进制
    // Go语言采用UTF8编码
    // 一个中文字符相当于3个字节
    // 一个英文字符相当于1个字节
    //=>12
    fmt.Println(len(s))

    // 将字符串s转换为字节切片，
    // 再以字符串的格式输出
    //=>诗意代码
    fmt.Printf("%s\n", []byte(s))

    // 将字符串s转换为字节切片，
    // 再以字节的格式输出（16进制表示）
    //=>E8AF97E6848FE4BBA3E7A081(24位16进制<==>96位二进制<==>12个字节)
    fmt.Printf("%X\n", []byte(s))

    // 将字符串s转换为字节切片，
    // 遍历字节切片，
    // 并输出每一个字节(16进制)
    for _,b := range []byte(s) {
        fmt.Printf("%X\n", b)
    }

    // 输出字符串s的长度（单位字符）
    fmt.Println(utf8.RuneCountInString(s))

    bytes := []byte(s)
    for len(bytes) > 0 {
        //fmt.Println(bytes)
        // utf8.DecodeRune(bytes)会从字节切片中\
        // 取出字符及其size(单位字节)
        // 只会取一次
        ch, size := utf8.DecodeRune(bytes)
        // bytes[size:]将已经被取出的字符删除\
        // （从其size代表的索引处切片）
        // 方便utf8.DecodeRune(bytes)重新取值
        bytes = bytes[size:]
        // 将从字节切片中取出的字符打印出来
        fmt.Printf("%c\n", ch)
    }

    // 遍历字符串
    // 并输出每一个字符(rune)（16进制表示）
    // i是每一个字符起始字节的索引
    for i,ch := range s {
        fmt.Printf("(%d, %X)\n", i, ch)
    }

    // 将字符串s转换为字符切片,遍历字符切片，
    // 输出每一个字符
    // i是每一个字符的索引
    for i, ch := range []rune(s) {
        fmt.Printf("(%d, %c)\n", i, ch)
    }

}
/*
 * @file: basic1.go
 * @brief: Go基础语法-内建变量类型
 * @author: feihu1996.cn
 * @date: 18-7-23
 * @version: 1.0
 */

 /*
     Go语言内建变量类型：
     1. 布尔类型：
         bool
     2. 字符串：
         string
     3. 整型(小括号中的u表示如果前面是u，则它是一个无符号整型)：
        (u)int: 长度或位数取决于操作系统
         (u)int8: 8位
         (u)int16: 16位
        (u)int32: 32位
         (u)int64: 64位
         uintptr: 指针类型，长度或位数取决于操作系统
     4. 字节型：
         byte: 8位长度，等价于uint8
     5. 字符型：
         rune: 32位长度，等价于uint32
     6. 浮点型：
         float32: 32位浮点型
         float64: 64位浮点型
     7. 复数类型：
        complex64: 实部和虚部都是float32
        complex128: 实部和虚部都是float64
 */

package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler() {
 	//c := 3 + 4i
 	//fmt.Println(cmplx.Abs(c)) // 复数取模

 	/*
 	欧拉公式
 	*/
 	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func main() {
	euler()
}

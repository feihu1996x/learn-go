/*
 * @file: pointer.go
 * @brief:  Go基础语法-指针
 * @author: feihu1996.cn
 * @date: 18-7-25
 * @version: 1.0
*/

/*
    Go语言指针不能运算
    Go语言中的参数传递都是值传递
*/

package main

import "fmt"

func swap(a, b *int){
    *b, *a = *a, *b
}

func main() {
    a, b := 3, 4
    swap(&a, &b)
    fmt.Println(a, b)
}

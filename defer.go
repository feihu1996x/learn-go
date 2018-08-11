/*
 * @file: defer.go
 * @brief: Go异常处理之defer调用
 * @author: feihu1996.cn
 * @date: 18-8-11
 * @version: 1.0
*/

package main

import (
    "fmt"
    "os"
    "bufio"
    "errors"
)

// 定义一个名为intGen的函数类型
type intGen2 func() int

func fibonacci2() intGen2 {
    /*
    闭包
    返回一个斐波纳契数列生成器函数
    */
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func tryDefer() {
    // defer确保了fmt.Println(1)\
    // 在tryDefer函数结束时调用\
    // 即使中间出现return或者异常被抛出
    defer fmt.Println(1)

    defer fmt.Println(2)

    i := 100
    defer fmt.Println(i)  // 参数在defer语句时计算
    i++

    fmt.Println(3)
}

func writeFile(filename string) {
    // file, err := os.Create(filename)
    file, err := os.OpenFile(
        filename, os.O_EXCL|os.O_CREATE, 0666)
    if err != nil {
        // 异常处理

        // panic(err)

        // fmt.Println("file already exists.")

        // fmt.Println("Error:", err.Error())

        // fmt.Println("Error:", err)

        //if pathError, ok := err.(*os.PathError); !ok {
        //    panic(err)
        //} else {
        //    fmt.Println(pathError.Op,  // open
        //        pathError.Path,  // abc.txt
        //            pathError.Err)  // file exits
        //}

        // err只是一个普通的值，这个值也可以自定义
        err := errors.New("something wrong")
        panic(err)

        return
    }

    // 确保writeFile函数结束时关闭文件
    defer file.Close()

    writer := bufio.NewWriter(file)

    // 确保writeFile函数结束时\
    // 将内存中缓存的数据写入文件
    defer writer.Flush()

    f := fibonacci2()

    // 将前20个斐波纳契数写入文件
    for i:=0;i<20;i++ {
        fmt.Fprintln(writer, f())
    }
}

func main() {
    tryDefer()  // 3 100 2 1

    writeFile("abc.txt")
}

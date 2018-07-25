/*
 * @file: loop.go
 * @brief: Go基础语法-循环
 * @author: feihu1996.cn
 * @date: 18-7-25
 * @version: 1.0
*/

package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
)

func sum() int {
    res := 0
    for i:=1; i<=100; i++ { // for的条件里不需要括号，初始条件、结束条件、递增表达式都可以省略
       res += i
    }
    return res
}

func convertToBin(n int) string {
    /*
        十进制转二进制
    */
    result := ""
    for ; n>0; n/=2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    return result
}

func printFile(filename string) {
    /*
        按行读取文件
    */
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {  // 只有结束条件的for语句，等价于其他语言中的while循环
        fmt.Println(scanner.Text())
    }
}

func forever() {
    for {
        fmt.Println("hello")  // 结束条件都没有？那就是个死循环
    }
}

func main() {
    fmt.Println(sum())
    fmt.Println(
        convertToBin(5), // 101
        convertToBin(13), // 1101
        convertToBin(7234566),
        convertToBin(0),
    )
    printFile("abc.txt")
    forever()
}

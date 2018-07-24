/*
 * @file: branch.go
 * @brief: Go基础语法-条件语句
 * @author: feihu1996.cn
 * @date: 18-7-24
 * @version: 1.0
*/

package main

import (
    "io/ioutil"
    "fmt"
)

const filename1 = "abc.txt"

func testIf() {
    contents, err := ioutil.ReadFile(filename1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

func testIf1() {
    if contents, err := ioutil.ReadFile(filename1); err != nil {  // if的条件里可以定义变量，但变量的作用域仅限此块级作用域
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

func testSwitch(a, b int, op string) int {
    var result int
    switch op {
    case "+":
        result = a + b
        // switch会自动break, 除非使用fallthrough
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        result = a / b
    default:
        panic("unsupported operator:" + op)
    }
    return result
}

func grade(score int) string {
    g := ""
    switch { // switch后可以没有表达式
    case score < 0 || score > 100:
        panic(fmt.Sprintf(
            "Wrong Score: %d", score))
    case score < 60:
        g = "F"
    case score < 80:
        g = "C"
    case score < 90:
        g = "B"
    case score <= 100:
        g = "A"
    }
    return g
}

func main() {
    testIf()
    testIf1()
    fmt.Println(testSwitch(1, 2, "+"))
    fmt.Println(
        grade(10),
        grade(70),
        grade(85),
        grade(100),
        //grade(200),
    )
}

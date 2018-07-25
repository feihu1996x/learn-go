/*
 * @file: func.go
 * @brief: Go基础语法-函数
 * @author: feihu1996.cn
 * @date: 18-7-25
 * @version: 1.0
*/

package main

import (
    "fmt"
    "reflect"
    "runtime"
    "math"
)

func eval(a,b int, op string) (int, error) {  // 多个返回值的函数, 其中一个返回错误
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        q, _ := div(a, b)
        return q, nil
    default:
        return 0, fmt.Errorf("unsupported operation: %s", op)
    }
}

func div(a, b int) (int, int) {  // 多个返回值的函数
    return a / b, a % b
}

func div1(a, b int) (q, r int) {  // 多个返回值的函数，为返回值命名
    //q = a / b
    //r = a % b
    //return
    return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {  // 函数作为参数
    p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()  // 拿到函数名
    fmt.Printf("Calling function %s with args " + "(%d, %d)\n", opName, a, b)
    return op(a, b)
}

func pow(a, b int) int {
    return int(math.Pow(float64(a), float64(b)))
}

func total(numbers ...int) int {  // 可变参数列表
    s := 0
    for i := range numbers {
        s += numbers[i]
    }
    return s
}


func main() {
    if result, err := eval(12, 13, "x"); err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println(result)
    }

    fmt.Println(div(13, 3))

    a, b := div1(13, 3)
    fmt.Println(a)
    fmt.Println(b)

    fmt.Println(apply(pow, 3, 4))

    fmt.Println(apply(func(a, b int) int {
        return int(math.Pow(float64(a), float64(b)))
    }, 3, 4))

    fmt.Println(total(1, 2, 3, 4, 5))
}

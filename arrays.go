/*
 * @file: arrays.go
 * @brief: Go内建容器——数组
 * @author: feihu1996.cn
 * @date: 18-8-5
 * @version: 1.0
*/

package main

import "fmt"

func printArray(arr [5]int){
    // 函数接收一个长度为5的整型数组作为参数
    // Go语言中，
    // 数组是值类型，
    // 传参时会对数组进行一次深拷贝
    arr[0] = 10000
    for i,v := range arr{
        fmt.Println(i, v)
    }
}

func printArray2(arr *[5]int){
    arr[0] = 10000
    for i,v := range arr{
        fmt.Println(i, v)
    }
}

func main(){
    // 声明一个长度为5的整型数组
    // 其中每个元素的初始值均为0
    var arr1 [5]int

    // 定义一个长度为3的整型数组
    arr2 := [3]int{1, 3, 5}

    // 定义一个不定长数组
    arr3 := [...]int{2, 4, 6, 8, 10}

    // 声明一个二维数组
    var grid [4][5]int

    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)
    fmt.Println(grid)

    // 遍历数组——方式一
    for i := 0; i < len(arr3); i++ {
        fmt.Println(arr3[i])
    }

    // 遍历数组——方式二
    for i := range arr3 {
        fmt.Println(arr3[i])
    }

    // 遍历数组——方式三
    for _,v := range arr3 {
        fmt.Println(v)
    }

    printArray(arr3)
    fmt.Println(arr3)

    printArray2(&arr3)
    fmt.Println(arr3)

}

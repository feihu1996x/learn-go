/*
 * @file: slices.go
 * @brief: Go内建容器——切片
 * @author: feihu1996.cn
 * @date: 18-8-5
 * @version: 1.0
*/

package main

import (
    "fmt"
)

func updateSlice(s []int){
    // 函数接收一个整型切片作为参数
    // 数组切片被不是值类型，
    // 本身没有数据，
    // 只是对原数组的一种视图(view)
    // “视图”可以理解为从不同的视角观察到的数组
    s[0] = 10000000
}

func printSlice(s []int){
    /*
        给定一个切片，
        打印其长度和容量
    */
    fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main(){
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

    fmt.Println("arr[2:6]", arr[2:6])
    fmt.Println("arr[:6]", arr[:6])
    fmt.Println("arr[2:]", arr[2:])
    fmt.Println("arr[:]", arr[:])

    s := arr[3:]
    updateSlice(s)
    fmt.Println(s)
    fmt.Println(arr)

    // 切片本身还可以进一步切片，
    // 这个过程被称之为reslice
    fmt.Println("arr[2:][2:]", arr[2:][2:])

    // 切片“越界”后，
    // 会自动基于对原数组的视图进行扩展
    // 但是这种扩展只能向后扩展
    fmt.Println("arr[2:][3:5]", arr[2:][3:5])

    // 查看切片的长度
    fmt.Println(len(arr[2:5]))

    // 查看切片的容量（包含了切片扩展）
    fmt.Println(cap(arr[2:5]))

    /*
        向切片添加元素
    */
    arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s1 := arr1[2:5]
    fmt.Println("original: cap(s1)=", cap(s1)) // 6

    s1 = append(s1, 8)
    fmt.Println(s1)  // [2 3 4 8]
    fmt.Println(arr1)  // [0 1 2 3 4 8 6 7]

    s1 = append(s1, 9)
    fmt.Println(s1)  // [2 3 4 8 9]
    fmt.Println(arr1)  // [0 1 2 3 4 8 9 7]

    s1 = append(s1, 10)
    fmt.Println(s1)  // [2 3 4 8 9 10]
    fmt.Println(arr1)  // [0 1 2 3 4 8 9 10]

    s1 = append(s1, 11)
    fmt.Println(s1)  // [2 3 4 8 9 10 11]
    fmt.Println("now: cap(s1)=", cap(s1))  // 12
    // s1的容量是6
    // 向切片中添加元素时，
    // 如果超过了其容量
    // 切片不再view原数组，
    // 系统会为其分配更大的底层数组
    fmt.Println(arr1)  // [0 1 2 3 4 8 9 10]

    // 声明一个整型切片，
    // 其值为nil
    var anotherSlice []int
    for i := 0; i < 100; i++{
        // 观察切片长度和容量的变化
        printSlice(anotherSlice)
        anotherSlice = append(anotherSlice, 2 * i + 1)
    }
    fmt.Println(anotherSlice)

    // 创建一个包含初始值的整型切片
    // 切片容量默认等于其长度
    // 元素的默认值均为0
    slice2 := []int{2, 4, 6, 8}
    printSlice(slice2)

    // 创建一个长度为16的整型切片
    // 切片容量默认等于其长度
    slice3 := make([]int, 16)
    printSlice(slice3)

    // 创建一个长度为10，容量为20的整型切片
    slice4 := make([]int, 10, 20)
    printSlice(slice4)

    /*
        切片拷贝
    */
    arr4 := [3]int{1, 3, 5}
    slice5 := make([]int, 10, 20)
    slice6 := arr4[:]
    // 将slice6拷贝一份到slice5中
    copy(slice5, slice6)
    printSlice(slice5)  // [1 3 5 0 0 0 0 0 0 0]


    /*
        从切片中移除元素
    */
    // 删除中间某个元素
    //printSlice(slice5[:2])
    //printSlice(slice5[3:])
    slice5 = append(slice5[:2], slice5[3:]...)
    printSlice(slice5)

    // 删除头元素
    slice5 = slice5[1:]
    printSlice(slice5)

    // 删除尾元素
    slice5 = slice5[:len(slice5) - 1]
    printSlice(slice5)
}

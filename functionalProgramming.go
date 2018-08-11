/*
 * @file: functionalProgramming.go
 * @brief: Go函数式编程
 * @author: feihu1996.cn
 * @date: 18-8-10
 * @version: 1.0
*/

package main

import (
    "fmt"
    "strings"
    "io"
    "bufio"
)

/*
一、函数式编程概述：
1. 函数是一等公民，参数、变量、返回值都可以是函数
2. 高阶函数
3. 函数->闭包
二、”正统“函数式编程
1. 不可变性：不能有状态，只有常量和函数
2. 函数只能有一个参数
*/

// 定义一个名为iAdder的函数类型
type iAdder func(value int) (int, iAdder)

func adder() func (value int) int {
    /*
        闭包
        返回一个类加器函数
    */
    sum := 0
    return func (value int) int {
        sum += value
        return sum
    }
}

func adder2(base int) iAdder {
    /*
        闭包
        更加正统的函数式编程
        实现累加器
    */
    return func(value int) (int, iAdder) {
        return base + value, adder2(base + value)
    }
}

// 定义一个名为intGen的函数类型
type intGen func() int

// 为intGen函数类型定义一个Read方法
// 因此，它实现了io.Reader接口
// 从而使得intGen函数类型可以像文件一样被读取
func (g intGen) Read(p []byte) (n int, err error) {
    next := g()
    if next > 10000 {  // io.EOF表示已经到达文件末尾了
        return 0, io.EOF
    }
    s := fmt.Sprintf("%d\n", next)

    // TODO：incorrect if p is too small!
    return strings.NewReader(s).Read(p)
}

func fibonacci() intGen {
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

func printFileContents2(reader io.Reader) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

// 定义一个名为treeNodeF的结构体类型
type treeNodeF struct {
    value int
    left, right *treeNodeF
}

// 为treeNodeF结构体定义方法
// 获取结构成员value的值
// "node treeNodeF"被称之为接收者参数
func (node treeNodeF) getValue() {
    // 输出treeNodeF结构体类型变量的value值
    fmt.Println(node.value)
}

// 为treeNodeF结构体定义方法
// 设置结构成员value的值
func (node *treeNodeF) setValue(value int) {
    if node == nil {
        fmt.Println("Setting value to nil node, Ignored.")
        return
    }
    // 接收者参数传递也是值传递
    // 只有使用指针才可以改变结构内容
    node.value = value
}

// 为treeNodeF结构体定义方法
// 遍历当前结构树
func (node *treeNodeF) traverse(){
    node.traverseFunc(func(n *treeNodeF){
        n.getValue()
    })
}

// 为treeNodeF结构体定义方法
// 遍历当前结构树(使用函数)
func (node *treeNodeF) traverseFunc(f func(*treeNodeF)){
    /*
    传入不同的函数参数，
    在遍历树的过程中，
    可以对节点进行不同的操作
    */
    if node == nil {
        return
    }
    node.left.traverseFunc(f)
    f(node)
    node.right.traverseFunc(f)
}

func main() {
    adder := adder()
    for i:=0;i<10;i++{
       fmt.Println(adder(i))
    }

    adder2 := adder2(0)
    for i:=0;i<10;i++{
        var s int
        s, adder2 = adder2(i)
        fmt.Println(s)
    }

    // 生成斐波纳契数列生成器
    f := fibonacci()

    //fmt.Println(f())
    //fmt.Println(f())
    //fmt.Println(f())
    //fmt.Println(f())
    //fmt.Println(f())

    printFileContents2(f)

    root := treeNodeF{3, nil, nil}
    root.left = &treeNodeF{6, nil, nil}
    root.right = &treeNodeF{5, nil, nil}
    root.right.left = new(treeNodeF)

    // 遍历root树
    root.traverse()

    nodeCount := 0
    root.traverseFunc(func(*treeNodeF) {
        nodeCount++
    })
    fmt.Println(nodeCount)  // 输出root树节点数量

}

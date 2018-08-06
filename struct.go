/*
 * @file: struct.go
 * @brief: Go面向对象-结构体
 * @author: feihu1996.cn
 * @date: 18-8-6
 * @version: 1.0
*/

/*
Go语言面向对象仅支持封装，不支持继承和多态
Go语言没有class，只有struct(结构体)
*/

package main

import "fmt"

// 定义一个名为treeNode的结构体类型
type treeNode struct {
    value int
    left, right *treeNode
}

// 为treeNode结构体定义方法
// 获取结构成员value的值
// "node treeNode"被称之为接收者参数
func (node treeNode) getValue() {
    // 输出treeNode结构体类型变量的value值
    fmt.Println(node.value)
}

// 为treeNode结构体定义方法
// 设置结构成员value的值
func (node *treeNode) setValue(value int) {
    if node == nil {
        fmt.Println("Setting value to nil node, Ignored.")
        return
    }
    // 接收者参数传递也是值传递
    // 只有使用指针才可以改变结构内容
    node.value = value
}

// 为treeNode结构体定义方法
// 遍历当前结构树
func (node *treeNode) traverse(){
    if node == nil {
        return
    }
    node.left.traverse()
    node.getValue()
    node.right.traverse()
}

// 创建结构体指针的工厂函数
func createNode(value int) *treeNode {
    // 返回了一个局部变量的地址！
    // Go语言编译器会自动进行垃圾回收
    return &treeNode{value: value}
}

func main() {
    // 声明一个treeNode类型的变量root
    var root treeNode

    // 初始化root
    root = treeNode{value:3}

    // 初始化root的成员left
    root.left = &treeNode{}

    // 初始化root的成员right
    root.right = &treeNode{5, nil, nil}

    // 初始化root的成员right指向的treeNode类型的值
    // new函数用于获取参数的内存地址
    // 不论是指针还是结构体本身，一律使用.来访问成员
    root.right.left = new(treeNode)

    // 初始化root的成员left指向的treeNode类型的值
    root.left.right = createNode(2)

    fmt.Println("root:", root)

    // 定义一个结构体切片
    nodes := []treeNode {
       {value:3},
       {},
       {6, nil, &root},
    }

    fmt.Println("nodes", nodes)

    // 调用root的getValue方法
    root.getValue()

    // 调用root.left的setValue方法
    root.left.setValue(10)

    // 调用root.left的getValue方法
    root.left.getValue()

    // nil指针也可以调用方法
    var pRoot *treeNode
    fmt.Println(pRoot)
    pRoot.setValue(200)

    // 遍历root树
    root.traverse()
}
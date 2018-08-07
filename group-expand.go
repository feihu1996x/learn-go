/*
 * @file: group-expand.go
 * @brief: Go面向对象-扩展已有的类型(使用组合)
 * @author: feihu1996.cn
 * @date: 18-8-7
 * @version: 1.0
*/

package main

import "fmt"

// 定义一个名为anotherTreeNode的结构体类型
type anotherTreeNode struct {
    value int
    left, right *anotherTreeNode
}

// 为anotherTreeNode结构体定义getValue方法
func (node anotherTreeNode) getValue() {
    fmt.Println(node.value)
}

// 为anotherTreeNode结构体定义setValue方法
func (node *anotherTreeNode) setValue(value int) {
    if node == nil {
        fmt.Println("Setting value to nil node, Ignored.")
        return
    }
    node.value = value
}

// 为anotherTreeNode结构体定义traverse方法
// 遍历当前树
func (node *anotherTreeNode) traverse() {
    if node == nil {
        return
    }
    node.left.traverse()
    node.getValue()
    node.right.traverse()
}

// 定义一个名为myAnotherTreeNode的结构体类型
// 其内部含有一个node指针,指向myAnotherTreeNode
// 从而实现自定义类型的扩展
type myAnotherTreeNode struct {
    node *anotherTreeNode
}

// 为myAnotherTreeNode定义一个方法
// 换种方式遍历当前树
func (myNode *myAnotherTreeNode) postOrder() {
    if myNode == nil || myNode.node == nil{
        return
    }

    left := myAnotherTreeNode{myNode.node.left}
    left.postOrder()

    right := myAnotherTreeNode{myNode.node.right}
    right.postOrder()

    myNode.node.getValue()
}

func main(){
    root := anotherTreeNode{
        value:3,
        left:&anotherTreeNode{
            value:4,
            left:&anotherTreeNode{
                value:5,
            },
            right:&anotherTreeNode{
                value:7,
            },
        },
        right:&anotherTreeNode{
            value:6,
            left:&anotherTreeNode{
                value:8,
            },
            right:&anotherTreeNode{
                value:9,
            },
        },
    }

    // 使用扩展的方法遍历root树
    myRoot := myAnotherTreeNode{&root}
    myRoot.postOrder()

}
/*
 * @file: alias-expand.go
 * @brief: Go面向对象-扩展已有的类型(使用别名)
 * @author: feihu1996.cn
 * @date: 18-8-7
 * @version: 1.0
*/

package main

import "fmt"

// 定义一个名为queue的整型切片类型
// queue相当于是系统切片类型的别名
type queue []int

// 为queue类型定义一个push方法
func (q *queue) push(v int) {
    *q = append(*q, v)
}

// 为queue类型定义一个pop方法
func (q *queue) pop() int {
    head := (*q)[0]
    *q = (*q)[1:]
    return head
}

// 为queue类型定义一个isEmpty方法
func (q *queue) isEmpty() bool {
    return len(*q) == 0
}

func main(){
    // 定义一个queue类型的变量
    // 初始成员是1
    q := queue{1}

    q.push(2)
    q.push(3)
    fmt.Println(q.pop())
    fmt.Println(q.pop())
    fmt.Println(q.isEmpty())
    fmt.Println(q.pop())
    fmt.Println(q.isEmpty())

}
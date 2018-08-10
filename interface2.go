/*
 * @file: interface2.go
 * @brief: Go面向接口--接口的值类型
 * @author: feihu1996.cn
 * @date: 18-8-10
 * @version: 1.0
*/


package main

import "fmt"

type retriever2 interface {
    get(url string) string
}

type mockRetriever2 struct {
    contents string
}

type anotherMockRetriever2 struct {
    contents string
}

// 为mockRetriever2结构体类型定义一个get方法
// 它实现了retriever2接口
// 接收者是一个值
func (r mockRetriever2) get(url string) string {
    return r.contents
}

// 为anotherMockRetriever2结构体类型定义一个get方法
// 它实现了retriever2接口
// 接收者是一个指针
func (r *anotherMockRetriever2) get(url string) string {
    return r.contents
}

func inspect(r retriever2) {
    // 查看接口变量之type switch
    fmt.Printf("type:%T\nvalue:%v\n", r, r)
    switch v := r.(type) {
    case mockRetriever2:
        fmt.Println("contents:", v.contents)
    case *anotherMockRetriever2:
        fmt.Println("contents:", v.contents)
    }
}

// Go语言中, interface{}可以表示任何类型
// 定义一个名为queue2的任意类型切片
type queue2 []interface{}

// 为queue2类型定义一个push方法
func (q *queue2) push(v interface{}) {
    *q = append(*q, v)
}

// 为queue类型定义一个pop方法
func (q *queue2) pop()interface{} {
    head := (*q)[0]
    *q = (*q)[1:]
    return head
    // 限定pop方法值返回head(interface{})时，\
    // 取出其包含的int值，如果head不包含int值， \
    // 那么就会出现运行错误
    //return head.(int)
}

func main() {
    // 声明两个接口变量r2和r3
    // 它们由两部分组成：
    //  1. 实现者的类型
    //  2. 实现者的值（或者实现者的指针）
    // 接口变量自带指针，
    // 接口变量同样采用值传递，几乎不需要使用接口的指针
    var r2, r3 retriever2

    // 值接收者既能以指针方式使用，也能以值方式使用
    r2 = mockRetriever2{"this is a fake www.feihu1996.cn"}
    //r2 = &mockRetriever2{"this is a fake www.feihu1996.cn"}
    inspect(r2)

    // anotherMockRetriever2在实现retriever2接口时使用了指针作为接收者
    // 指针接收者只能以指针方式使用
    r3 = &anotherMockRetriever2{"this is a fake www.feihu1996.cn"}
    inspect(r3)

    // 查看接口变量之type assertion
    if r4, ok := r3.(mockRetriever2); ok {
        fmt.Println(r4.contents)
    } else {
        fmt.Println("r3 is not mockRetriever2")
    }

    // 可以向切片queue2中添加任意类型的值
    q := queue2{1}
    q.push("abc")
    q.push(true)
    fmt.Println(q.pop())
    //fmt.Println(q.pop())
    //fmt.Println(q.pop())
}

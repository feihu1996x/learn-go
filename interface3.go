/*
 * @file: interface3.go
 * @brief: Go面向接口--接口的组合
 * @author: feihu1996.cn
 * @date: 18-8-10
 * @version: 1.0
*/

package main

import "fmt"

// 定义一个名为getter的接口
type getter interface {
    get(url string) string
}

// 定义一个名为poster的接口
type poster interface {
    post(url string, form map[string]string) string
}

// 定义一个名为gettterAndPoster的接口
// 它组合了getter和poster两个接口
type gettterAndPoster interface {
    getter
    poster
}

// 定义一个名为ahRetriever的结构体类型
type ahRetriever struct {
    contents string
}

// 为ahRetriever定义一个post方法
// 它实现了poster接口
func (r *ahRetriever) post(url string, form map[string]string) string {
    r.contents = form["contents"]
    return "ok"
}

// 为ahRetriever定义一个get方法
// 它实现了getter接口,
// 同时它也实现了gettterAndPoster接口
func (r *ahRetriever) get(url string) string{
    return r.contents
}

func post(poster poster) string{
    return poster.post("http://www.feihu1996.cn", map[string]string{
        "contents": "hello",
    })
}

func get(getter getter) string{
    return getter.get("http://www.feihu1996.cn")
}

func session(s gettterAndPoster) {
    fmt.Println(post(s))
    fmt.Println(get(s))
}

func main() {
    s := ahRetriever{"this is aha retriever"}
    session(&s)
}

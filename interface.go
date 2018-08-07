/*
 * @file: interface.go
 * @brief: Go面向接口--接口的定义和实现
 * @author: feihu1996.cn
 * @date: 18-8-7
 * @version: 1.0
*/

package main

import (
    "fmt"
    "time"
    "net/http"
    "net/http/httputil"
)

// 定义一个名为retriever的接口类型
type retriever interface {
    get(url string) string
}

// 定义一个名为mockRetriever的结构体类型
type mockRetriever struct {
    contents string
}

// 定义一个名为realRetriever的结构体类型
type realRetriever struct {
    userAgent string
    timeOut time.Duration
}

// 为mockRetriever结构体定义一个get方法
// mockRetriever结构体只要实现了get方法,
// 那么它就实现了retriever接口(duck typing)
func (r mockRetriever) get(url string) string {
    return r.contents
}

// 为realRetriever结构体定义一个get方法
// realRetriever结构体只要实现了get方法,
// 那么它就实现了retriever接口(duck typing)
func (r realRetriever) get(url string) string {
    if res, err := http.Get(url); err != nil {
        panic(err)
    } else {
        if result, err := httputil.DumpResponse(res, true); err != nil {
            panic(err)
        } else {
            res.Body.Close()
            return string(result)
        }

    }
}

func download (r retriever) string {
    return r.get("http://www.feihu1996.cn")
}

func main() {
    // 声明一个r变量,
    // 它应当实现了retriever接口
    var r retriever

    r = mockRetriever{"this is a fake www.feihu1996.cn"}
    fmt.Println(download(r))

    r = realRetriever{}
    fmt.Println(download(r))
}
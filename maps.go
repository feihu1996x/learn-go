/*
 * @file: maps.go
 * @brief: Go内建容器——map
 * @author: feihu1996.cn
 * @date: 18-8-5
 * @version: 1.0
*/

package main

import "fmt"

func main(){
    // 定义一个字符串map
    // map中的key不能是任意类型：
    // 1. map使用哈希表，必须可以比较相等
    // 2. 除了slice，map和function的内建类型都可以作为key
    // 3. Struct类型不包含上述字段，也可作为key
    m := map[string]string{
        "a": "A",
        "b": "B",
        "c": "C",
    }

    // 创建一个空map
    // m1 == empty map
    m1 := make(map[string]int)

    // 声明一个map
    // m2 == nil
    // Go语言中,nil值也是可以参与运算的
    var m2 map[string]int

    fmt.Println(m)
    fmt.Println(m1)
    fmt.Println(m2)

    // 遍历map
    // map是无序的
    for k,v := range m {
        fmt.Println(k, v)
    }
    for k := range m {
        fmt.Println(k)
    }
    for _,v := range m {
        fmt.Println(v)
    }

    // 访问map中的元素
    if d,ok := m["d"]; ok {
        fmt.Println(d, ok)
    } else {
        // 当key不存在时
        // 获得的是value类型的初始值
        // 这里是一个空串
        fmt.Println(len(d), ok)
    }

    // 删除map中的元素
    fmt.Println(len(m["a"]))
    delete(m, "a")
    fmt.Println(len(m["a"]))
}
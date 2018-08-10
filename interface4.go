/*
 * @file: interface4.go
 * @brief: Go面向接口——常用系统接口
 * @author: feihu1996.cn
 * @date: 18-8-10
 * @version: 1.0
*/

package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "strings"
)

/*
    常用系统接口：
        1. Stringer(fmt.Stringer)
        2. Reader(io.Reader)
        3. Writer(io.Writer)
*/

// 定义一个名为stillRetriever的结构体类型
type stillRetriever struct {
    contents string
}

// 为stillRetriever结构体类型定义一个String方法
// 它实现了Stringer接口
func (r *stillRetriever) String() string{
    return fmt.Sprintf("stillRetriever: {contents: %s}",
        r.contents)
}

func printFile2(filename string) {
    /*
        按行读取文件
    */
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    // file实现了io.Reader接口
    printFileContents(file)
}

func printFileContents(reader io.Reader) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func main(){
    var s fmt.Stringer
    s = &stillRetriever{"this is still retriever"}
    fmt.Println(s.String())

    printFile2("abc.txt")
    printFileContents(strings.NewReader(`sasasas
        sas
        aasasa`))
}

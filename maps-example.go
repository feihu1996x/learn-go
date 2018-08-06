/*
 * @file: maps.go
 * @brief: Go内建容器——map例题
 * @author: feihu1996.cn
 * @date: 18-8-6
 * @version: 1.0
*/

package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
    /*
    寻找最长不含重复字符的子串:
    对于每一个字母x,
    如果lastOccured[x]不存在，或者 < start，那么无需操作，
    如果lastOccured[x] >= start，那么更新start，
    无论如何，都要更新lastOccured[x]和maxLength
    */

    lastOccured := make(map[byte]int)
    start := 0
    maxLength := 0

    for i, ch := range []byte(s) {
        if lastI, ok := lastOccured[ch]; ok && lastI >= start {
            start = lastI + 1
        }
        if i - start + 1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccured[ch] = i
    }

    return maxLength
}

func main() {
    fmt.Println(lengthOfNonRepeatingSubStr(
        "abcabcbb"))
    fmt.Println(lengthOfNonRepeatingSubStr(
        ""))
    fmt.Println(lengthOfNonRepeatingSubStr(
        "a"))
    fmt.Println(lengthOfNonRepeatingSubStr(
        "abcdefg"))
    fmt.Println(lengthOfNonRepeatingSubStr(
        "诗意"))
}
package main

import (
    "net/http"
    _ "net/http/pprof"
)

var m = make(map[int]int)

func main() {
    go func() {
        for i := 0; ; i++ {
            if i%(9*60) == 0 && i%7 == 0 { // 添加判断条件，防止增长过快打满内存
                m[i] = i
            }
        }
    }()

    http.ListenAndServe("0.0.0.0:8080", nil)
}
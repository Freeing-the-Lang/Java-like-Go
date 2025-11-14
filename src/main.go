package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: java-like-go <source.jgo>")
        return
    }

    inputFile := os.Args[1]
    code, err := os.ReadFile(inputFile)
    if err != nil {
        panic(err)
    }

    parsed := ParseJavaLike(string(code))
    goCode := TranspileToGo(parsed)

    out := "out.go"
    os.WriteFile(out, []byte(goCode), 0644)

    fmt.Println("Generated:", out)
}
